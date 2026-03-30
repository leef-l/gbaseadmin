package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gbaseadmin/utility/snowflake"
)

// PayByBalanceReq 余额支付请求
type PayByBalanceReq struct {
	g.Meta   `path:"/payment/pay_by_balance" method:"post" tags:"支付记录" summary:"余额支付"`
	OrderID  snowflake.JsonInt64 `json:"orderID" v:"required#订单ID不能为空" dc:"订单ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" v:"required#会员ID不能为空" dc:"会员ID"`
}

// PayByBalanceRes 余额支付响应
type PayByBalanceRes struct {
	g.Meta `mime:"application/json"`
}
