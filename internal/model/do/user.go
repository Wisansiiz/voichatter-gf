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
	UserId           interface{} //
	Username         interface{} //
	Email            interface{} //
	PasswordHash     interface{} //
	AvatarUrl        interface{} //
	RegistrationDate *gtime.Time //
	LastLoginDate    *gtime.Time //
	DeletedAt        *gtime.Time //
}
