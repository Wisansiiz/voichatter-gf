// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package group

import (
	"context"

	"voichatter/api/group/v1"
)

type IGroupV1 interface {
	GroupList(ctx context.Context, req *v1.GroupListReq) (res *v1.GroupListRes, err error)
}
