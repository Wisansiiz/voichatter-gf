// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Message is the golang structure for table message.
type Message struct {
	MessageId    uint64      `json:"message_id"     ` // 消息id
	MessageType  string      `json:"message_type"   ` // 消息类型
	ServerId     uint64      `json:"server_id"      ` // 服务器id
	ChannelId    uint64      `json:"channel_id"     ` // 频道id
	SenderUserId uint64      `json:"sender_user_id" ` // 发送者id
	Content      string      `json:"content"        ` // 内容
	Attachment   string      `json:"attachment"     ` // 引用
	SendDate     *gtime.Time `json:"send_date"      ` // 发送日期
	DeletedAt    *gtime.Time `json:"deleted_at"     ` // 删除日期
}
