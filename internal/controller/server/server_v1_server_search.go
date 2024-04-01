package server

import (
	"context"
	"voichatter/internal/service"

	"voichatter/api/server/v1"
)

func (c *ControllerV1) ServerSearch(ctx context.Context, req *v1.ServerSearchReq) (res *v1.ServerSearchRes, err error) {
	search, err := service.Server().ServerSearch(ctx, req.ServerName)
	if err != nil {
		return nil, err
	}
	return search, nil
}
