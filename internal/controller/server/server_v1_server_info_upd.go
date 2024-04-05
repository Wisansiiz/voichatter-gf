package server

import (
	"context"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/server/v1"
)

func (c *ControllerV1) ServerInfoUpd(ctx context.Context, req *v1.ServerInfoUpdReq) (res *v1.ServerInfoUpdRes, err error) {
	upd, err := service.Server().ServerInfoUpd(ctx, model.ServerInfoUpdInput{
		ServerId:          req.ServerId,
		ServerName:        req.ServerName,
		ServerType:        req.ServerType,
		ServerDescription: req.ServerDescription,
	})
	if err != nil {
		return nil, err
	}
	return upd, nil
}
