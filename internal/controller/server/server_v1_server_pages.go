package server

import (
	"context"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/server/v1"
)

func (c *ControllerV1) ServerPages(ctx context.Context, req *v1.ServerPagesReq) (res *v1.ServerPagesRes, err error) {
	pages, total, err := service.Server().ServerPages(ctx, model.ServerPagesInput{
		ServerName: req.ServerName,
		Page:       req.Page,
		PageSize:   req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	return &v1.ServerPagesRes{
		Servers:   pages,
		PageTotal: total,
	}, nil
}
