package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ServerDelReq struct {
	g.Meta   `path:"/servers/{serverId}" method:"delete" tags:"ServerService" summary:"servers"`
	ServerId uint64 `p:"serverId" v:"required"`
}

type ServerDelRes struct{}
