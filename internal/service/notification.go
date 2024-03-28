// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "voichatter/api/notification/v1"
	"voichatter/internal/model"
)

type (
	INotification interface {
		NotificationCreate(ctx context.Context, in model.NotificationCreateInput) (res *v1.NotificationCreateRes, err error)
		NotificationGet(ctx context.Context, serverId uint64) (res *v1.NotificationGetRes, err error)
	}
)

var (
	localNotification INotification
)

func Notification() INotification {
	if localNotification == nil {
		panic("implement not found for interface INotification, forgot register?")
	}
	return localNotification
}

func RegisterNotification(i INotification) {
	localNotification = i
}
