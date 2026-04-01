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
	"gbaseadmin/app/play/internal/controller/activity_step_log"
	"gbaseadmin/app/play/internal/controller/balance_log"
	"gbaseadmin/app/play/internal/controller/banner"
	"gbaseadmin/app/play/internal/controller/category"
	"gbaseadmin/app/play/internal/controller/coach"
	"gbaseadmin/app/play/internal/controller/coach_apply"
	"gbaseadmin/app/play/internal/controller/coach_level"
	"gbaseadmin/app/play/internal/controller/coupon"
	"gbaseadmin/app/play/internal/controller/coupon_member"
	"gbaseadmin/app/play/internal/controller/goods"
	"gbaseadmin/app/play/internal/controller/member"
	"gbaseadmin/app/play/internal/controller/member_level"
	"gbaseadmin/app/play/internal/controller/message"
	"gbaseadmin/app/play/internal/controller/oauth"
	"gbaseadmin/app/play/internal/controller/order"
	"gbaseadmin/app/play/internal/controller/payment"
	playapiCtrl "gbaseadmin/app/play/internal/controller/playapi"
	"gbaseadmin/app/play/internal/controller/profit_log"
	"gbaseadmin/app/play/internal/controller/recharge_order"
	"gbaseadmin/app/play/internal/controller/recharge_plan"
	"gbaseadmin/app/play/internal/controller/review"
	"gbaseadmin/app/play/internal/controller/shop"
	"gbaseadmin/app/play/internal/controller/statistics"
	"gbaseadmin/app/play/internal/controller/withdraw"

	"gbaseadmin/app/play/internal/middleware"
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
					group.Middleware(middleware.Auth)
					group.Bind(
						activity.Activity,
						activity_join.ActivityJoin,
						activity_reward.ActivityReward,
						activity_step.ActivityStep,
						activity_step_log.ActivityStepLog,
						balance_log.BalanceLog,
						banner.Banner,
						category.Category,
						coach.Coach,
						coach_apply.CoachApply,
						coach_level.CoachLevel,
						coupon.Coupon,
						coupon_member.CouponMember,
						goods.Goods,
						member.Member,
						member_level.MemberLevel,
						message.Message,
						oauth.Oauth,
						order.Order,
						payment.Payment,
						profit_log.ProfitLog,
						recharge_order.RechargeOrder,
						recharge_plan.RechargePlan,
						review.Review,
						shop.Shop,
						statistics.Statistics,
						withdraw.Withdraw,
					)
				})
				// C端API路由组
				group.Group("/api/playapi", func(group *ghttp.RouterGroup) {
					// 公开接口（无需登录，但有 token 时解析用户信息）
					group.Middleware(middleware.MemberAuthOptional)
					group.Bind(
						playapiCtrl.Auth,
						playapiCtrl.GoodsPublic,
						playapiCtrl.CoachPublic,
						playapiCtrl.ReviewPublic,
						playapiCtrl.ActivityPublic,
						playapiCtrl.CouponPublic,
						playapiCtrl.SearchPublic,
						playapiCtrl.PaymentNotify,
						playapiCtrl.RechargeNotify,
						playapiCtrl.BannerPublic,
					)
					// 需要会员登录
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(middleware.MemberAuth)
						group.Bind(
							playapiCtrl.Member,
							playapiCtrl.Order,
							playapiCtrl.Payment,
							playapiCtrl.Coupon,
							playapiCtrl.Activity,
							playapiCtrl.Recharge,
							playapiCtrl.Review,
							playapiCtrl.CoachApply,
							playapiCtrl.Message,
						)
						// 需要陪玩师身份
						group.Group("/", func(group *ghttp.RouterGroup) {
							group.Middleware(middleware.CoachOnly)
							group.Bind(
								playapiCtrl.CoachWork,
								playapiCtrl.OrderCoach,
								playapiCtrl.ReviewCoach,
								playapiCtrl.Withdraw,
							)
						})
					})
				})
			})
			s.Run()
			return nil
		},
	}
)
