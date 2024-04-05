package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type ServerCreateReq struct {
	g.Meta       `path:"/servers" method:"post" tags:"ServerService" summary:""`
	ServerName   string `p:"serverName" v:"required|length:4,10#请输入服务器名|服务器名长度为:{min}到:{max}位"`
	ServerType   string `p:"serverType" v:"in:public,private##服务器类型只能是public或private"`
	ServerImgUrl string `p:"serverImgUrl"`
}
type ServerCreateRes struct {
	Server *model.Server `json:"server"`
}
