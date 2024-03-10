// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Member is the golang structure for table member.
type Member struct {
	MemberId     uint64      `json:"member_id"     ` // 成员id
	ServerId     uint64      `json:"server_id"     ` // 服务器id
	UserId       uint64      `json:"user_id"       ` // 用户id
	JoinDate     *gtime.Time `json:"join_date"     ` // 加入日期
	SPermissions string      `json:"s_permissions" ` // 权限
	DeletedAt    *gtime.Time `json:"deleted_at"    ` // 删除日期
}
