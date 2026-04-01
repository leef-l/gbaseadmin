// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayReview is the golang structure for table play_review.
type PlayReview struct {
	Id            uint64      `orm:"id"             description:"评价ID（Snowflake）"`      // 评价ID（Snowflake）
	OrderId       uint64      `orm:"order_id"       description:"订单ID"`                 // 订单ID
	MemberId      uint64      `orm:"member_id"      description:"评价会员ID"`               // 评价会员ID
	CoachId       uint64      `orm:"coach_id"       description:"被评陪玩师ID"`              // 被评陪玩师ID
	Score         int         `orm:"score"          description:"评分（乘100，如 500=5.00分）"` // 评分（乘100，如 500=5.00分）
	ReviewContent string      `orm:"review_content" description:"评价内容"`                 // 评价内容
	ReviewImage   string      `orm:"review_image"   description:"评价图片（多张逗号分隔）"`         // 评价图片（多张逗号分隔）
	ReplyContent  string      `orm:"reply_content"  description:"陪玩师回复内容"`              // 陪玩师回复内容
	ReplyAt       *gtime.Time `orm:"reply_at"       description:"回复时间"`                 // 回复时间
	IsAnonymous   int         `orm:"is_anonymous"   description:"是否匿名:0=否,1=是"`         // 是否匿名:0=否,1=是
	Status        int         `orm:"status"         description:"状态:0=隐藏,1=显示"`         // 状态:0=隐藏,1=显示
	CreatedBy     uint64      `orm:"created_by"     description:"创建人ID"`                // 创建人ID
	DeptId        uint64      `orm:"dept_id"        description:"所属部门ID"`               // 所属部门ID
	CreatedAt     *gtime.Time `orm:"created_at"     description:"创建时间"`                 // 创建时间
	UpdatedAt     *gtime.Time `orm:"updated_at"     description:"更新时间"`                 // 更新时间
	DeletedAt     *gtime.Time `orm:"deleted_at"     description:"软删除时间"`                // 软删除时间
}
