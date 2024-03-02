package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model/entity"
)

type ServerListReq struct {
	g.Meta `path:"/api/servers-list" method:"get" tags:"UserService" summary:"servers-list"`
}

type ServerListRes struct {
	ServerList *[]entity.Server
}
