package server

import (
	"context"
	"voichatter/internal/service"

	"voichatter/api/server/v1"
)

func (c *ControllerV1) ServerInviteLink(ctx context.Context, req *v1.ServerInviteLinkReq) (res *v1.ServerInviteLinkRes, err error) {
	link, err := service.Server().ServerInviteLink(ctx, req.ServerId)
	if err != nil {
		return nil, err
	}
	return &v1.ServerInviteLinkRes{
		Link: link,
	}, nil
}
