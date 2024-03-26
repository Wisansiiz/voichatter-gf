// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Group is the golang structure for table group.
type Group struct {
	GroupId   uint64      `json:"group_id"   ` //
	ServerId  uint64      `json:"server_id"  ` // 服务器id
	GroupName string      `json:"group_name" ` // 分组名称
	DeletedAt *gtime.Time `json:"deleted_at" ` // 删除日期
}
