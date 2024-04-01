// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Message is the golang structure of table message for DAO operations like Where/Data.
type Message struct {
	g.Meta       `orm:"table:message, do:true"`
	MessageId    interface{} // 消息id
	MessageType  interface{} // 消息类型
	ServerId     interface{} // 服务器id
	ChannelId    interface{} // 频道id
	SenderUserId interface{} // 发送者id
	Content      interface{} // 内容
	Attachment   interface{} // 引用
	SendDate     *gtime.Time // 发送日期
	DeletedAt    *gtime.Time // 删除日期
}
