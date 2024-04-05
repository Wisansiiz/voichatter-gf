package server

import (
	"context"
	"voichatter/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"voichatter/api/server/v1"
)

func (c *ControllerV1) ServerInfo(ctx context.Context, req *v1.ServerInfoReq) (res *v1.ServerInfoRes, err error) {
	info, err := service.Server().ServerInfo(ctx, req.ServerId)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取服务器信息失败")
	}
	return info, nil
}
