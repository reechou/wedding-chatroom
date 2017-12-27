package models

import (
	"fmt"
	"time"

	"github.com/reechou/holmes"
)

type Lottery struct {
	ID           int64  `xorm:"pk autoincr" json:"id"`
	WeddingId    int64  `xorm:"not null default 0 int index" json:"weddingId"`
	CreateUserId int64  `xorm:"not null default 0 int index" json:"createUserId"`
	Name         string `xorm:"not null default '' varchar(256)" json:"name"`
	Num          int64  `xorm:"not null default 0 int" json:"num"`
	Status       int64  `xorm:"not null default 0 int" json:"status"`
	CreatedAt    int64  `xorm:"not null default 0 int" json:"createdAt"`
	UpdatedAt    int64  `xorm:"not null default 0 int" json:"-"`
}

func CreateLottery(info *Lottery) error {
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create lottery prize error: %v", err)
		return err
	}
	holmes.Info("create lottery prize [%v] success.", info)

	return nil
}

func DelLottery(info *Lottery) error {
	if info.ID == 0 {
		return fmt.Errorf("del id cannot be nil.")
	}
	_, err := x.ID(info.ID).Delete(info)
	if err != nil {
		holmes.Error("del lottery error: %v", err)
		return err
	}

	return nil
}

func UpdateLottery(info *Lottery) error {
	now := time.Now().Unix()
	info.UpdatedAt = now
	affected, err := x.ID(info.ID).Cols("name", "num", "updated_at").Update(info)
	if affected == 0 {
		return fmt.Errorf("lottery update error")
	}
	return err
}

func GetLotteryList(weddingId int64) ([]Lottery, error) {
	var list []Lottery
	err := x.Where("wedding_id = ?", weddingId).Find(&list)
	if err != nil {
		holmes.Error("get all lottery list error: %v", err)
		return nil, err
	}
	return list, nil
}
