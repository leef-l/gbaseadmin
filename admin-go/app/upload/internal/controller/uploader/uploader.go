package uploader

import (
	"context"

	v1 "gbaseadmin/app/upload/api/upload/v1"
	"gbaseadmin/app/upload/internal/service"
)

var Uploader = cUploader{}

type cUploader struct{}

func (c *cUploader) Upload(ctx context.Context, req *v1.UploaderUploadReq) (res *v1.UploaderUploadRes, err error) {
	err = service.Uploader().Upload(ctx)
	return
}
