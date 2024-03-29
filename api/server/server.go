// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package server

import (
	"context"

	"voichatter/api/server/v1"
)

type IServerV1 interface {
	ServerCreate(ctx context.Context, req *v1.ServerCreateReq) (res *v1.ServerCreateRes, err error)
	ServerDel(ctx context.Context, req *v1.ServerDelReq) (res *v1.ServerDelRes, err error)
	ServerJoin(ctx context.Context, req *v1.ServerJoinReq) (res *v1.ServerJoinRes, err error)
	ServerList(ctx context.Context, req *v1.ServerListReq) (res *v1.ServerListRes, err error)
	ServerModifyName(ctx context.Context, req *v1.ServerModifyNameReq) (res *v1.ServerModifyNameRes, err error)
}
