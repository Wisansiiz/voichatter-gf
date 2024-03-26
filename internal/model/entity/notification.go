// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Notification is the golang structure for table notification.
type Notification struct {
	NotificationId uint64      `json:"notification_id" ` // 通知id
	Title          string      `json:"title"           ` // 标题
	Content        string      `json:"content"         ` // 内容
	ServerId       uint64      `json:"server_id"       ` // 属于哪个服务器
	CreateUserId   uint64      `json:"create_user_id"  ` // 创建者id
	DeletedAt      *gtime.Time `json:"deleted_at"      ` // 删除时间
}
