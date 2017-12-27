package controller

import (
	"encoding/json"
	"net/http"

	"github.com/reechou/holmes"
	"github.com/reechou/wedding-chatroom/models"
	"github.com/reechou/wedding-chatroom/proto"
)

func (self *Logic) CreateLottery(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &proto.CreateLotteryReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("json decode error: %v", err)
		return
	}

	lottery := &models.Lottery{
		WeddingId:    req.WeddingId,
		CreateUserId: req.CreateUserId,
		Name:         req.Name,
		Num:          req.Num,
		Status:       LOTTERY_STATUS_NOT_START,
	}
	if err := models.CreateLottery(lottery); err != nil {
		holmes.Error("create lottery error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		rsp.Msg = proto.MSG_ERROR_SYSTEM
	}
}

func (self *Logic) DeleteLottery(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &proto.DeleteLotteryReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("json decode error: %v", err)
		return
	}

	lottery := &models.Lottery{ID: req.LotteryId}
	if err := models.DelLottery(lottery); err != nil {
		holmes.Error("delete lottery error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		rsp.Msg = proto.MSG_ERROR_SYSTEM
	}
}

func (self *Logic) UpdateLottery(w http.ResponseWriter, r *http.Request) {
	rsp := &proto.Response{Code: proto.RESPONSE_OK}
	defer func() {
		WriteJSON(w, http.StatusOK, rsp)
	}()

	if r.Method != "POST" {
		return
	}

	req := &proto.UpdateLotteryReq{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		holmes.Error("json decode error: %v", err)
		return
	}

	lottery := &models.Lottery{ID: req.LotteryId, Name: req.Name, Num: req.Num}
	if err := models.UpdateLottery(lottery); err != nil {
		holmes.Error("update lottery error: %v", err)
		rsp.Code = proto.RESPONSE_ERR
		rsp.Msg = proto.MSG_ERROR_SYSTEM
	}
}
