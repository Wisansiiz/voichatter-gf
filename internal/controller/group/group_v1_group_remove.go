package group

import (
	"context"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/group/v1"
)

func (c *ControllerV1) GroupRemove(ctx context.Context, req *v1.GroupRemoveReq) (res *v1.GroupRemoveRes, err error) {
	_, err = service.Group().GroupRemove(ctx, model.GroupRemoveInput{
		GroupId:  req.GroupId,
		ServerId: req.ServerId,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
