package controller

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/reechou/holmes"
	"github.com/reechou/wedding-chatroom/models"
	"github.com/reechou/wedding-chatroom/proto"
	"github.com/reechou/wedding-chatroom/ext"
)

func (self *Logic) runRpc(w http.ResponseWriter, r *http.Request) {
	serviceName := r.Header.Get("ServerName")
	if serviceName != SERVICE_NAME {
		holmes.Error("rpc service name[%s] is not self[%s]", serviceName, SERVICE_NAME)
		return
	}
	methodName := r.Header.Get("MethodName")
	switch methodName {
	case METHOD_CREATE_CHATROOM:
		self.CreateSceneChatroom(w, r)
	case METHOD_ENTER_CHATROOM:
		self.EnterChatroom(w, r)
	case METHOD_SEND_MSG:
		self.SendChatroomMsg(w, r)
	case METHOD_GET_MSG_LIST:
		self.GetChatroomMessageList(w, r)
	case METHOD_GET_CHATROOM_MEMBER_LIST:
		self.GetChatroomMemberList(w, r)
	}
}

type MessageDetail struct {
	Msg  *models.ChatroomMessage `json:"msg"`
	User *ext.UserInfoData       `json:"user"`
}

type ChatroomMemberList struct {
	Count int64              `json:"count"`
	List  []ext.UserInfoData `json:"list"`
}

func (self *Logic) systemMsg(chatroomId, weddingId, userId int64, msg string) {
	message := &models.ChatroomMessage{
		ChatroomId: chatroomId,
		UserId:     userId,
		MsgType:    CHATROOM_MSG_TYPE_SYSTEM,
		Msg:        msg,
	}
	md := &MessageDetail{}
	if message.UserId != 0 && weddingId != 0 {
		userIdReq := []int64{message.UserId}
		getUserListReq := &ext.GetWeddingUserListReqData{
			WeddingId: weddingId,
			UserList:  userIdReq,
		}
		userList, err := self.weddingExt.GetWeddingUserList(getUserListReq)
		if err != nil {
			holmes.Error("get wedding user list error: %v", err)
		} else {
			if len(userList) != 0 {
				message.Msg = fmt.Sprintf(msg, userList[0].NickName)
				md.User = &userList[0]
			}
		}
	}
	// create chatroom system message
	err := models.CreateChatroomMessage(message)
	if err != nil {
		holmes.Error("create chatroom system message error: %v", err)
		return
	}
	md.Msg = message
	memberList, err := models.GetAllChatroomMemberList(chatroomId)
	if err != nil {
		holmes.Error("get all chatroom member list error: %v", err)
	} else {
		chatroomMemberList := make([]int64, len(memberList))
		for i := 0; i < len(memberList); i++ {
			chatroomMemberList[i] = memberList[i].UserId
		}
		// broadcast
		self.broadcastChatroomMsgV2(chatroomMemberList, md, ext.BROADCAST_MSG_NOT_NOTICE)
	}
}

func (self *Logic) broadcastChatroomMsgV2(userIdList []int64,
	msg *MessageDetail,
	isNotice int64) {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		holmes.Error("broadcast msg json marshal error: %v", err)
		return
	}
	broadcastMsgReq := &ext.BroadcastMsgReqData{
		UserList: userIdList,
		MsgType:  ext.BROADCAST_MSG_TYPE_CHATROOM,
		Content:  string(msgBytes),
		IsNotice: isNotice,
	}
	self.weddingExt.BroadcastMsg(broadcastMsgReq)
}

