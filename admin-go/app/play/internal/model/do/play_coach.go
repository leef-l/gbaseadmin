// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayCoach is the golang structure of table play_coach for DAO operations like Where/Data.
type PlayCoach struct {
	g.Meta        `orm:"table:play_coach, do:true"`
	Id            any         // 陪玩师ID（Snowflake）
	MemberId      any         // 关联会员ID
	CoachLevelId  any         // 陪玩师等级ID
	ShopId        any         // 所属店铺ID（0表示无店铺）
	RealName      any         // 真实姓名
	Intro         any         // 个人简介
	CoverImage    any         // 封面图
	TotalOrders   any         // 总接单数
	TotalScore    any         // 总评分（乘100，如 500=5.00）
	ScoreNum      any         // 评分人数
	IncomeTotal   any         // 累计收入（分）
	IncomeBalance any         // 可提现余额（分）
	IsOnline      any         // 是否在线:0=离线,1=在线
	Sort          any         // 排序（升序）
	Status        any         // 状态:0=禁用,1=正常
	CreatedBy     any         // 创建人ID
	DeptId        any         // 所属部门ID
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 软删除时间
}
