// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	UserId           uint64      `json:"user_id"           ` // 用户id
	Username         string      `json:"username"          ` // 用户名
	Email            string      `json:"email"             ` // 邮箱
	PasswordHash     string      `json:"password_hash"     ` // 密码
	AvatarUrl        string      `json:"avatar_url"        ` // 头像链接
	RegistrationDate *gtime.Time `json:"registration_date" ` // 注册时间
	LastLoginDate    *gtime.Time `json:"last_login_date"   ` // 最后登录日期
	DeletedAt        *gtime.Time `json:"deleted_at"        ` // 注销日期
}
