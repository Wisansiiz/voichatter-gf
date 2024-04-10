// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Activity is the golang structure for table activity.
type Activity struct {
	ActivityId    uint64      `json:"activity_id"     ` // 活动id
	ServerId      uint64      `json:"server_id"       ` // 服务器id
	ActivityTitle string      `json:"activity_title"  ` // 活动主题/内容
	ActivityDesc  string      `json:"activity_desc"   ` // 活动描述
	CreatorUserId uint64      `json:"creator_user_id" ` // 活动创建者id
	StartDate     *gtime.Time `json:"start_date"      ` // 开始日期
	EndDate       *gtime.Time `json:"end_date"        ` // 结束日期
	DeletedAt     *gtime.Time `json:"deleted_at"      ` // 删除日期
}
