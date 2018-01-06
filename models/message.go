package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/reechou/holmes"
)

const (
	CHATROOM_MESSAGE_NUM = 100
)

type ChatroomMessage struct {
	ID         int64  `xorm:"pk autoincr" json:"id"`
	ChatroomId int64  `xorm:"not null default 0 int index" json:"chatroomId"`
	UserId     int64  `xorm:"not null default 0 int" json:"userId"`
	MsgType    int64  `xorm:"not null default 0 int index" json:"msgType"`
	Msg        string `xorm:"not null default '' varchar(1024)" json:"msg"` // json
	CreatedAt  int64  `xorm:"not null default 0 int" json:"createdAt"`
	UpdatedAt  int64  `xorm:"not null default 0 int" json:"-"`
}

func (self *ChatroomMessage) TableName() string {
	return "chatroom_message_" + strconv.Itoa(int(self.ChatroomId)%CHATROOM_MESSAGE_NUM)
}

func CreateChatroomMessage(info *ChatroomMessage) error {
	if info.ChatroomId == 0 || info.UserId == 0 {
		return fmt.Errorf("chatroom or user cannot be 0")
	}

	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create chatroom message error: %v", err)
		return err
	}
	holmes.Info("create chatroom message[%v] success.", info)

	return nil
}

func GetChatroomMessageList(chatroomId, lastId, num int64) ([]ChatroomMessage, error) {
	var list []ChatroomMessage
	var err error
	if lastId == 0 {
		err = x.Table(&ChatroomMessage{ChatroomId: chatroomId}).
			Where("chatroom_id = ?", chatroomId).
			Desc("id").
			Limit(int(num)).Find(&list)
	} else {
		err = x.Table(&ChatroomMessage{ChatroomId: chatroomId}).
			Where("chatroom_id = ?", chatroomId).
			And("id < ?", lastId).
			Desc("id").
			Limit(int(num)).Find(&list)
	}
	if err != nil {
		holmes.Error("get chatroom member list error: %v", err)
		return nil, err
	}
	return list, nil
}

type BroadcastMessage struct {
	ID         int64  `xorm:"pk autoincr" json:"id"`
	ChatroomId int64  `xorm:"not null default 0 int index" json:"chatroomId"`
	UserId     int64  `xorm:"not null default 0 int" json:"userId"`
	MsgType    int64  `xorm:"not null default 0 int index" json:"msgType"`
	Msg        string `xorm:"not null default '' varchar(1024)" json:"msg"` // json
	CreatedAt  int64  `xorm:"not null default 0 int" json:"createdAt"`
	UpdatedAt  int64  `xorm:"not null default 0 int" json:"-"`
}

func CreateBroadcastMessage(info *BroadcastMessage) error {
	if info.ChatroomId == 0 || info.UserId == 0 {
		return fmt.Errorf("chatroom or user cannot be 0")
	}

	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create broadcast message error: %v", err)
		return err
	}
	holmes.Info("create broadcast message[%v] success.", info)

	return nil
}
