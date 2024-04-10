package notification

import (
	"context"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/notification/v1"
)

func (c *ControllerV1) NotificationModify(ctx context.Context, req *v1.NotificationModifyReq) (res *v1.NotificationModifyRes, err error) {
	update, err := service.Notification().NotificationUpdate(ctx, model.NotificationUpdateInput{
		NotificationId: req.NotificationId,
		Title:          req.Title,
		Content:        req.Content,
		ServerId:       req.ServerId,
	})
	if err != nil {
		return nil, err
	}
	return &v1.NotificationModifyRes{
		Notification: update,
	}, nil
}
