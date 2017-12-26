package ext

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/reechou/holmes"
	"github.com/reechou/wedding-chatroom/config"
)

type WeddingServiceExt struct {
	cfg    *config.Config
	client *http.Client
}

func NewWeddingServiceExt(cfg *config.Config) *WeddingServiceExt {
	return &WeddingServiceExt{
		cfg:    cfg,
		client: &http.Client{},
	}
}

func (self *WeddingServiceExt) GetWeddingUserList(reqData *GetWeddingUserListReqData) ([]UserInfoData, error) {
	request := &WeddingServiceReq{
		ActionName: ACTION_NAME_GET_USER_LIST,
		Data:       reqData,
	}

	reqBytes, err := json.Marshal(request)
	if err != nil {
		holmes.Error("json encode error: %v", err)
		return nil, err
	}

	url := "http://" + self.cfg.WeddingService.Host + WEDDING_SERVICE_RPC_URI
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBytes))
	if err != nil {
		holmes.Error("http new request error: %v", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := self.client.Do(req)
	if err != nil {
		holmes.Error("http do request error: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		holmes.Error("ioutil ReadAll error: %v", err)
		return nil, err
	}
	var response GetWeddingUserListRsp
	err = json.Unmarshal(rspBody, &response)
	if err != nil {
		holmes.Error("json decode error: %v [%s]", err, string(rspBody))
		return nil, err
	}
	if response.Code != WEDDING_SERVICE_STATUS_OK {
		holmes.Error("[wedding service] get user list [%v] result code error: %d %s", request, response.Code, response.Msg)
		return nil, fmt.Errorf("[wedding service] get user list error.")
	}

	return response.Data, nil
}

func (self *WeddingServiceExt) BroadcastMsg(reqData *BroadcastMsgReqData) error {
	request := &WeddingServiceReq{
		ActionName: ACTION_NAME_BROADCAST_MSG,
		Data:       reqData,
	}

	reqBytes, err := json.Marshal(request)
	if err != nil {
		holmes.Error("json encode error: %v", err)
		return err
	}

	url := "http://" + self.cfg.WeddingService.Host + WEDDING_SERVICE_RPC_URI
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBytes))
	if err != nil {
		holmes.Error("http new request error: %v", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := self.client.Do(req)
	if err != nil {
		holmes.Error("http do request error: %v", err)
		return err
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		holmes.Error("ioutil ReadAll error: %v", err)
		return err
	}
	var response WeddingServiceRsp
	err = json.Unmarshal(rspBody, &response)
	if err != nil {
		holmes.Error("json decode error: %v [%s]", err, string(rspBody))
		return err
	}
	if response.Code != WEDDING_SERVICE_STATUS_OK {
		holmes.Error("[wedding service] broadcast [%v] result code error: %d %s", request, response.Code, response.Msg)
		return fmt.Errorf("[wedding service] broadcast error.")
	}
	holmes.Debug("broadcast msg type[%d] content[%s] success.", reqData.MsgType, reqData.Content)

	return nil
}
