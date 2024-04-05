package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type ServerInfoUpdReq struct {
	g.Meta            `path:"/server-info/{serverId}" method:"put" tags:"ServerService" summary:"server"`
	ServerId          uint64 `p:"serverId"          v:"required#请输入服务器id"`
	ServerName        string `p:"serverName"        v:"length:4,10#服务器名长度为:{min}到:{max}位"`
	ServerType        string `p:"serverType"        v:"in:public,private#服务器类型只能是public或private"`
	ServerDescription string `p:"serverDescription" v:"length:4,255#服务器描述长度为:{min}到:{max}位"`
}

type ServerInfoUpdRes struct {
	ServerInfo *model.Server `json:"serverInfo"`
}
