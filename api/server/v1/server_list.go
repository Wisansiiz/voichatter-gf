package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model/entity"
)

type ServerListReq struct {
	g.Meta `path:"/servers" method:"get" tags:"ServerService" summary:"servers"`
}

type ServerListRes struct {
	ServerList *[]entity.Server
}
