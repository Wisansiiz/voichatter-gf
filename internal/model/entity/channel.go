// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Channel is the golang structure for table channel.
type Channel struct {
	ChannelId    uint64      `json:"channel_id"     ` // 频道id
	ChannelName  string      `json:"channel_name"   ` // 频道名称
	ServerId     uint64      `json:"server_id"      ` // 服务器id
	GroupId      uint64      `json:"group_id"       ` // 分组表
	Type         string      `json:"type"           ` // 服务器类型
	CreationDate *gtime.Time `json:"creation_date"  ` // 服务器创建时间
	CreateUserId uint64      `json:"create_user_id" ` // 服务器创建者id
	DeletedAt    *gtime.Time `json:"deleted_at"     ` // 删除时间
}
