// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	UserId           uint64      `json:"user_id"           ` //
	Username         string      `json:"username"          ` //
	Email            string      `json:"email"             ` //
	PasswordHash     string      `json:"password_hash"     ` //
	AvatarUrl        string      `json:"avatar_url"        ` //
	RegistrationDate *gtime.Time `json:"registration_date" ` //
	LastLoginDate    *gtime.Time `json:"last_login_date"   ` //
	DeletedAt        *gtime.Time `json:"deleted_at"        ` //
}
