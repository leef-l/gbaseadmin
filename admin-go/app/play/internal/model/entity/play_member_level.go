// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayMemberLevel is the golang structure for table play_member_level.
type PlayMemberLevel struct {
	Id        uint64      `orm:"id"         description:"等级ID（Snowflake）"`                       // 等级ID（Snowflake）
	Title     string      `orm:"title"      description:"等级名称"`                                  // 等级名称
	Level     int         `orm:"level"      description:"等级:1=普通会员,2=白银会员,3=黄金会员,4=铂金会员,5=钻石会员"` // 等级:1=普通会员,2=白银会员,3=黄金会员,4=铂金会员,5=钻石会员
	Icon      string      `orm:"icon"       description:"等级图标"`                                  // 等级图标
	MinExp    int         `orm:"min_exp"    description:"所需最低经验值"`                               // 所需最低经验值
	Discount  int         `orm:"discount"   description:"折扣（百分比，如 90 表示九折）"`                     // 折扣（百分比，如 90 表示九折）
	Sort      int         `orm:"sort"       description:"排序（升序）"`                                // 排序（升序）
	Status    int         `orm:"status"     description:"状态:0=关闭,1=开启"`                          // 状态:0=关闭,1=开启
	CreatedBy uint64      `orm:"created_by" description:"创建人ID"`                                 // 创建人ID
	DeptId    uint64      `orm:"dept_id"    description:"所属部门ID"`                                // 所属部门ID
	CreatedAt *gtime.Time `orm:"created_at" description:"创建时间"`                                  // 创建时间
	UpdatedAt *gtime.Time `orm:"updated_at" description:"更新时间"`                                  // 更新时间
	DeletedAt *gtime.Time `orm:"deleted_at" description:"软删除时间"`                                 // 软删除时间
}
