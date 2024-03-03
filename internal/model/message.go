package model

import "github.com/gogf/gf/v2/os/gtime"

type Message struct {
	ServerId  uint64
	ChannelId uint64
}

type MessageInfo struct {
	MessageId    uint64      `json:"message_id"     `
	SenderUserId uint64      `json:"sender_user_id" ` //
	ChannelId    uint64      `json:"channel_id"     ` //
	Content      string      `json:"content"        ` //
	Attachment   string      `json:"attachment"     ` //
	SendDate     *gtime.Time `json:"send_date"      ` //
	ServerId     uint64      `json:"server_id"      ` //
}
