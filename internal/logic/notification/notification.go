package notification

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "voichatter/api/notification/v1"
	"voichatter/internal/consts"
	"voichatter/internal/dao"
	"voichatter/internal/model"
	"voichatter/internal/service"
	"voichatter/utility/cache"
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
	time := gtime.Now()
	id, err := dao.Notification.Ctx(ctx).InsertAndGetId(model.Notification{
		ServerId:     in.ServerId,
		Title:        in.Title,
		Content:      in.Content,
		CreateUserId: userId,
		LastEditDate: time,
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
			LastEditDate:   time,
		},
	}, nil
}

func (s *sNotification) NotificationGet(ctx context.Context, serverId uint64) (res *v1.NotificationGetRes, err error) {
	var notification *model.Notification
	value, err := g.Redis().Get(ctx, fmt.Sprintf("%s-%d", consts.Notification, serverId))
	if err != nil {
		return nil, errResponse.OperationFailed("获取公告失败")
	}
	if err = value.Struct(&notification); err != nil {
		return nil, errResponse.OperationFailed("获取公告失败")
	}
	if notification != nil {
		return &v1.NotificationGetRes{
			Notification: notification,
		}, nil
	}
	// 从数据库获取
	err = dao.Notification.Ctx(ctx).
		Where("server_id = ?", serverId).
		Scan(&notification)
	if err != nil {
		return nil, errResponse.DbOperationError("获取公告失败")
	}
	if err = g.Redis().SetEX(ctx, fmt.Sprintf("%s-%d", consts.Notification, serverId), notification, int64(gtime.D)); err != nil {
		return nil, errResponse.OperationFailed("设置缓存失败")
	}
	return &v1.NotificationGetRes{
		Notification: notification,
	}, nil
}

func (s *sNotification) NotificationUpdate(ctx context.Context, in model.NotificationUpdateInput) (res *model.Notification, err error) {
	userId := gconv.Uint64(ctx.Value("userId"))
	count, err := dao.Member.Ctx(ctx).
		Where("server_id = ? AND user_id = ? AND s_permissions IN (?)", in.ServerId, userId, g.SliceStr{"owner", "admin"}).
		Count()
	if err != nil || count == 0 {
		return nil, errResponse.DbOperationError("权限不足")
	}
	_, err = dao.Notification.Ctx(ctx).
		Where("notification_id = ?", in.NotificationId).
		Update(in)
	if err != nil {
		return nil, errResponse.DbOperationError("操作失败")
	}

	if err = cache.DelNotification(ctx, in.ServerId); err != nil {
		return nil, errResponse.OperationFailed("删除缓存失败")
	}
	err = dao.Notification.Ctx(ctx).
		Where("notification_id = ?", in.NotificationId).
		Scan(&res)
	if err != nil {
		return nil, errResponse.DbOperationError("操作失败")
	}
	return res, nil
}

func (s *sNotification) NotificationDelete(ctx context.Context, in model.NotificationDeleteInput) error {
	userId := gconv.Uint64(ctx.Value("userId"))
	count, err := dao.Member.Ctx(ctx).
		Where("server_id = ? AND user_id = ? AND s_permissions IN (?)", in.ServerId, userId, g.SliceStr{"owner", "admin"}).
		Count()
	if err != nil || count == 0 {
		return errResponse.DbOperationError("权限不足")
	}
	_, err = dao.Notification.Ctx(ctx).
		Where("notification_id = ?", in.NotificationId).
		Delete()
	if err != nil {
		return errResponse.DbOperationError("操作失败")
	}
	return cache.DelNotification(ctx, in.ServerId)
}
