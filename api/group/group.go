// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package group

import (
	"context"

	"voichatter/api/group/v1"
)

type IGroupV1 interface {
	GroupCreate(ctx context.Context, req *v1.GroupCreateReq) (res *v1.GroupCreateRes, err error)
	GroupList(ctx context.Context, req *v1.GroupListReq) (res *v1.GroupListRes, err error)
	GroupModify(ctx context.Context, req *v1.GroupModifyReq) (res *v1.GroupModifyRes, err error)
	GroupRemove(ctx context.Context, req *v1.GroupRemoveReq) (res *v1.GroupRemoveRes, err error)
}
