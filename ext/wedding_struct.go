package ext

const (
	WEDDING_SERVICE_STATUS_OK = iota
	WEDDING_SERVICE_STATUS_ERR
)

const (
	WEDDING_SERVICE_RPC_URI = "/socket/response.do"
)

const (
	ACTION_NAME_GET_USER_LIST = "get_user_list"
	ACTION_NAME_BROADCAST_MSG = "broadcast_msg"
)

const (
	BROADCAST_MSG_TYPE_CHATROOM = 1
)

const (
	BROADCAST_MSG_NOTICE     = 1
	BROADCAST_MSG_NOT_NOTICE = 2
)

type GetWeddingUserListReqData struct {
	WeddingId int64   `json:"card_id"`
	UserList  []int64 `json:"user_list"`
}

type BroadcastMsgReqData struct {
	UserList []int64 `json:"user_list"`
	MsgType  int64   `json:"msg_type"`
	Content  string  `json:"content"`
	IsNotice int64   `json:"is_notice"`
}

type WeddingServiceReq struct {
	ActionName string      `json:"action_name"`
	Data       interface{} `json:"data"`
}

type WeddingServiceRsp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type UserInfoData struct {
	ID       int64  `json:"id"`
	NickName string `json:"nick_name"`
	Pic      string `json:"pic"`
	UserRole int64  `json:"user_status"`
}

type GetWeddingUserListRsp struct {
	Code int64          `json:"code"`
	Msg  string         `json:"msg"`
	Data []UserInfoData `json:"data"`
}
