package notification

import (
	"context"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/notification/v1"
)

func (c *ControllerV1) NotificationDelete(ctx context.Context, req *v1.NotificationDeleteReq) (res *v1.NotificationDeleteRes, err error) {
	err = service.Notification().NotificationDelete(ctx, model.NotificationDeleteInput{
		NotificationId: req.NotificationId,
		ServerId:       req.ServerId,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
