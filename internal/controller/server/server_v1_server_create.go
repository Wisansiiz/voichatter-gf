package server

import (
	"context"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/server/v1"
)

func (c *ControllerV1) ServerCreate(ctx context.Context, req *v1.ServerCreateReq) (res *v1.ServerCreateRes, err error) {
	server, err := service.Server().ServerCreate(ctx, model.ServerCreateInput{
		ServerName:   req.ServerName,
		ServerType:   req.ServerType,
		ServerImgUrl: req.ServerImgUrl,
	})
	if err != nil {
		return nil, err
	}
	return server, nil
}
