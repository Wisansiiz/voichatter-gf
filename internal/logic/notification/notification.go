package notification

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "voichatter/api/notification/v1"
	"voichatter/internal/dao"
	"voichatter/internal/model"
	"voichatter/internal/service"
	"voichatter/utility/errResponse"
)

type (
	sNotification struct{}
)

func init() {
	service.RegisterNotification(New())
}

func New() service.INotification {
	return &sNotification{}
}

func (s *sNotification) NotificationCreate(ctx context.Context, in model.NotificationCreateInput) (res *v1.NotificationCreateRes, err error) {
	userId := gconv.Uint64(ctx.Value("userId"))
	count, err := dao.Member.Ctx(ctx).
		Where("server_id = ? AND user_id = ? AND s_permissions IN (?)", in.ServerId, userId, g.SliceStr{"owner", "admin"}).
		Count()
	if err != nil || count == 0 {
		return nil, errResponse.DbOperationError("权限不足")
	}

	id, err := dao.Notification.Ctx(ctx).InsertAndGetId(model.Notification{
		ServerId:     in.ServerId,
		Title:        in.Title,
		Content:      in.Content,
		CreateUserId: userId,
	})
	if err != nil {
		return nil, errResponse.DbOperationError("操作失败")
	}
	return &v1.NotificationCreateRes{
		Notification: &model.Notification{
			NotificationId: uint64(id),
			ServerId:       in.ServerId,
			Title:          in.Title,
			Content:        in.Content,
			CreateUserId:   userId,
		},
	}, nil
}
