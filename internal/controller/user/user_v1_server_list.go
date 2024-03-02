package user

import (
	"context"
	"voichatter/internal/service"

	"voichatter/api/user/v1"
)

func (c *ControllerV1) ServerList(ctx context.Context, req *v1.ServerListReq) (res *v1.ServerListRes, err error) {
	list, err := service.User().ServerList(ctx, req)
	if err != nil {
		return nil, err
	}
	return list, nil
}
