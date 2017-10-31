package controller

const (
	SERVICE_NAME = "Chatroom"
)

const (
	METHOD_CREATE_CHATROOM          = "CreateChatroom"
	METHOD_ENTER_CHATROOM           = "EnterChatroom"
	METHOD_SEND_MSG                 = "SendMessage"
	METHOD_GET_MSG_LIST             = "GetMessageList"
	METHOD_GET_CHATROOM_MEMBER_LIST = "GetChatroomMemberList"
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
)

// system msg
const (
	SYSTEM_MSG_ENTER_ROOM = "宾客 %v 进入房间"
)
