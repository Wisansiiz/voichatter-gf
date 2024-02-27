// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Message is the golang structure for table message.
type Message struct {
	SenderUserId uint64      `json:"sender_user_id" ` //
	ChannelId    uint64      `json:"channel_id"     ` //
	Content      string      `json:"content"        ` //
	Attachment   string      `json:"attachment"     ` //
	SendDate     *gtime.Time `json:"send_date"      ` //
	DeletedAt    *gtime.Time `json:"deleted_at"     ` //
	MessageId    uint64      `json:"message_id"     ` //
}
