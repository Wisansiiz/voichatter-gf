package group

import (
	"context"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/group/v1"
)

func (c *ControllerV1) GroupCreate(ctx context.Context, req *v1.GroupCreateReq) (res *v1.GroupCreateRes, err error) {
	groupCreate, err := service.Group().GroupCreate(ctx, model.GroupCreateInput{
		ServerId:  req.ServerId,
		GroupName: req.GroupName,
	})
	if err != nil {
		return nil, err
	}
	return groupCreate, nil
}
