package qiniu

import (
	"context"
	"voichatter/internal/service"

	"voichatter/api/qiniu/v1"
)

func (c *ControllerV1) UploadFile(ctx context.Context, req *v1.UploadFileReq) (res *v1.UploadFileRes, err error) {
	url, err := service.Qiniu().UploadFile(ctx, req.File)
	if err != nil {
		return nil, err
	}
	return url, nil
}
