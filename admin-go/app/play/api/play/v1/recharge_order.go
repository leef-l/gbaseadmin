package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// RechargeOrder API

// RechargeOrderCreateReq 创建充值订单表请求
type RechargeOrderCreateReq struct {
	g.Meta `path:"/recharge_order/create" method:"post" tags:"充值订单表" summary:"创建充值订单表"`
	OrderNo string `json:"orderNo" v:"required#充值订单号不能为空" dc:"充值订单号"`
	MemberID snowflake.JsonInt64 `json:"memberID" v:"required#会员ID不能为空" dc:"会员ID"`
	RechargePlanID snowflake.JsonInt64 `json:"rechargePlanID" v:"required#充值方案ID不能为空" dc:"充值方案ID"`
	Amount int64 `json:"amount" v:"required#充值金额（分）不能为空" dc:"充值金额（分）"`
	GiftAmount int64 `json:"giftAmount"  dc:"赠送金额（分）"`
	PayType int `json:"payType"  dc:"支付方式"`
	TradeNo string `json:"tradeNo"  dc:"第三方交易号"`
	PayStatus int `json:"payStatus"  dc:"支付状态"`
	PayAt *gtime.Time `json:"payAt"  dc:"支付时间"`
}

// RechargeOrderCreateRes 创建充值订单表响应
type RechargeOrderCreateRes struct {
	g.Meta `mime:"application/json"`
}

// RechargeOrderUpdateReq 更新充值订单表请求
type RechargeOrderUpdateReq struct {
	g.Meta `path:"/recharge_order/update" method:"put" tags:"充值订单表" summary:"更新充值订单表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"充值订单表ID"`
	OrderNo string `json:"orderNo" dc:"充值订单号"`
	MemberID snowflake.JsonInt64 `json:"memberID" dc:"会员ID"`
	RechargePlanID snowflake.JsonInt64 `json:"rechargePlanID" dc:"充值方案ID"`
	Amount int64 `json:"amount" dc:"充值金额（分）"`
	GiftAmount int64 `json:"giftAmount" dc:"赠送金额（分）"`
	PayType int `json:"payType" dc:"支付方式"`
	TradeNo string `json:"tradeNo" dc:"第三方交易号"`
	PayStatus int `json:"payStatus" dc:"支付状态"`
	PayAt *gtime.Time `json:"payAt" dc:"支付时间"`
}

// RechargeOrderUpdateRes 更新充值订单表响应
type RechargeOrderUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// RechargeOrderDeleteReq 删除充值订单表请求
type RechargeOrderDeleteReq struct {
	g.Meta `path:"/recharge_order/delete" method:"delete" tags:"充值订单表" summary:"删除充值订单表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"充值订单表ID"`
}

// RechargeOrderDeleteRes 删除充值订单表响应
type RechargeOrderDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// RechargeOrderDetailReq 获取充值订单表详情请求
type RechargeOrderDetailReq struct {
	g.Meta `path:"/recharge_order/detail" method:"get" tags:"充值订单表" summary:"获取充值订单表详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"充值订单表ID"`
}

// RechargeOrderDetailRes 获取充值订单表详情响应
type RechargeOrderDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.RechargeOrderDetailOutput
}

// RechargeOrderListReq 获取充值订单表列表请求
type RechargeOrderListReq struct {
	g.Meta   `path:"/recharge_order/list" method:"get" tags:"充值订单表" summary:"获取充值订单表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	PayType int `json:"payType" dc:"支付方式"`
	PayStatus int `json:"payStatus" dc:"支付状态"`
}

// RechargeOrderListRes 获取充值订单表列表响应
type RechargeOrderListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.RechargeOrderListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

