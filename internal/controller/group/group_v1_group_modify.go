package group

import (
	"context"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/group/v1"
)

func (c *ControllerV1) GroupModify(ctx context.Context, req *v1.GroupModifyReq) (res *v1.GroupModifyRes, err error) {
	modify, err := service.Group().GroupModify(ctx, model.Group{
		GroupId:   req.GroupId,
		ServerId:  req.ServerId,
		GroupName: req.GroupName,
	})
	if err != nil {
		return nil, err
	}
	return modify, nil
}
