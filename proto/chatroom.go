package proto

const (
	RESPONSE_USER_NOT_IN_CHATROOM = 10000
)

const (
	MSG_ERROR_SYSTEM               = "系统错误"
	MSG_ERROR_USER_NOT_IN_CHATROOM = "该用户不在房间内"
	MSG_ERROR_CHATROOM_NOT_FOUND   = "该房间不存在"
	MSG_ERROR_CHATROOM_GOSSIP      = "该房间禁言中"
)

// 创建现场房间
type CreateSceneChatroomReq struct {
	WeddingId int64  `json:"weddingId"`
	Name      string `json:"name"`
}

// 进入现场房间
type EnterChatroomReq struct {
	WeddingId int64   `json:"weddingId"`
	UserId    int64   `json:"userId"`
	Latitude  float32 `json:"latitude"`  // 纬度
	Longitude float32 `json:"longitude"` // 经度
}

type TextMsg struct {
	Content string `json:"content"`
}

type PicMsg struct {
	Content string `json:"content"`
}

type AudioMsg struct {
	AudioTime int64  `json:"audioTime"`
	Content   string `json:"content"`
}

// 发送房间消息
type SendChatroomMsgReq struct {
	ChatroomId int64  `json:"chatroomId"`
	WeddingId  int64  `json:"weddingId"`
	UserId     int64  `json:"userId"`
	MsgType    int64  `json:"msgType"`
	Msg        string `json:"msg"`
}

// 获取房间历史消息
type GetChatroomMsgListReq struct {
	ChatroomId int64 `json:"chatroomId"`
	WeddingId  int64 `json:"weddingId"`
	UserId     int64 `json:"userId"`
	LastId     int64 `json:"lastId"`
}

// 获取房间成员列表
type GetChatroomMemberListReq struct {
	ChatroomId int64 `json:"chatroomId"`
	Offset     int64 `json:"offset"`
	Num        int64 `json:"num"`
	WeddingId  int64 `json:"weddingId"`
}

// 设置房间状态
// status 0: 正常状态, 1: 禁言状态
type SetChatroomStatusReq struct {
	ChatroomId int64 `json:"chatroomId"`
	UserId     int64 `json:"userId"` // 操作人 user id
	Status     int64 `json:"status"`
}
