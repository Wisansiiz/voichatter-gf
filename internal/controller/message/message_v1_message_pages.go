package message

import (
	"context"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/message/v1"
)

func (c *ControllerV1) MessagePages(ctx context.Context, req *v1.MessagePagesReq) (res *v1.MessagePagesRes, err error) {
	pages, err := service.Message().MessagePages(ctx, model.MessagePagesRes{
		ServerId:  req.ServerId,
		ChannelId: req.ChannelId,
		Page:      req.Page,
		PageSize:  req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	return pages, nil
}
