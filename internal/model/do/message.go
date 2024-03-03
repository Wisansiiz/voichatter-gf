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
	SenderUserId interface{} //
	ChannelId    interface{} //
	Content      interface{} //
	Attachment   interface{} //
	SendDate     *gtime.Time //
	DeletedAt    *gtime.Time //
	MessageId    interface{} //
	ServerId     interface{} //
}
