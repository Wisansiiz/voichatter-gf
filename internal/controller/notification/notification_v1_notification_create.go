package notification

import (
	"context"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/notification/v1"
)

func (c *ControllerV1) NotificationCreate(ctx context.Context, req *v1.NotificationCreateReq) (res *v1.NotificationCreateRes, err error) {
	create, err := service.Notification().NotificationCreate(ctx, model.NotificationCreateInput{
		ServerId: req.ServerId,
		Title:    req.Title,
		Content:  req.Content,
	})
	if err != nil {
		return nil, err
	}
	return create, nil
}
