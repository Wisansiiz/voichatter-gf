// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package message

import (
	"context"

	"voichatter/api/message/v1"
)

type IMessageV1 interface {
	MessageList(ctx context.Context, req *v1.MessageListReq) (res *v1.MessageListRes, err error)
}
