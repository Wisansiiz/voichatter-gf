package upload

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

func FileUpload(ctx context.Context, r *ghttp.Request) (url string, err error) {
	fileUpload := r.GetUploadFile("file")
	if fileUpload == nil {
		return "", gerror.New("无文件")
	}

	dirPath := "./upload/"
	key, err := fileUpload.Save(dirPath, true)
	if err != nil {
		return
	}
	localFile := dirPath + key
	accessKey := g.Cfg().MustGet(ctx, "qiniu.accessKey").String()
	secretKey := g.Cfg().MustGet(ctx, "qiniu.secretKey").String()
	mac := auth.New(accessKey, secretKey)
	bucket := g.Cfg().MustGet(ctx, "qiniu.bucket").String()
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseHTTPS:      true,
		UseCdnDomains: false,
	}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	err = formUploader.PutFile(ctx, &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		return
	}
	url = g.Cfg().MustGet(ctx, "qiniu.url").String() + ret.Key
	g.Dump(url)
	return
}
