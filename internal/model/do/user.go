// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta           `orm:"table:user, do:true"`
	UserId           interface{} // 用户id
	Username         interface{} // 用户名
	Email            interface{} // 邮箱
	PasswordHash     interface{} // 密码
	AvatarUrl        interface{} // 头像链接
	RegistrationDate *gtime.Time // 注册时间
	LastLoginDate    *gtime.Time // 最后登录日期
	DeletedAt        *gtime.Time // 注销日期
}
