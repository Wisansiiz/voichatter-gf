// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Activity is the golang structure of table activity for DAO operations like Where/Data.
type Activity struct {
	g.Meta        `orm:"table:activity, do:true"`
	ActivityId    interface{} // 活动id
	ServerId      interface{} // 服务器id
	ActivityTitle interface{} // 活动主题/内容
	ActivityDesc  interface{} // 活动描述
	CreatorUserId interface{} // 活动创建者id
	StartDate     *gtime.Time // 开始日期
	EndDate       *gtime.Time // 结束日期
	DeletedAt     *gtime.Time // 删除日期
}
