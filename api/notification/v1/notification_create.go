package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type NotificationCreateReq struct {
	g.Meta   `path:"/notification/{serverId}" method:"post" tags:"NotificationService" summary:"创建notification"`
	ServerId uint64 `p:"serverId" v:"required"`
	Title    string `p:"title" v:"required"`
	Content  string `p:"content" v:"required"`
}

type NotificationCreateRes struct {
	Notification *model.Notification `json:"notification"`
}
