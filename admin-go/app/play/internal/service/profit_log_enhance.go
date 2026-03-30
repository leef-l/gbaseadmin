package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
)

type IProfitLogEnhance interface {
	SettleOrder(ctx context.Context, in *model.SettleOrderInput) error
}

func ProfitLogEnhance() IProfitLogEnhance {
	return localProfitLog.(IProfitLogEnhance)
}
