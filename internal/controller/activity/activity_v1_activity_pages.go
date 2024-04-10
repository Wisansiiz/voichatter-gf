package activity

import (
	"context"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/activity/v1"
)

func (c *ControllerV1) ActivityPages(ctx context.Context, req *v1.ActivityPagesReq) (res *v1.ActivityPagesRes, err error) {
	pages, total, err := service.Activity().ActivityPages(ctx, model.ActivityPagesInput{
		Page:          req.Page,
		PageSize:      req.PageSize,
		ActivityTitle: req.ActivityTitle,
	})
	if err != nil {
		return nil, err
	}
	return &v1.ActivityPagesRes{
		ActivityPages: pages,
		PageTotal:     total,
	}, nil
}
