// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayCoachLevel is the golang structure for table play_coach_level.
type PlayCoachLevel struct {
	Id             uint64      `orm:"id"              description:"等级ID（Snowflake）"`             // 等级ID（Snowflake）
	Title          string      `orm:"title"           description:"等级名称"`                        // 等级名称
	Level          int         `orm:"level"           description:"等级:1=青铜,2=白银,3=黄金,4=铂金,5=钻石"` // 等级:1=青铜,2=白银,3=黄金,4=铂金,5=钻石
	Icon           string      `orm:"icon"            description:"等级图标"`                        // 等级图标
	MinOrders      int         `orm:"min_orders"      description:"所需最低接单数"`                     // 所需最低接单数
	MinScore       int         `orm:"min_score"       description:"所需最低评分（乘100存储，如 450=4.50分）"`  // 所需最低评分（乘100存储，如 450=4.50分）
	CommissionRate int         `orm:"commission_rate" description:"平台抽成比例（百分比，如 20 表示 20%）"`     // 平台抽成比例（百分比，如 20 表示 20%）
	Sort           int         `orm:"sort"            description:"排序（升序）"`                      // 排序（升序）
	Status         int         `orm:"status"          description:"状态:0=关闭,1=开启"`                // 状态:0=关闭,1=开启
	CreatedBy      uint64      `orm:"created_by"      description:"创建人ID"`                       // 创建人ID
	DeptId         uint64      `orm:"dept_id"         description:"所属部门ID"`                      // 所属部门ID
	CreatedAt      *gtime.Time `orm:"created_at"      description:"创建时间"`                        // 创建时间
	UpdatedAt      *gtime.Time `orm:"updated_at"      description:"更新时间"`                        // 更新时间
	DeletedAt      *gtime.Time `orm:"deleted_at"      description:"软删除时间"`                       // 软删除时间
}
