package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type IPayment interface {
	Create(ctx context.Context, in *model.PaymentCreateInput) error
	Update(ctx context.Context, in *model.PaymentUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.PaymentDetailOutput, err error)
	List(ctx context.Context, in *model.PaymentListInput) (list []*model.PaymentListOutput, total int, err error)
}

var localPayment IPayment

func Payment() IPayment {
	return localPayment
}

func RegisterPayment(i IPayment) {
	localPayment = i
}
