package activity

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"voichatter/internal/dao"
	"voichatter/internal/model"
	"voichatter/internal/service"
	"voichatter/utility/errResponse"
)

type (
	sActivity struct{}
)

func init() {
	service.RegisterActivity(New())
}

func New() service.IActivity {
	return &sActivity{}
}

func (s *sActivity) ActivityPages(ctx context.Context, in model.ActivityPagesInput) (res []*model.ActivityPages, total int, err error) {
	if err = dao.Activity.Ctx(ctx).
		Where("end_date > ?", gtime.Now()).
		Page(in.Page, in.PageSize).
		Scan(&res); err != nil {
		return nil, 0, errResponse.DbOperationErrorDefault()
	}
	count, err := dao.Activity.Ctx(ctx).Count()
	if err != nil {
		return nil, 0, errResponse.DbOperationErrorDefault()
	}
	return res, count, nil
}
