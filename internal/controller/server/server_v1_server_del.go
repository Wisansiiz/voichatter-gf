package server

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"voichatter/api/server/v1"
)

func (c *ControllerV1) ServerDel(ctx context.Context, req *v1.ServerDelReq) (res *v1.ServerDelRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
