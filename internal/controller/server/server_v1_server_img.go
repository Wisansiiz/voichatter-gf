package server

import (
	"context"
	"voichatter/internal/service"

	"voichatter/api/server/v1"
)

func (c *ControllerV1) ServerImg(ctx context.Context, req *v1.ServerImgReq) (res *v1.ServerImgRes, err error) {
	res, err = service.Server().ServerImg(ctx, req.ServerId, req.File)
	if err != nil {
		return nil, err
	}
	return res, nil
}
