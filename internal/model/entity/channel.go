// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Channel is the golang structure for table channel.
type Channel struct {
	ChannelId    uint64      `json:"channel_id"     ` //
	ChannelName  string      `json:"channel_name"   ` //
	ServerId     uint64      `json:"server_id"      ` //
	Type         string      `json:"type"           ` //
	CreationDate *gtime.Time `json:"creation_date"  ` //
	CreateUserId uint64      `json:"create_user_id" ` //
}
