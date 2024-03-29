// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package qiniu

import (
	"context"

	"voichatter/api/qiniu/v1"
)

type IQiniuV1 interface {
	UploadFile(ctx context.Context, req *v1.UploadFileReq) (res *v1.UploadFileRes, err error)
}
