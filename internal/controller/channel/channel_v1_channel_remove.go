package channel

import (
	"context"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/channel/v1"
)

func (c *ControllerV1) ChannelRemove(ctx context.Context, req *v1.ChannelRemoveReq) (res *v1.ChannelRemoveRes, err error) {
	_, err = service.Channel().ChannelRemove(ctx, model.ChannelRemoveInput{
		ChannelId: req.ChannelId,
		ServerId:  req.ServerId,
	})
	if err != nil {
		return nil, err
	}
	return
}
