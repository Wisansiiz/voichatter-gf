package channel

import (
	"context"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/channel/v1"
)

func (c *ControllerV1) ChannelCreate(ctx context.Context, req *v1.ChannelCreateReq) (res *v1.ChannelCreateRes, err error) {
	create, err := service.Channel().ChannelCreate(ctx, model.ChannelCreateInput{
		Type:        req.Type,
		ServerId:    req.ServerId,
		ChannelName: req.ChannelName,
		GroupId:     req.GroupId,
	})
	if err != nil {
		return nil, err
	}
	return create, nil
}
