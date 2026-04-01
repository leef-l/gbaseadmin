package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type IMessage interface {
	Create(ctx context.Context, in *model.MessageCreateInput) error
	Update(ctx context.Context, in *model.MessageUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	BatchDelete(ctx context.Context, ids []snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.MessageDetailOutput, err error)
	List(ctx context.Context, in *model.MessageListInput) (list []*model.MessageListOutput, total int, err error)
	Export(ctx context.Context, in *model.MessageListInput) (list []*model.MessageListOutput, err error)
}

var localMessage IMessage

func Message() IMessage {
	return localMessage
}

func RegisterMessage(i IMessage) {
	localMessage = i
}
