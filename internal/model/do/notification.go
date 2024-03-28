// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Notification is the golang structure of table notification for DAO operations like Where/Data.
type Notification struct {
	g.Meta         `orm:"table:notification, do:true"`
	NotificationId interface{} // 通知id
	Title          interface{} // 标题
	Content        interface{} // 内容
	ServerId       interface{} // 属于哪个服务器
	CreateUserId   interface{} // 创建者id
	LastEditDate   *gtime.Time // 最后编辑时间
	DeletedAt      *gtime.Time // 删除时间
}
