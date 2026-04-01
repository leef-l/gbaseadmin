// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayCouponMember is the golang structure for table play_coupon_member.
type PlayCouponMember struct {
	Id        uint64      `orm:"id"         description:"记录ID（Snowflake）"`        // 记录ID（Snowflake）
	CouponId  uint64      `orm:"coupon_id"  description:"优惠券模板ID"`                // 优惠券模板ID
	MemberId  uint64      `orm:"member_id"  description:"会员ID"`                   // 会员ID
	OrderId   uint64      `orm:"order_id"   description:"使用的订单ID（0表示未使用）"`        // 使用的订单ID（0表示未使用）
	UseStatus int         `orm:"use_status" description:"使用状态:0=未使用,1=已使用,2=已过期"` // 使用状态:0=未使用,1=已使用,2=已过期
	ClaimAt   *gtime.Time `orm:"claim_at"   description:"领取时间"`                   // 领取时间
	UseAt     *gtime.Time `orm:"use_at"     description:"使用时间"`                   // 使用时间
	ExpireAt  *gtime.Time `orm:"expire_at"  description:"过期时间"`                   // 过期时间
	CreatedBy uint64      `orm:"created_by" description:"创建人ID"`                  // 创建人ID
	DeptId    uint64      `orm:"dept_id"    description:"所属部门ID"`                 // 所属部门ID
	CreatedAt *gtime.Time `orm:"created_at" description:"创建时间"`                   // 创建时间
	UpdatedAt *gtime.Time `orm:"updated_at" description:"更新时间"`                   // 更新时间
	DeletedAt *gtime.Time `orm:"deleted_at" description:"软删除时间"`                  // 软删除时间
}
