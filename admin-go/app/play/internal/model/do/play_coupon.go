// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayCoupon is the golang structure of table play_coupon for DAO operations like Where/Data.
type PlayCoupon struct {
	g.Meta       `orm:"table:play_coupon, do:true"`
	Id           any         // 优惠券ID（Snowflake）
	Title        any         // 优惠券名称
	Type         any         // 优惠券类型:1=满减券,2=折扣券,3=无门槛券
	IsNewMember  any         // 是否新人专享:0=否,1=是
	FaceValue    any         // 面值（分，满减/无门槛时为抵扣额，折扣时为折扣值如 85=8.5折）
	MinAmount    any         // 最低消费金额（分，0表示无门槛）
	TotalNum     any         // 发放总量（0表示不限）
	UsedNum      any         // 已使用数量
	ClaimNum     any         // 已领取数量
	PerLimit     any         // 每人限领张数
	ValidStartAt *gtime.Time // 有效期开始时间
	ValidEndAt   *gtime.Time // 有效期结束时间
	Sort         any         // 排序（升序）
	Status       any         // 状态:0=关闭,1=开启
	CreatedBy    any         // 创建人ID
	DeptId       any         // 所属部门ID
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	DeletedAt    *gtime.Time // 软删除时间
}
