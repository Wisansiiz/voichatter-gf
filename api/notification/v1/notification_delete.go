package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type NotificationDeleteReq struct {
	g.Meta         `path:"/notification/{serverId}/{notificationId}" method:"delete" tags:"NotificationService" summary:"删除notification"`
	ServerId       uint64 `p:"serverId" v:"required"`
	NotificationId uint64 `p:"notificationId" v:"required"`
}

type NotificationDeleteRes struct {
}
