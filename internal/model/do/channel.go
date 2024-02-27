// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Channel is the golang structure of table channel for DAO operations like Where/Data.
type Channel struct {
	g.Meta       `orm:"table:channel, do:true"`
	ChannelId    interface{} //
	ChannelName  interface{} //
	ServerId     interface{} //
	Type         interface{} //
	CreationDate *gtime.Time //
	CreateUserId interface{} //
}
