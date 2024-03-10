package server

import (
	"context"
	"voichatter/internal/service"

	"voichatter/api/server/v1"
)

func (c *ControllerV1) ServerDel(ctx context.Context, req *v1.ServerDelReq) (res *v1.ServerDelRes, err error) {
	_, err = service.Server().ServerDel(ctx, req.ServerId)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
