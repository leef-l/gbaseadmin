// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayOrder is the golang structure for table play_order.
type PlayOrder struct {
	Id             uint64      `orm:"id"               description:"订单ID（Snowflake）"`                                // 订单ID（Snowflake）
	OrderNo        string      `orm:"order_no"         description:"订单编号"`                                           // 订单编号
	MemberId       uint64      `orm:"member_id"        description:"下单会员ID"`                                         // 下单会员ID
	CoachId        uint64      `orm:"coach_id"         description:"陪玩师ID"`                                          // 陪玩师ID
	ShopId         uint64      `orm:"shop_id"          description:"店铺ID（0表示无店铺）"`                                   // 店铺ID（0表示无店铺）
	GoodsId        uint64      `orm:"goods_id"         description:"商品ID"`                                           // 商品ID
	GoodsTitle     string      `orm:"goods_title"      description:"商品名称（冗余）"`                                       // 商品名称（冗余）
	GoodsPrice     int64       `orm:"goods_price"      description:"商品单价（分，下单时快照）"`                                  // 商品单价（分，下单时快照）
	Quantity       int         `orm:"quantity"         description:"数量"`                                             // 数量
	TotalAmount    int64       `orm:"total_amount"     description:"订单总额（分）"`                                        // 订单总额（分）
	DiscountAmount int64       `orm:"discount_amount"  description:"会员折扣金额（分）"`                                      // 会员折扣金额（分）
	CouponAmount   int64       `orm:"coupon_amount"    description:"优惠券抵扣金额（分）"`                                     // 优惠券抵扣金额（分）
	PayAmount      int64       `orm:"pay_amount"       description:"实付金额（分）"`                                        // 实付金额（分）
	CouponMemberId uint64      `orm:"coupon_member_id" description:"使用的优惠券领取记录ID"`                                   // 使用的优惠券领取记录ID
	PayType        int         `orm:"pay_type"         description:"支付方式:0=未支付,1=微信支付,2=支付宝支付,3=余额支付"`               // 支付方式:0=未支付,1=微信支付,2=支付宝支付,3=余额支付
	OrderStatus    int         `orm:"order_status"     description:"订单状态:0=待支付,1=已支付,2=进行中,3=已完成,4=已取消,5=退款中,6=已退款"` // 订单状态:0=待支付,1=已支付,2=进行中,3=已完成,4=已取消,5=退款中,6=已退款
	PayAt          *gtime.Time `orm:"pay_at"           description:"支付时间"`                                           // 支付时间
	StartAt        *gtime.Time `orm:"start_at"         description:"服务开始时间"`                                         // 服务开始时间
	FinishAt       *gtime.Time `orm:"finish_at"        description:"服务完成时间"`                                         // 服务完成时间
	CancelAt       *gtime.Time `orm:"cancel_at"        description:"取消时间"`                                           // 取消时间
	CancelReason   string      `orm:"cancel_reason"    description:"取消原因"`                                           // 取消原因
	Remark         string      `orm:"remark"           description:"订单备注"`                                           // 订单备注
	CreatedBy      uint64      `orm:"created_by"       description:"创建人ID"`                                          // 创建人ID
	DeptId         uint64      `orm:"dept_id"          description:"所属部门ID"`                                         // 所属部门ID
	CreatedAt      *gtime.Time `orm:"created_at"       description:"创建时间"`                                           // 创建时间
	UpdatedAt      *gtime.Time `orm:"updated_at"       description:"更新时间"`                                           // 更新时间
	DeletedAt      *gtime.Time `orm:"deleted_at"       description:"软删除时间"`                                          // 软删除时间
}
