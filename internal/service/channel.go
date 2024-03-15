// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "voichatter/api/channel/v1"
	"voichatter/internal/model"
)

type (
	IChannel interface {
		ChannelCreate(ctx context.Context, in model.ChannelCreateInput) (res *v1.ChannelCreateRes, err error)
		ChannelModify(ctx context.Context, in model.ChannelModifyInput) (res *v1.ChannelModifyRes, err error)
		ChannelRemove(ctx context.Context, in model.ChannelRemoveInput) (res *v1.ChannelRemoveRes, err error)
	}
)

var (
	localChannel IChannel
)

func Channel() IChannel {
	if localChannel == nil {
		panic("implement not found for interface IChannel, forgot register?")
	}
	return localChannel
}

func RegisterChannel(i IChannel) {
	localChannel = i
}
