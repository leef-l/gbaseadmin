package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// Order API

// OrderCreateReq 创建订单表请求
type OrderCreateReq struct {
	g.Meta `path:"/order/create" method:"post" tags:"订单表" summary:"创建订单表"`
	OrderNo string `json:"orderNo" v:"required#订单编号不能为空" dc:"订单编号"`
	MemberID snowflake.JsonInt64 `json:"memberID" v:"required#下单会员ID不能为空" dc:"下单会员ID"`
	CoachID snowflake.JsonInt64 `json:"coachID" v:"required#陪玩师ID不能为空" dc:"陪玩师ID"`
	ShopID snowflake.JsonInt64 `json:"shopID"  dc:"店铺ID（0表示无店铺）"`
	GoodsID snowflake.JsonInt64 `json:"goodsID" v:"required#商品ID不能为空" dc:"商品ID"`
	GoodsTitle string `json:"goodsTitle" v:"required#商品名称（冗余）不能为空" dc:"商品名称（冗余）"`
	GoodsPrice int64 `json:"goodsPrice" v:"required#商品单价（分，下单时快照）不能为空" dc:"商品单价（分，下单时快照）"`
	Quantity int `json:"quantity"  dc:"数量"`
	TotalAmount int64 `json:"totalAmount"  dc:"订单总额（分）"`
	DiscountAmount int64 `json:"discountAmount"  dc:"会员折扣金额（分）"`
	CouponAmount int64 `json:"couponAmount"  dc:"优惠券抵扣金额（分）"`
	PayAmount int64 `json:"payAmount"  dc:"实付金额（分）"`
	CouponMemberID snowflake.JsonInt64 `json:"couponMemberID"  dc:"使用的优惠券领取记录ID"`
	PayType int `json:"payType"  dc:"支付方式"`
	OrderStatus int `json:"orderStatus"  dc:"订单状态"`
	PayAt *gtime.Time `json:"payAt"  dc:"支付时间"`
	StartAt *gtime.Time `json:"startAt"  dc:"服务开始时间"`
	FinishAt *gtime.Time `json:"finishAt"  dc:"服务完成时间"`
	CancelAt *gtime.Time `json:"cancelAt"  dc:"取消时间"`
	CancelReason string `json:"cancelReason"  dc:"取消原因"`
	Remark string `json:"remark"  dc:"订单备注"`
}

// OrderCreateRes 创建订单表响应
type OrderCreateRes struct {
	g.Meta `mime:"application/json"`
}

// OrderUpdateReq 更新订单表请求
type OrderUpdateReq struct {
	g.Meta `path:"/order/update" method:"put" tags:"订单表" summary:"更新订单表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"订单表ID"`
	OrderNo string `json:"orderNo" dc:"订单编号"`
	MemberID snowflake.JsonInt64 `json:"memberID" dc:"下单会员ID"`
	CoachID snowflake.JsonInt64 `json:"coachID" dc:"陪玩师ID"`
	ShopID snowflake.JsonInt64 `json:"shopID" dc:"店铺ID（0表示无店铺）"`
	GoodsID snowflake.JsonInt64 `json:"goodsID" dc:"商品ID"`
	GoodsTitle string `json:"goodsTitle" dc:"商品名称（冗余）"`
	GoodsPrice int64 `json:"goodsPrice" dc:"商品单价（分，下单时快照）"`
	Quantity int `json:"quantity" dc:"数量"`
	TotalAmount int64 `json:"totalAmount" dc:"订单总额（分）"`
	DiscountAmount int64 `json:"discountAmount" dc:"会员折扣金额（分）"`
	CouponAmount int64 `json:"couponAmount" dc:"优惠券抵扣金额（分）"`
	PayAmount int64 `json:"payAmount" dc:"实付金额（分）"`
	CouponMemberID snowflake.JsonInt64 `json:"couponMemberID" dc:"使用的优惠券领取记录ID"`
	PayType int `json:"payType" dc:"支付方式"`
	OrderStatus int `json:"orderStatus" dc:"订单状态"`
	PayAt *gtime.Time `json:"payAt" dc:"支付时间"`
	StartAt *gtime.Time `json:"startAt" dc:"服务开始时间"`
	FinishAt *gtime.Time `json:"finishAt" dc:"服务完成时间"`
	CancelAt *gtime.Time `json:"cancelAt" dc:"取消时间"`
	CancelReason string `json:"cancelReason" dc:"取消原因"`
	Remark string `json:"remark" dc:"订单备注"`
}

// OrderUpdateRes 更新订单表响应
type OrderUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// OrderDeleteReq 删除订单表请求
type OrderDeleteReq struct {
	g.Meta `path:"/order/delete" method:"delete" tags:"订单表" summary:"删除订单表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"订单表ID"`
}

// OrderDeleteRes 删除订单表响应
type OrderDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// OrderDetailReq 获取订单表详情请求
type OrderDetailReq struct {
	g.Meta `path:"/order/detail" method:"get" tags:"订单表" summary:"获取订单表详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"订单表ID"`
}

// OrderDetailRes 获取订单表详情响应
type OrderDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.OrderDetailOutput
}

// OrderListReq 获取订单表列表请求
type OrderListReq struct {
	g.Meta   `path:"/order/list" method:"get" tags:"订单表" summary:"获取订单表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	PayType int `json:"payType" dc:"支付方式"`
	OrderStatus int `json:"orderStatus" dc:"订单状态"`
}

// OrderListRes 获取订单表列表响应
type OrderListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.OrderListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