func (self *Logic) broadcastChatroomMsg(userIdList []int64,
	msg *models.ChatroomMessage,
	weddingId int64,
	isNotice int64) {
	md := &MessageDetail{
		Msg: msg,
	}
	if msg.UserId != 0 && weddingId != 0 {
		userIdReq := []int64{msg.UserId}
		getUserListReq := &ext.GetWeddingUserListReqData{
			WeddingId: weddingId,
			UserList:  userIdReq,
		}
		userList, err := self.weddingExt.GetWeddingUserList(getUserListReq)
		if err != nil {
			holmes.Error("get wedding user list error: %v", err)
		} else {
			if len(userList) != 0 {
				md.User = &userList[0]
			}
		}
	}
	msgBytes, err := json.Marshal(md)
	if err != nil {
		holmes.Error("broadcast msg json marshal error: %v", err)
		return
	}
	broadcastMsgReq := &ext.BroadcastMsgReqData{
		UserList: userIdList,
		MsgType:  ext.BROADCAST_MSG_TYPE_CHATROOM,
		Content:  string(msgBytes),
		IsNotice: isNotice,
	}
	self.weddingExt.BroadcastMsg(broadcastMsgReq)
}

func (self *Logic) CreateSceneChatroom(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &proto.CreateSceneChatroomReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("CreateChatroom json decode error: %v", err)
		return
	}

	chatroom := &models.Chatroom{
		WeddingId: req.WeddingId,
		ChatType:  CHATROOM_WEDDING_SCENE,
		Name:      req.Name,
	}
	if err := models.CreateChatroom(chatroom); err != nil {
		holmes.Error("create chatroom error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		rsp.Msg = proto.MSG_ERROR_SYSTEM
	}
}

