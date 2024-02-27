package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"voichatter/api/user/v1"
)

func (c *ControllerV1) SignOut(ctx context.Context, req *v1.SignOutReq) (res *v1.SignOutRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
