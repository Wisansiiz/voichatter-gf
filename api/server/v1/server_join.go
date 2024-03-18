package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type ServerJoinReq struct {
	g.Meta   `path:"/servers/{serverId}" method:"post" tags:"ServerService" summary:"加入服务器"`
	ServerId uint64 `p:"serverId" v:"required"`
}

type ServerJoinRes struct {
	Server *model.Server `json:"server"`
}
