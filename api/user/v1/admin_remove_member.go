package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserRemoveReq struct {
	g.Meta   `path:"/user/{serverId}/{userId}" method:"delete" tags:"UserService" summary:""`
	ServerId uint64 `p:"serverId" v:"required"`
	UserId   uint64 `p:"userId" v:"required"`
}

type UserRemoveRes struct {
}
