// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Member is the golang structure for table member.
type Member struct {
	MemberId     uint64      `json:"member_id"     ` //
	ServerId     uint64      `json:"server_id"     ` //
	UserId       uint64      `json:"user_id"       ` //
	JoinDate     *gtime.Time `json:"join_date"     ` //
	SPermissions string      `json:"s_permissions" ` //
	CPermissions string      `json:"c_permissions" ` //
	DeletedAt    *gtime.Time `json:"deleted_at"    ` //
}
