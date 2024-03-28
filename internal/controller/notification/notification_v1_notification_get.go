package notification

import (
	"context"
	"voichatter/internal/service"

	"voichatter/api/notification/v1"
)

func (c *ControllerV1) NotificationGet(ctx context.Context, req *v1.NotificationGetReq) (res *v1.NotificationGetRes, err error) {
	get, err := service.Notification().NotificationGet(ctx, req.ServerId)
	if err != nil {
		return nil, err
	}
	if get == nil {
		return nil, nil
	}
	return get, nil
}
