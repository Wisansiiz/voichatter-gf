// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package channel

import (
	"context"

	"voichatter/api/channel/v1"
)

type IChannelV1 interface {
	ChannelCreate(ctx context.Context, req *v1.ChannelCreateReq) (res *v1.ChannelCreateRes, err error)
	ChannelModify(ctx context.Context, req *v1.ChannelModifyReq) (res *v1.ChannelModifyRes, err error)
	ChannelRemove(ctx context.Context, req *v1.ChannelRemoveReq) (res *v1.ChannelRemoveRes, err error)
}
