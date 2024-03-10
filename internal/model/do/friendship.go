// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Friendship is the golang structure of table friendship for DAO operations like Where/Data.
type Friendship struct {
	g.Meta       `orm:"table:friendship, do:true"`
	FriendshipId interface{} // 关系id
	UserId1      interface{} // 用户1
	UserId2      interface{} // 用户2
	Date         *gtime.Time // 日期
	DeletedAt    *gtime.Time // 删除日期
}
