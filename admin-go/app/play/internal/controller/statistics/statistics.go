package statistics

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/service"
)

var Statistics = cStatistics{}

type cStatistics struct{}

// Overview 总览统计
func (c *cStatistics) Overview(ctx context.Context, req *v1.StatisticsOverviewReq) (res *v1.StatisticsOverviewRes, err error) {
	res, err = service.Statistics().Overview(ctx)
	return
}
