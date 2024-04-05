// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package server

import (
	"context"

	"voichatter/api/server/v1"
)

type IServerV1 interface {
	ServerCount(ctx context.Context, req *v1.ServerCountReq) (res *v1.ServerCountRes, err error)
	ServerCreate(ctx context.Context, req *v1.ServerCreateReq) (res *v1.ServerCreateRes, err error)
	ServerDel(ctx context.Context, req *v1.ServerDelReq) (res *v1.ServerDelRes, err error)
	ServerImg(ctx context.Context, req *v1.ServerImgReq) (res *v1.ServerImgRes, err error)
	ServerInfo(ctx context.Context, req *v1.ServerInfoReq) (res *v1.ServerInfoRes, err error)
	ServerInfoUpd(ctx context.Context, req *v1.ServerInfoUpdReq) (res *v1.ServerInfoUpdRes, err error)
	ServerJoin(ctx context.Context, req *v1.ServerJoinReq) (res *v1.ServerJoinRes, err error)
	ServerList(ctx context.Context, req *v1.ServerListReq) (res *v1.ServerListRes, err error)
	ServerModifyName(ctx context.Context, req *v1.ServerModifyNameReq) (res *v1.ServerModifyNameRes, err error)
	ServerSearch(ctx context.Context, req *v1.ServerSearchReq) (res *v1.ServerSearchRes, err error)
}
