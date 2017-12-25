package controller

const (
	SERVICE_NAME = "Chatroom"
)

const (
	METHOD_CREATE_CHATROOM          = "CreateChatroom"
	METHOD_ENTER_CHATROOM           = "EnterChatroom"
	METHOD_ENTER_CHATROOM_WITH_INFO = "EnterChatroomWithInfo"
	METHOD_SEND_MSG                 = "SendMessage"
	METHOD_GET_MSG_LIST             = "GetMessageList"
	METHOD_GET_CHATROOM_MEMBER_LIST = "GetChatroomMemberList"
	METHOD_SET_CHATROOM_STATUS      = "SetChatroomStatus"
)

// chatroom type
const (
	CHATROOM_WEDDING_SCENE = "wedding-scene"
)

// msg type
const (
	CHATROOM_MSG_TYPE_SYSTEM = iota
	CHATROOM_MSG_TYPE_TEXT
	CHATROOM_MSG_TYPE_AUDIO
	CHATROOM_MSG_TYPE_PIC
	CHATROOM_MSG_TYPE_RED_ENVELOPES
)

const (
	CHATROOM_MSG_TYPE_STATUS = 99
)

// chatroom status
const (
	CHATROOM_STATUS_OK = iota
	CHATROOM_STATUS_GOSSIP
)

// system msg
const (
	SYSTEM_MSG_ENTER_ROOM  = "宾客 %v 进入房间"
	SYSTEM_MSG_ROOM_GOSSIP = "房间开始禁言"
	SYSTEM_MSG_ROOM_RESUME = "房间已恢复"
)
