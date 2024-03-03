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
}
