// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Server is the golang structure for table server.
type Server struct {
	ServerId      uint64      `json:"server_id"       ` //
	ServerName    string      `json:"server_name"     ` //
	CreatorUserId uint64      `json:"creator_user_id" ` //
	CreationDate  *gtime.Time `json:"creation_date"   ` //
	ServerTheme   string      `json:"server_theme"    ` //
	CreateDate    *gtime.Time `json:"create_date"     ` //
	ServerType    string      `json:"server_type"     ` //
	ServerImgUrl  string      `json:"server_img_url"  ` //
	DeletedAt     *gtime.Time `json:"deleted_at"      ` //
}
