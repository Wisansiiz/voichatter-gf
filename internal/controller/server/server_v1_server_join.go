package server

import (
	"context"
	"voichatter/internal/service"

	"voichatter/api/server/v1"
)

func (c *ControllerV1) ServerJoin(ctx context.Context, req *v1.ServerJoinReq) (res *v1.ServerJoinRes, err error) {
	server, err := service.Server().ServerJoin(ctx, req.ServerId)
	if err != nil {
		return nil, err
	}
	return server, nil
}
