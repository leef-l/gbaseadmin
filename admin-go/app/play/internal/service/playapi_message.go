package service

import (
	"context"

	v1 "gbaseadmin/api/playapi/v1"
)

type IPlayapiMessage interface {
	List(ctx context.Context, memberID int64, msgType string, page, pageSize int) (list []v1.MessageItem, total int, err error)
	Read(ctx context.Context, memberID int64, id string) error
	ReadAll(ctx context.Context, memberID int64) error
}

var localPlayapiMessage IPlayapiMessage

func PlayapiMessage() IPlayapiMessage {
	return localPlayapiMessage
}

func RegisterPlayapiMessage(i IPlayapiMessage) {
	localPlayapiMessage = i
}
