package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type NotificationGetReq struct {
	g.Meta   `path:"/notification/{serverId}" method:"get" tags:"NotificationService" summary:"创建notification"`
	ServerId uint64 `p:"serverId" v:"required"`
}

type NotificationGetRes struct {
	Notification *model.Notification `json:"notification"`
}
