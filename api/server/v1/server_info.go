package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type ServerInfoReq struct {
	g.Meta   `path:"/server/{serverId}" method:"get" tags:"ServerService" summary:"servers"`
	ServerId uint64 `p:"serverId" v:"required"`
}

type ServerInfoRes struct {
	ServerInfo *model.Server `json:"serverInfo"`
}
