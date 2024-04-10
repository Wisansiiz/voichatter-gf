package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type NotificationModifyReq struct {
	g.Meta         `path:"/notification/{serverId}" method:"put" tags:"NotificationService" summary:"更新notification"`
	ServerId       uint64 `p:"serverId" v:"required"`
	NotificationId uint64 `p:"notificationId" v:"required"`
	Title          string `p:"title" v:"required"`
	Content        string `p:"content" v:"required"`
}

type NotificationModifyRes struct {
	Notification *model.Notification `json:"notification"`
}
