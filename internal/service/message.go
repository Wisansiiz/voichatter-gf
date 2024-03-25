// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "voichatter/api/message/v1"
	"voichatter/internal/model"
)

type (
	IMessage interface {
		MessageList(ctx context.Context, in model.Message) (res *v1.MessageListRes, err error)
		MessagePages(ctx context.Context, in model.MessagePagesRes) (res *v1.MessagePagesRes, err error)
	}
)

var (
	localMessage IMessage
)

func Message() IMessage {
	if localMessage == nil {
		panic("implement not found for interface IMessage, forgot register?")
	}
	return localMessage
}

func RegisterMessage(i IMessage) {
	localMessage = i
}
