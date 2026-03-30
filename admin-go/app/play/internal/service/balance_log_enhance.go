package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
)

type IBalanceLogEnhance interface {
	AddLog(ctx context.Context, in *model.AddBalanceLogInput) error
}

func BalanceLogEnhance() IBalanceLogEnhance {
	return localBalanceLog.(IBalanceLogEnhance)
}
