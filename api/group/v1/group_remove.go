package v1

import "github.com/gogf/gf/v2/frame/g"

type GroupRemoveReq struct {
	g.Meta   `path:"/group/{serverId}/{groupId}" method:"delete" tags:"GroupService" summary:"删除分组"`
	ServerId uint64 `p:"serverId" v:"required"`
	GroupId  uint64 `p:"groupId" v:"required"`
}

type GroupRemoveRes struct{}
