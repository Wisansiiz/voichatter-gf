// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Friendship is the golang structure for table friendship.
type Friendship struct {
	FriendshipId int64       `json:"friendship_id" ` // 关系id
	UserId1      int64       `json:"user_id_1"     ` // 用户1
	UserId2      int64       `json:"user_id_2"     ` // 用户2
	Date         *gtime.Time `json:"date"          ` // 日期
	DeletedAt    *gtime.Time `json:"deleted_at"    ` // 删除日期
}
