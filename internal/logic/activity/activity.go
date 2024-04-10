package activity

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"voichatter/internal/dao"
	"voichatter/internal/model"
	"voichatter/internal/service"
	"voichatter/utility/auth"
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
		Where("activity_title like ?", "%"+in.ActivityTitle+"%").
		Page(in.Page, in.PageSize).
		Scan(&res); err != nil {
		return nil, 0, errResponse.DbOperationErrorDefault()
	}
	count, err := dao.Activity.Ctx(ctx).
		Where("activity_title like ?", "%"+in.ActivityTitle+"%").
		Count()
	if err != nil {
		return nil, 0, errResponse.DbOperationErrorDefault()
	}
	return res, count, nil
}

func (s *sActivity) ActivityCreate(ctx context.Context, in model.ActivityCreateInput) (res *model.Activity, err error) {
	userId := gconv.Uint64(ctx.Value("userId"))
	in.CreatorUserId = userId
	id, err := dao.Activity.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return nil, errResponse.DbOperationError("创建失败")
	}
	if err = dao.Activity.Ctx(ctx).
		Where("activity_id = ?", id).
		Scan(&res); err != nil {
		return nil, errResponse.DbOperationErrorDefault()
	}
	return res, nil
}

func (s *sActivity) ActivityDelete(ctx context.Context, in model.ActivityDeleteInput) (err error) {
	if err = auth.IsActivityCreator(ctx, in.ActivityId); err != nil {
		return err
	}
	if _, err = dao.Activity.Ctx(ctx).
		Delete("activity_id = ?", in.ActivityId); err != nil {
		return errResponse.DbOperationError("删除失败")
	}
	return nil
}

func (s *sActivity) ActivityUpdate(ctx context.Context, in model.ActivityUpdateInput) (err error) {
	if err = auth.IsActivityCreator(ctx, in.ActivityId); err != nil {
		return err
	}
	if _, err = dao.Activity.Ctx(ctx).Update(in, in); err != nil {
		return errResponse.DbOperationErrorDefault()
	}
	return nil
}

func (s *sActivity) ActivitySearch(ctx context.Context, in model.ActivitySearchInput) (res []*model.Activity, err error) {
	if err = dao.Activity.Ctx(ctx).
		Where("activity_title like ?", "%"+in.ActivityTitle+"%").
		Scan(&res); err != nil {
		return nil, errResponse.DbOperationErrorDefault()
	}
	return res, nil
}
