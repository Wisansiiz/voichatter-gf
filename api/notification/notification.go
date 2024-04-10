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
	NotificationDelete(ctx context.Context, req *v1.NotificationDeleteReq) (res *v1.NotificationDeleteRes, err error)
	NotificationGet(ctx context.Context, req *v1.NotificationGetReq) (res *v1.NotificationGetRes, err error)
	NotificationModify(ctx context.Context, req *v1.NotificationModifyReq) (res *v1.NotificationModifyRes, err error)
}
