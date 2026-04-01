// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayActivityStep is the golang structure of table play_activity_step for DAO operations like Where/Data.
type PlayActivityStep struct {
	g.Meta      `orm:"table:play_activity_step, do:true"`
	Id          any         // 步骤ID（Snowflake）
	ActivityId  any         // 活动ID
	StepNum     any         // 步骤序号
	Title       any         // 步骤标题
	StepType    any         // 步骤类型：1=文字 2=链接 3=图片
	ExampleText any         // 示例文字或链接URL
	DescContent any         // 步骤说明（富文本，支持图文）
	StepImage   any         // 步骤示例图片
	IsRequired  any         // 是否需要填写:0=不需要,1=需要
	Sort        any         // 排序（升序）
	CreatedBy   any         // 创建人ID
	DeptId      any         // 所属部门ID
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 软删除时间
}
