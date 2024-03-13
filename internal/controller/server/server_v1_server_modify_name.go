package server

import (
	"context"
	"voichatter/internal/service"

	"voichatter/api/server/v1"
)

func (c *ControllerV1) ServerModifyName(ctx context.Context, req *v1.ServerModifyNameReq) (res *v1.ServerModifyNameRes, err error) {
	data, err := service.Server().ServerModifyName(ctx, req.ServerId, req.ServerName)
	if err != nil {
		return nil, err
	}
	return data, nil
}
