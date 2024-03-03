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
	g.Meta        `orm:"table:server, do:true"`
	ServerId      interface{} //
	ServerName    interface{} //
	CreatorUserId interface{} //
	CreationDate  *gtime.Time //
	CreateDate    *gtime.Time //
	ServerType    interface{} //
	ServerImgUrl  interface{} //
	DeletedAt     *gtime.Time //
}
