// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Server is the golang structure for table server.
type Server struct {
	ServerId          uint64      `json:"server_id"          ` // 服务器id
	ServerName        string      `json:"server_name"        ` // 服务器名称
	ServerDescription string      `json:"server_description" ` // 服务器描述
	CreatorUserId     uint64      `json:"creator_user_id"    ` // 服务器创建者id
	CreateDate        *gtime.Time `json:"create_date"        ` // 创建日期
	ServerType        string      `json:"server_type"        ` // 服务器类型
	ServerImgUrl      string      `json:"server_img_url"     ` // 服务器头像链接
	DeletedAt         *gtime.Time `json:"deleted_at"         ` // 删除日期
}
