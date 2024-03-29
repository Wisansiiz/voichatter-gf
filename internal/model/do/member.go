// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Member is the golang structure of table member for DAO operations like Where/Data.
type Member struct {
	g.Meta       `orm:"table:member, do:true"`
	MemberId     interface{} // 成员id
	ServerId     interface{} // 服务器id
	UserId       interface{} // 用户id
	JoinDate     *gtime.Time // 加入日期
	SPermissions interface{} // 权限
	DeletedAt    *gtime.Time // 删除日期
}
