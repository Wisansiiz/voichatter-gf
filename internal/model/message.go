package model

import "github.com/gogf/gf/v2/os/gtime"

type Message struct {
	ServerId  uint64
	ChannelId uint64
}

type MessagePagesReq struct {
	ServerId  uint64
	ChannelId uint64
	Page      int
	PageSize  int
}

type MessagePagesRes struct {
	MessageInfo []*MessageInfo `json:"messageInfo"`
	PageTotal   int            `json:"pageTotal"`
}

type MessageInfo struct {
	MessageId    uint64      `json:"messageId"`
	MessageType  string      `json:"messageType"`
	SenderUserId uint64      `json:"senderUserId"`
	ChannelId    uint64      `json:"channelId"`
	Content      string      `json:"content"`
	Attachment   string      `json:"attachment"`
	SendDate     *gtime.Time `json:"sendDate"`
	ServerId     uint64      `json:"serverId"`
	AvatarUrl    string      `json:"avatarUrl"`
	Username     string      `json:"senderName"`
}
