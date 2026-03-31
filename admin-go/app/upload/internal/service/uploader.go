package service

import "context"

type IUploader interface {
	Upload(ctx context.Context) error
}

var localUploader IUploader

func RegisterUploader(s IUploader) {
	localUploader = s
}

func Uploader() IUploader {
	return localUploader
}
