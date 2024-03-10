package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type ServerCreateReq struct {
	g.Meta       `path:"/servers" method:"post" tags:"ServerService" summary:""`
	ServerName   string `p:"server_name" v:"required|length:4,30#请输入服务器名|服务器名长度为:{min}到:{max}位"`
	ServerType   string `p:"server_type" v:"required#请输入服务器类型"`
	ServerImgUrl string `p:"server_img_url"`
}
type ServerCreateRes struct {
	Server *model.Server `json:"server"`
}
