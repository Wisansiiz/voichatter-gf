// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package activity

import (
	"context"

	"voichatter/api/activity/v1"
)

type IActivityV1 interface {
	ActivityCreate(ctx context.Context, req *v1.ActivityCreateReq) (res *v1.ActivityCreateRes, err error)
	ActivityPages(ctx context.Context, req *v1.ActivityPagesReq) (res *v1.ActivityPagesRes, err error)
	ActivitySearch(ctx context.Context, req *v1.ActivitySearchReq) (res *v1.ActivitySearchRes, err error)
}