func (self *Logic) EnterChatroom(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &proto.EnterChatroomReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("EnterChatroom json decode error: %v", err)
		return
	}

	chatroom := &models.Chatroom{
		WeddingId: req.WeddingId,
		ChatType:  CHATROOM_WEDDING_SCENE,
	}
	has, err := models.GetChatroom(chatroom)
	if err != nil {
		holmes.Error("get chatroom error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		rsp.Msg = proto.MSG_ERROR_SYSTEM
		return
	}
	if !has {
		if err = models.CreateChatroom(chatroom); err != nil {
			holmes.Error("create chatroom error: %v", err)
			rsp.Code = proto.RESPONSE_ERR
			rsp.Msg = proto.MSG_ERROR_SYSTEM
			return
		}
	}

	member := &models.ChatroomMember{
		ChatroomId: chatroom.ID,
		UserId:     req.UserId,
	}
	has, err = models.GetChatroomMember(member)
	if err != nil {
		holmes.Error("get chatroom member error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		rsp.Msg = proto.MSG_ERROR_SYSTEM
		return
	}
	if has {
		rsp.Data = chatroom.ID
		return
	}
	if err := models.CreateChatroomMember(member); err != nil {
		holmes.Error("create chatroom member error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		rsp.Msg = proto.MSG_ERROR_SYSTEM
		return
	}
	rsp.Data = chatroom.ID
	// broadcast
	self.systemMsg(chatroom.ID, req.WeddingId, req.UserId, SYSTEM_MSG_ENTER_ROOM)
}

func (self *Logic) SendChatroomMsg(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &proto.SendChatroomMsgReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("SendChatroomMsg json decode error: %v", err)
		return
	}
	
	chatroomMember := &models.ChatroomMember{
		ChatroomId: req.ChatroomId,
		UserId:     req.UserId,
	}
	has, err := models.GetChatroomMember(chatroomMember)
	if err != nil {
		holmes.Error("get chatroom member error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		rsp.Msg = proto.MSG_ERROR_SYSTEM
		return
	}
	if !has {
		rsp.Code = proto.RESPONSE_USER_NOT_IN_CHATROOM
		rsp.Msg = proto.MSG_ERROR_USER_NOT_IN_CHATROOM
		return
	}

	chatroomMessage := &models.ChatroomMessage{
		ChatroomId: req.ChatroomId,
		UserId:     req.UserId,
		MsgType:    req.MsgType,
		Msg:        req.Msg,
	}
	if err := models.CreateChatroomMessage(chatroomMessage); err != nil {
		holmes.Error("create chatroom message error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		rsp.Msg = proto.MSG_ERROR_SYSTEM
		return
	}
	// broadcast
	memberList, err := models.GetAllChatroomMemberList(req.ChatroomId)
	if err != nil {
		holmes.Error("get all chatroom member list error: %v", err)
	} else {
		chatroomMemberList := make([]int64, len(memberList))
		for i := 0; i < len(memberList); i++ {
			chatroomMemberList[i] = memberList[i].UserId
		}
		self.broadcastChatroomMsg(chatroomMemberList, chatroomMessage, req.WeddingId, ext.BROADCAST_MSG_NOT_NOTICE)
	}
}

func (self *Logic) GetChatroomMessageList(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()
	
	if r.Method != "POST" {
		return
	}
	
	req := &proto.GetChatroomMsgListReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("GetChatroomMessageList json decode error: %v", err)
		return
	}
	
	msgList, err := models.GetChatroomMessageList(req.ChatroomId, req.LastId, 10)
	if err != nil {
		holmes.Error("get chatroom message list error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		rsp.Msg = proto.MSG_ERROR_SYSTEM
		return
	}
	holmes.Debug("msg list: %v", msgList)
	// get message user info list
	var userIdList []int64
	userMap := make(map[int64]*ext.UserInfoData)
	for _, v := range msgList {
		if _, ok := userMap[v.UserId]; ok {
			continue
		}
		userIdList = append(userIdList, v.UserId)
		userMap[v.UserId] = nil
	}
	getUserListReq := &ext.GetWeddingUserListReqData{
		WeddingId: req.WeddingId,
		UserList:  userIdList,
	}
	userList, err := self.weddingExt.GetWeddingUserList(getUserListReq)
	if err != nil {
		holmes.Error("get wedding user list error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		rsp.Msg = proto.MSG_ERROR_SYSTEM
		return
	}
	for i := 0; i < len(userList); i++ {
		userMap[userList[i].ID] = &userList[i]
	}
	holmes.Debug("user map: %v", userMap)
	var msgs []MessageDetail
	for i := 0; i < len(msgList); i++ {
		md := MessageDetail{
			Msg: &msgList[i],
		}
		uv, ok := userMap[msgList[i].UserId]
		if ok {
			md.User = uv
		}
		msgs = append(msgs, md)
	}
	rsp.Data = msgs
}

func (self *Logic) GetChatroomMemberList(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()
	
	if r.Method != "POST" {
		return
	}
	
	req := &proto.GetChatroomMemberListReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("GetChatroomMemberList json decode error: %v", err)
		return
	}
	
	list := &ChatroomMemberList{}
	var err error
	list.Count, err = models.GetChatroomMemberCount(req.ChatroomId)
	if err != nil {
		holmes.Error("get chatroom[%d] member count error: %v", req.ChatroomId, err)
		rsp.Code = proto.RESPONSE_ERR
		rsp.Msg = proto.MSG_ERROR_SYSTEM
		return
	}
	
	memberList, err := models.GetChatroomMemberList(req.ChatroomId, req.Offset, req.Num)
	if err != nil {
		holmes.Error("get chatroom[%d] member list error: %v", req.ChatroomId, err)
		rsp.Code = proto.RESPONSE_ERR
		rsp.Msg = proto.MSG_ERROR_SYSTEM
		return
	}
	var userIdList []int64
	for i := 0; i < len(memberList); i++ {
		userIdList = append(userIdList, memberList[i].UserId)
	}
	getUserListReq := &ext.GetWeddingUserListReqData{
		WeddingId: req.WeddingId,
		UserList:  userIdList,
	}
	holmes.Debug("get user list req: %v", getUserListReq)
	list.List, err = self.weddingExt.GetWeddingUserList(getUserListReq)
	if err != nil {
		holmes.Error("get chatroom[%d] wedding user list error: %v", req.ChatroomId, err)
		rsp.Code = proto.RESPONSE_ERR
		rsp.Msg = proto.MSG_ERROR_SYSTEM
		return
	}
	rsp.Data = list
}
