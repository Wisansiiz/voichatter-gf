package server

import (
	"context"
	"voichatter/internal/service"

	"voichatter/api/server/v1"
)

func (c *ControllerV1) ServerCount(ctx context.Context, req *v1.ServerCountReq) (res *v1.ServerCountRes, err error) {
	count, err := service.Server().ServerCount(ctx, req)
	if err != nil {
		return nil, err
	}
	return count, nil
}
