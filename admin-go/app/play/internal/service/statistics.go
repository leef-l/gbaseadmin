package service

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
)

type IStatistics interface {
	Overview(ctx context.Context) (out *v1.StatisticsOverviewRes, err error)
}

var localStatistics IStatistics

func Statistics() IStatistics {
	return localStatistics
}

func RegisterStatistics(i IStatistics) {
	localStatistics = i
}
