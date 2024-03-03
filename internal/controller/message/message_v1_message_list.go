package message

import (
	"context"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/message/v1"
)

func (c *ControllerV1) MessageList(ctx context.Context, req *v1.MessageListReq) (res *v1.MessageListRes, err error) {
	messageList, err := service.Message().MessageList(ctx, model.Message{
		ServerId:  req.ServerId,
		ChannelId: req.ChannelId,
	})
	if err != nil {
		return nil, err
	}
	return messageList, nil
}
