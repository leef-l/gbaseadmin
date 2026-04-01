// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayOrder is the golang structure of table play_order for DAO operations like Where/Data.
type PlayOrder struct {
	g.Meta         `orm:"table:play_order, do:true"`
	Id             any         // 订单ID（Snowflake）
	OrderNo        any         // 订单编号
	MemberId       any         // 下单会员ID
	CoachId        any         // 陪玩师ID
	ShopId         any         // 店铺ID（0表示无店铺）
	GoodsId        any         // 商品ID
	GoodsTitle     any         // 商品名称（冗余）
	GoodsPrice     any         // 商品单价（分，下单时快照）
	Quantity       any         // 数量
	TotalAmount    any         // 订单总额（分）
	DiscountAmount any         // 会员折扣金额（分）
	CouponAmount   any         // 优惠券抵扣金额（分）
	PayAmount      any         // 实付金额（分）
	CouponMemberId any         // 使用的优惠券领取记录ID
	PayType        any         // 支付方式:0=未支付,1=微信支付,2=支付宝支付,3=余额支付
	OrderStatus    any         // 订单状态:0=待支付,1=已支付,2=进行中,3=已完成,4=已取消,5=退款中,6=已退款
	PayAt          *gtime.Time // 支付时间
	StartAt        *gtime.Time // 服务开始时间
	FinishAt       *gtime.Time // 服务完成时间
	CancelAt       *gtime.Time // 取消时间
	CancelReason   any         // 取消原因
	Remark         any         // 订单备注
	CreatedBy      any         // 创建人ID
	DeptId         any         // 所属部门ID
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 软删除时间
}
