package channel

import (
	"context"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/channel/v1"
)

func (c *ControllerV1) ChannelModify(ctx context.Context, req *v1.ChannelModifyReq) (res *v1.ChannelModifyRes, err error) {
	modify, err := service.Channel().ChannelModify(ctx, model.ChannelModifyInput{
		ChannelId:   req.ChannelId,
		ChannelName: req.ChannelName,
		ServerId:    req.ServerId,
	})
	if err != nil {
		return nil, err
	}
	return modify, nil
}
