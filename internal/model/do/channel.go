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
	ChannelId    interface{} // 频道id
	ChannelName  interface{} // 频道名称
	ServerId     interface{} // 服务器id
	GroupId      interface{} // 分组表
	Type         interface{} // 服务器类型
	CreationDate *gtime.Time // 服务器创建时间
	CreateUserId interface{} // 服务器创建者id
	DeletedAt    *gtime.Time // 删除时间
}
