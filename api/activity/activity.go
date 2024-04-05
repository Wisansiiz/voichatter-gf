// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package activity

import (
	"context"

	"voichatter/api/activity/v1"
)

type IActivityV1 interface {
	ActivityPages(ctx context.Context, req *v1.ActivityPagesReq) (res *v1.ActivityPagesRes, err error)
}
