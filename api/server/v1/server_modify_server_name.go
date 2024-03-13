package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type ServerModifyNameReq struct {
	g.Meta     `path:"/server/{serverId}" method:"put" tags:"ServerService" summary:"servers"`
	ServerId   uint64 `p:"serverId" v:"required"`
	ServerName string `p:"serverName" v:"required|length:4,30#请输入服务器名|服务器名长度为:{min}到:{max}位"`
}

type ServerModifyNameRes struct {
	ServerInfo *model.Server `json:"serverInfo"`
}
