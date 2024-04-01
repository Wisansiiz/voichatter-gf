package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type ServerSearchReq struct {
	g.Meta     `path:"/servers/search" method:"get" tags:"ServerService" summary:""`
	ServerName string `p:"serverName" v:"required"`
}

type ServerSearchRes struct {
	Servers []*model.Server `json:"servers"`
}
