// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayCoach is the golang structure for table play_coach.
type PlayCoach struct {
	Id            uint64      `orm:"id"             description:"陪玩师ID（Snowflake）"`     // 陪玩师ID（Snowflake）
	MemberId      uint64      `orm:"member_id"      description:"关联会员ID"`               // 关联会员ID
	CoachLevelId  uint64      `orm:"coach_level_id" description:"陪玩师等级ID"`              // 陪玩师等级ID
	ShopId        uint64      `orm:"shop_id"        description:"所属店铺ID（0表示无店铺）"`       // 所属店铺ID（0表示无店铺）
	RealName      string      `orm:"real_name"      description:"真实姓名"`                 // 真实姓名
	Intro         string      `orm:"intro"          description:"个人简介"`                 // 个人简介
	CoverImage    string      `orm:"cover_image"    description:"封面图"`                  // 封面图
	TotalOrders   int         `orm:"total_orders"   description:"总接单数"`                 // 总接单数
	TotalScore    int         `orm:"total_score"    description:"总评分（乘100，如 500=5.00）"` // 总评分（乘100，如 500=5.00）
	ScoreNum      int         `orm:"score_num"      description:"评分人数"`                 // 评分人数
	IncomeTotal   int64       `orm:"income_total"   description:"累计收入（分）"`              // 累计收入（分）
	IncomeBalance int64       `orm:"income_balance" description:"可提现余额（分）"`             // 可提现余额（分）
	IsOnline      int         `orm:"is_online"      description:"是否在线:0=离线,1=在线"`       // 是否在线:0=离线,1=在线
	Sort          int         `orm:"sort"           description:"排序（升序）"`               // 排序（升序）
	Status        int         `orm:"status"         description:"状态:0=禁用,1=正常"`         // 状态:0=禁用,1=正常
	CreatedBy     uint64      `orm:"created_by"     description:"创建人ID"`                // 创建人ID
	DeptId        uint64      `orm:"dept_id"        description:"所属部门ID"`               // 所属部门ID
	CreatedAt     *gtime.Time `orm:"created_at"     description:"创建时间"`                 // 创建时间
	UpdatedAt     *gtime.Time `orm:"updated_at"     description:"更新时间"`                 // 更新时间
	DeletedAt     *gtime.Time `orm:"deleted_at"     description:"软删除时间"`                // 软删除时间
}
