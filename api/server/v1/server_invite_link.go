package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ServerInviteLinkReq struct {
	g.Meta   `path:"/servers/link/{serverId}" method:"post" tags:"ServerService" summary:"邀请链接"`
	ServerId uint64 `p:"serverId" v:"required"`
}

type ServerInviteLinkRes struct {
	Link string `json:"link"`
}
