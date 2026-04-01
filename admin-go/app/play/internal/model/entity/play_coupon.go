// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayCoupon is the golang structure for table play_coupon.
type PlayCoupon struct {
	Id           uint64      `orm:"id"             description:"优惠券ID（Snowflake）"`                   // 优惠券ID（Snowflake）
	Title        string      `orm:"title"          description:"优惠券名称"`                              // 优惠券名称
	Type         int         `orm:"type"           description:"优惠券类型:1=满减券,2=折扣券,3=无门槛券"`           // 优惠券类型:1=满减券,2=折扣券,3=无门槛券
	IsNewMember  int         `orm:"is_new_member"  description:"是否新人专享:0=否,1=是"`                     // 是否新人专享:0=否,1=是
	FaceValue    int64       `orm:"face_value"     description:"面值（分，满减/无门槛时为抵扣额，折扣时为折扣值如 85=8.5折）"` // 面值（分，满减/无门槛时为抵扣额，折扣时为折扣值如 85=8.5折）
	MinAmount    int64       `orm:"min_amount"     description:"最低消费金额（分，0表示无门槛）"`                   // 最低消费金额（分，0表示无门槛）
	TotalNum     int         `orm:"total_num"      description:"发放总量（0表示不限）"`                        // 发放总量（0表示不限）
	UsedNum      int         `orm:"used_num"       description:"已使用数量"`                              // 已使用数量
	ClaimNum     int         `orm:"claim_num"      description:"已领取数量"`                              // 已领取数量
	PerLimit     int         `orm:"per_limit"      description:"每人限领张数"`                             // 每人限领张数
	ValidStartAt *gtime.Time `orm:"valid_start_at" description:"有效期开始时间"`                            // 有效期开始时间
	ValidEndAt   *gtime.Time `orm:"valid_end_at"   description:"有效期结束时间"`                            // 有效期结束时间
	Sort         int         `orm:"sort"           description:"排序（升序）"`                             // 排序（升序）
	Status       int         `orm:"status"         description:"状态:0=关闭,1=开启"`                       // 状态:0=关闭,1=开启
	CreatedBy    uint64      `orm:"created_by"     description:"创建人ID"`                              // 创建人ID
	DeptId       uint64      `orm:"dept_id"        description:"所属部门ID"`                             // 所属部门ID
	CreatedAt    *gtime.Time `orm:"created_at"     description:"创建时间"`                               // 创建时间
	UpdatedAt    *gtime.Time `orm:"updated_at"     description:"更新时间"`                               // 更新时间
	DeletedAt    *gtime.Time `orm:"deleted_at"     description:"软删除时间"`                              // 软删除时间
}
