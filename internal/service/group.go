// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "voichatter/api/group/v1"
	"voichatter/internal/model"
)

type (
	IGroup interface {
		GroupList(ctx context.Context, serverId uint64) (res *v1.GroupListRes, err error)
		GroupCreate(ctx context.Context, in model.GroupCreateInput) (res *v1.GroupCreateRes, err error)
		GroupModify(ctx context.Context, req *v1.GroupModifyReq) (res *v1.GroupModifyRes, err error)
		GroupRemove(ctx context.Context, req *v1.GroupRemoveReq) (res *v1.GroupRemoveRes, err error)
	}
)

var (
	localGroup IGroup
)

func Group() IGroup {
	if localGroup == nil {
		panic("implement not found for interface IGroup, forgot register?")
	}
	return localGroup
}

func RegisterGroup(i IGroup) {
	localGroup = i
}
