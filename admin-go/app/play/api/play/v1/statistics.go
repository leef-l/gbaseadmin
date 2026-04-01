package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// StatisticsOverviewReq 总览统计请求
type StatisticsOverviewReq struct {
	g.Meta `path:"/statistics/overview" method:"get" tags:"数据统计" summary:"总览统计"`
}

// StatisticsOverviewRes 总览统计响应
type StatisticsOverviewRes struct {
	TotalMembers  int64 `json:"totalMembers"`
	TotalCoaches  int64 `json:"totalCoaches"`
	TotalOrders   int64 `json:"totalOrders"`
	TotalRevenue  int64 `json:"totalRevenue" dc:"总营收(分)"`
	TodayOrders   int64 `json:"todayOrders"`
	TodayRevenue  int64 `json:"todayRevenue"`
	TodayNewUsers int64 `json:"todayNewUsers"`
}
