package main

import (
	_ "gbaseadmin/app/play/internal/packed"

	_ "gbaseadmin/app/play/internal/logic/activity"
	_ "gbaseadmin/app/play/internal/logic/activity_join"
	_ "gbaseadmin/app/play/internal/logic/activity_reward"
	_ "gbaseadmin/app/play/internal/logic/activity_step"
	_ "gbaseadmin/app/play/internal/logic/activity_step_log"
	_ "gbaseadmin/app/play/internal/logic/balance_log"
	_ "gbaseadmin/app/play/internal/logic/category"
	_ "gbaseadmin/app/play/internal/logic/coach"
	_ "gbaseadmin/app/play/internal/logic/coach_apply"
	_ "gbaseadmin/app/play/internal/logic/coach_level"
	_ "gbaseadmin/app/play/internal/logic/coupon"
	_ "gbaseadmin/app/play/internal/logic/coupon_member"
	_ "gbaseadmin/app/play/internal/logic/goods"
	_ "gbaseadmin/app/play/internal/logic/member"
	_ "gbaseadmin/app/play/internal/logic/member_level"
	_ "gbaseadmin/app/play/internal/logic/oauth"
	_ "gbaseadmin/app/play/internal/logic/order"
	_ "gbaseadmin/app/play/internal/logic/payment"
	_ "gbaseadmin/app/play/internal/logic/playapi"
	_ "gbaseadmin/app/play/internal/logic/profit_log"
	_ "gbaseadmin/app/play/internal/logic/recharge_order"
	_ "gbaseadmin/app/play/internal/logic/recharge_plan"
	_ "gbaseadmin/app/play/internal/logic/review"
	_ "gbaseadmin/app/play/internal/logic/shop"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"gbaseadmin/app/play/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
