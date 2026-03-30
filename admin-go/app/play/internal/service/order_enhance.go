package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
)

type IOrderEnhance interface {
	ChangeStatus(ctx context.Context, in *model.OrderChangeStatusInput) error
}
