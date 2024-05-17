package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type ServerPagesReq struct {
	g.Meta     `path:"/servers/page" method:"get" tags:"ServerService" summary:"服务器列表"`
	ServerName string `p:"serverName"`
	Page       int    `p:"page"     example:"1" d:"1" v:"min:1#页码最小值不能低于1"  dc:"当前页码"`
	PageSize   int    `p:"pageSize" example:"5" d:"5" v:"min:1|max:200#每页数量最小值不能低于1|最大值不能大于200" dc:"每页数量"`
}

type ServerPagesRes struct {
	Servers   []*model.ServerPages `json:"servers"  `
	PageTotal int                  `json:"pageTotal"`
}
