package statistics

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/service"
)

func init() {
	service.RegisterStatistics(New())
}

func New() *sStatistics {
	return &sStatistics{}
}

type sStatistics struct{}

// Overview 总览统计
func (s *sStatistics) Overview(ctx context.Context) (out *v1.StatisticsOverviewRes, err error) {
	out = &v1.StatisticsOverviewRes{}
	today := gtime.Now().Format("Y-m-d")

	// 总会员数
	cnt, e := dao.PlayMember.Ctx(ctx).
		Where(dao.PlayMember.Columns().DeletedAt, nil).
		Count()
	if e != nil {
		err = e
		return
	}
	out.TotalMembers = int64(cnt)

	// 总陪玩师数
	cnt, e = dao.PlayCoach.Ctx(ctx).
		Where(dao.PlayCoach.Columns().DeletedAt, nil).
		Count()
	if e != nil {
		err = e
		return
	}
	out.TotalCoaches = int64(cnt)

	// 总订单数
	cnt, e = dao.PlayOrder.Ctx(ctx).
		Where(dao.PlayOrder.Columns().DeletedAt, nil).
		Count()
	if e != nil {
		err = e
		return
	}
	out.TotalOrders = int64(cnt)

	// 总营收（已支付/进行中/已完成的订单）
	totalRevenueVal, e := dao.PlayOrder.Ctx(ctx).
		Fields("IFNULL(SUM(pay_amount),0) AS total").
		WhereIn(dao.PlayOrder.Columns().OrderStatus, g.Slice{1, 2, 3}).
		Where(dao.PlayOrder.Columns().DeletedAt, nil).
		Value()
	if e != nil {
		err = e
		return
	}
	out.TotalRevenue = totalRevenueVal.Int64()

	// 今日订单数
	cnt, e = dao.PlayOrder.Ctx(ctx).
		Where(dao.PlayOrder.Columns().DeletedAt, nil).
		WhereGTE(dao.PlayOrder.Columns().CreatedAt, today).
		Count()
	if e != nil {
		err = e
		return
	}
	out.TodayOrders = int64(cnt)

	// 今日营收
	todayRevenueVal, e := dao.PlayOrder.Ctx(ctx).
		Fields("IFNULL(SUM(pay_amount),0) AS total").
		WhereIn(dao.PlayOrder.Columns().OrderStatus, g.Slice{1, 2, 3}).
		Where(dao.PlayOrder.Columns().DeletedAt, nil).
		WhereGTE(dao.PlayOrder.Columns().CreatedAt, today).
		Value()
	if e != nil {
		err = e
		return
	}
	out.TodayRevenue = todayRevenueVal.Int64()

	// 今日新增用户
	cnt, e = dao.PlayMember.Ctx(ctx).
		Where(dao.PlayMember.Columns().DeletedAt, nil).
		WhereGTE(dao.PlayMember.Columns().CreatedAt, today).
		Count()
	if e != nil {
		err = e
		return
	}
	out.TodayNewUsers = int64(cnt)

	return
}
