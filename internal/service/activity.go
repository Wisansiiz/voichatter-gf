// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"voichatter/internal/model"
)

type (
	IActivity interface {
		ActivityPages(ctx context.Context, in model.ActivityPagesInput) (res []*model.ActivityPages, total int, err error)
		ActivityCreate(ctx context.Context, in model.ActivityCreateInput) (res *model.Activity, err error)
		ActivityDelete(ctx context.Context, in model.ActivityDeleteInput) (err error)
		ActivityUpdate(ctx context.Context, in model.ActivityUpdateInput) (err error)
		ActivitySearch(ctx context.Context, in model.ActivitySearchInput) (res []*model.Activity, err error)
	}
)

var (
	localActivity IActivity
)

func Activity() IActivity {
	if localActivity == nil {
		panic("implement not found for interface IActivity, forgot register?")
	}
	return localActivity
}

func RegisterActivity(i IActivity) {
	localActivity = i
}
