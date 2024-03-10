package group

import (
	"context"
	"voichatter/internal/service"

	"voichatter/api/group/v1"
)

func (c *ControllerV1) GroupList(ctx context.Context, req *v1.GroupListReq) (res *v1.GroupListRes, err error) {
	ans, err := service.Group().GroupList(ctx, req.ServerId)
	if err != nil {
		return nil, err
	}
	return ans, nil
}
