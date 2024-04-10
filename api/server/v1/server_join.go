package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type ServerJoinReq struct {
	g.Meta   `path:"/servers/{serverId}" method:"post" tags:"ServerService" summary:"加入服务器"`
	ServerId uint64 `p:"serverId" v:""`
	Link     string `p:"link"     v:"length:1,16#链接长度必须在:{min}到:{max}之间" dc:"链接"`
}

type ServerJoinRes struct {
	Server *model.Server `json:"server"`
}
