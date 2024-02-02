package model

type WsUserMsg struct {
	MsgId    uint32 `json:"msg_id"`
	SenderId string `json:"sender_id"`
	Content  string `json:"content"`
	MsgType  uint   `json:"msgType"`
	ReplayId uint   `json:"reply_id"`
	SendAt   int64  `json:"send_at"`
	DialogId uint32 `json:"dialog_id"`
}

type WsGroupMsg struct {
	MsgId    uint32 `json:"msg_id"`
	GroupId  int64  `json:"group_id"`
	UserId   string `json:"uid"`
	Content  string `json:"content"`
	MsgType  uint   `json:"msgType"`
	ReplayId uint   `json:"reply_id"`
	SendAt   int64  `json:"send_at"`
	DialogId uint32 `json:"dialog_id"`
}
