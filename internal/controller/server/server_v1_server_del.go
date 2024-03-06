package server

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"voichatter/internal/service"

	"voichatter/api/server/v1"
)

func (c *ControllerV1) ServerDel(ctx context.Context, req *v1.ServerDelReq) (res *v1.ServerDelRes, err error) {
	_, err = service.Server().ServerDel(ctx, req.ServerId)
	if err != nil {
		return nil, gerror.New("权限不足")
	}
	return nil, nil
}
