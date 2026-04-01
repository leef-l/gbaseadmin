// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayCouponMember is the golang structure of table play_coupon_member for DAO operations like Where/Data.
type PlayCouponMember struct {
	g.Meta    `orm:"table:play_coupon_member, do:true"`
	Id        any         // 记录ID（Snowflake）
	CouponId  any         // 优惠券模板ID
	MemberId  any         // 会员ID
	OrderId   any         // 使用的订单ID（0表示未使用）
	UseStatus any         // 使用状态:0=未使用,1=已使用,2=已过期
	ClaimAt   *gtime.Time // 领取时间
	UseAt     *gtime.Time // 使用时间
	ExpireAt  *gtime.Time // 过期时间
	CreatedBy any         // 创建人ID
	DeptId    any         // 所属部门ID
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 软删除时间
}
