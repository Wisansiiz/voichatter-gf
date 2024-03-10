// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Group is the golang structure of table group for DAO operations like Where/Data.
type Group struct {
	g.Meta    `orm:"table:group, do:true"`
	GroupId   interface{} //
	ServerId  interface{} // 服务器id
	GroupName interface{} // 分组名称
	ChannelId interface{} // 频道id
	DeletedAt *gtime.Time // 删除日期
}
