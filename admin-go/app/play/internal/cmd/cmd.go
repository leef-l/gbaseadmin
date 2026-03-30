package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gbaseadmin/app/play/internal/controller/activity"
	"gbaseadmin/app/play/internal/controller/activity_join"
	"gbaseadmin/app/play/internal/controller/activity_reward"
	"gbaseadmin/app/play/internal/controller/activity_step"
	"gbaseadmin/app/play/internal/controller/balance_log"
	"gbaseadmin/app/play/internal/controller/category"
	"gbaseadmin/app/play/internal/controller/coach"
	"gbaseadmin/app/play/internal/controller/coach_apply"
	"gbaseadmin/app/play/internal/controller/coach_level"
	"gbaseadmin/app/play/internal/controller/coupon"
	"gbaseadmin/app/play/internal/controller/coupon_member"
	"gbaseadmin/app/play/internal/controller/goods"
	"gbaseadmin/app/play/internal/controller/member"
	"gbaseadmin/app/play/internal/controller/member_level"
	"gbaseadmin/app/play/internal/controller/oauth"
	"gbaseadmin/app/play/internal/controller/order"
	"gbaseadmin/app/play/internal/controller/payment"
	"gbaseadmin/app/play/internal/controller/profit_log"
	"gbaseadmin/app/play/internal/controller/recharge_order"
	"gbaseadmin/app/play/internal/controller/recharge_plan"
	"gbaseadmin/app/play/internal/controller/review"
	"gbaseadmin/app/play/internal/controller/shop"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start play http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				// 后台管理路由组
				group.Group("/api/play", func(group *ghttp.RouterGroup) {
					group.Bind(
						member_level.MemberLevel,
						coach_level.CoachLevel,
						shop.Shop,
						recharge_plan.RechargePlan,
						member.Member,
						coach_apply.CoachApply,
						coach.Coach,
						category.Category,
						goods.Goods,
						order.Order,
						payment.Payment,
						recharge_order.RechargeOrder,
						balance_log.BalanceLog,
						activity.Activity,
						activity_reward.ActivityReward,
						activity_step.ActivityStep,
						activity_join.ActivityJoin,
						coupon.Coupon,
						coupon_member.CouponMember,
						oauth.Oauth,
						review.Review,
						profit_log.ProfitLog,
					)
				})
				// C端API路由组
				group.Group("/api/playapi", func(group *ghttp.RouterGroup) {
					group.Bind(
						category.Category,
						goods.Goods,
						coach.Coach,
						review.Review,
						shop.Shop,
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
