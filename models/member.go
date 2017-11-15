package models

import (
	"time"

	"github.com/reechou/holmes"
)

type ChatroomMember struct {
	ID         int64 `xorm:"pk autoincr" json:"id"`
	ChatroomId int64 `xorm:"not null default 0 int unique(chatroom_member)" json:"chatroomId"`
	UserId     int64 `xorm:"not null default 0 int unique(chatroom_member)" json:"userId"`
	CreatedAt  int64 `xorm:"not null default 0 int" json:"createdAt"`
	UpdatedAt  int64 `xorm:"not null default 0 int" json:"-"`
}

func CreateChatroomMember(info *ChatroomMember) error {
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create chatroom member error: %v", err)
		return err
	}
	holmes.Info("create chatroom member[%v] success.", info)

	return nil
}

func DeleteChatroomMember(info *ChatroomMember) error {
	_, err := x.Where("chatroom_id = ?", info.ChatroomId).
		And("user_id = ?", info.UserId).
		Delete(info)
	if err != nil {
		holmes.Error("del chatroom member error: %v", err)
		return err
	}

	return nil
}

func GetChatroomMember(info *ChatroomMember) (bool, error) {
	has, err := x.Where("chatroom_id = ?", info.ChatroomId).And("user_id = ?", info.UserId).Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		holmes.Debug("cannot find chatroom member from info[%v]", info)
		return false, nil
	}
	return true, nil
}

func GetAllChatroomMemberList(chatroomId int64) ([]ChatroomMember, error) {
	var list []ChatroomMember
	err := x.Where("chatroom_id = ?", chatroomId).Find(&list)
	if err != nil {
		holmes.Error("get all chatroom member list error: %v", err)
		return nil, err
	}
	return list, nil
}

func GetChatroomMemberCount(chatroomId int64) (int64, error) {
	count, err := x.Where("chatroom_id = ?", chatroomId).Count(&ChatroomMember{})
	if err != nil {
		holmes.Error("get chatroom member list count error: %v", err)
		return 0, err
	}
	return count, nil
}

func GetChatroomMemberList(chatroomId, offset, num int64) ([]ChatroomMember, error) {
	var list []ChatroomMember
	err := x.Where("chatroom_id = ?", chatroomId).Limit(int(num), int(offset)).Find(&list)
	if err != nil {
		holmes.Error("get chatroom member list error: %v", err)
		return nil, err
	}
	return list, nil
}
