package models

import (
	"time"

	"github.com/reechou/holmes"
)

type Chatroom struct {
	ID        int64  `xorm:"pk autoincr" json:"id"`
	WeddingId int64  `xorm:"not null default 0 int unique(wedding_chatrrom)" json:"weddingId"`
	ChatType  string `xorm:"not null default '' varchar(64) unique(wedding_chatrrom)" json:"chatType"`
	Name      string `xorm:"not null default '' varchar(128)" json:"name"`
	CreatedAt int64  `xorm:"not null default 0 int" json:"createdAt"`
	UpdatedAt int64  `xorm:"not null default 0 int" json:"-"`
}

func CreateChatroom(info *Chatroom) error {
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create chatroom error: %v", err)
		return err
	}
	holmes.Info("create chatroom[%v] success.", info)

	return nil
}

func GetChatroom(info *Chatroom) (bool, error) {
	has, err := x.Where("wedding_id = ?", info.WeddingId).And("chat_type = ?", info.ChatType).Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		holmes.Debug("cannot find chatroom from info[%v]", info)
		return false, nil
	}
	return true, nil
}
