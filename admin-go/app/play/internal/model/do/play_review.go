// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayReview is the golang structure of table play_review for DAO operations like Where/Data.
type PlayReview struct {
	g.Meta        `orm:"table:play_review, do:true"`
	Id            any         // 评价ID（Snowflake）
	OrderId       any         // 订单ID
	MemberId      any         // 评价会员ID
	CoachId       any         // 被评陪玩师ID
	Score         any         // 评分（乘100，如 500=5.00分）
	ReviewContent any         // 评价内容
	ReviewImage   any         // 评价图片（多张逗号分隔）
	ReplyContent  any         // 陪玩师回复内容
	ReplyAt       *gtime.Time // 回复时间
	IsAnonymous   any         // 是否匿名:0=否,1=是
	Status        any         // 状态:0=隐藏,1=显示
	CreatedBy     any         // 创建人ID
	DeptId        any         // 所属部门ID
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 软删除时间
}
