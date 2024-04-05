// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Server is the golang structure of table server for DAO operations like Where/Data.
type Server struct {
	g.Meta            `orm:"table:server, do:true"`
	ServerId          interface{} // 服务器id
	ServerName        interface{} // 服务器名称
	ServerDescription interface{} // 服务器描述
	CreatorUserId     interface{} // 服务器创建者id
	CreateDate        *gtime.Time // 创建日期
	ServerType        interface{} // 服务器类型
	ServerImgUrl      interface{} // 服务器头像链接
	DeletedAt         *gtime.Time // 删除日期
}
