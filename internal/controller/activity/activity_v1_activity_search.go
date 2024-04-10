package activity

import (
	"context"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/activity/v1"
)

func (c *ControllerV1) ActivitySearch(ctx context.Context, req *v1.ActivitySearchReq) (res *v1.ActivitySearchRes, err error) {
	search, err := service.Activity().ActivitySearch(ctx, model.ActivitySearchInput{
		ActivityTitle: req.ActivityTitle,
	})
	if err != nil {
		return nil, err
	}
	return &v1.ActivitySearchRes{
		Activities: search,
	}, nil
}
