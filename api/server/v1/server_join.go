package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ServerJoinReq struct {
	g.Meta   `path:"/servers/{serverId}" method:"post" tags:"ServerService" summary:""`
	ServerId uint64 `p:"serverId" v:"required"`
}

type ServerJoinRes struct {
}
