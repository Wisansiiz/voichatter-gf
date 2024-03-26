// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package notification

import (
	"context"

	"voichatter/api/notification/v1"
)

type INotificationV1 interface {
	NotificationCreate(ctx context.Context, req *v1.NotificationCreateReq) (res *v1.NotificationCreateRes, err error)
}
