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
	MemberId     interface{} //
	ServerId     interface{} //
	UserId       interface{} //
	JoinDate     *gtime.Time //
	SPermissions interface{} //
	CPermissions interface{} //
	DeletedAt    *gtime.Time //
}
