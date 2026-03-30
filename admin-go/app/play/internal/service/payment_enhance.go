package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
)

type IPaymentEnhance interface {
	BalancePay(ctx context.Context, in *model.BalancePayInput) error
}

func PaymentEnhance() IPaymentEnhance {
	return localPayment.(IPaymentEnhance)
}
