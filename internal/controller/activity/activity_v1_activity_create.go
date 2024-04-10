package activity

import (
	"context"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/activity/v1"
)

func (c *ControllerV1) ActivityCreate(ctx context.Context, req *v1.ActivityCreateReq) (res *v1.ActivityCreateRes, err error) {
	activity, err := service.Activity().ActivityCreate(ctx, model.ActivityCreateInput{
		ServerId:      req.ServerId,
		ActivityTitle: req.ActivityTitle,
		ActivityDesc:  req.ActivityDesc,
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
	})
	if err != nil {
		return nil, err
	}
	return &v1.ActivityCreateRes{
		Activities: activity,
	}, nil
}
