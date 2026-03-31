// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayActivityStep is the golang structure for table play_activity_step.
type PlayActivityStep struct {
	Id          uint64      `orm:"id"           description:"步骤ID（Snowflake）"`
	ActivityId  uint64      `orm:"activity_id"  description:"活动ID"`
	StepNum     int         `orm:"step_num"     description:"步骤序号"`
	Title       string      `orm:"title"        description:"步骤标题"`
	StepType    int         `orm:"step_type"    description:"步骤类型：1=文字 2=链接 3=图片"`
	ExampleText string      `orm:"example_text" description:"示例文字或链接URL"`
	DescContent string      `orm:"desc_content" description:"步骤说明"`
	StepImage   string      `orm:"step_image"   description:"步骤示例图片"`
	Sort        int         `orm:"sort"         description:"排序"`
	CreatedBy   uint64      `orm:"created_by"   description:"创建人ID"`
	DeptId      uint64      `orm:"dept_id"      description:"所属部门ID"`
	CreatedAt   *gtime.Time `orm:"created_at"   description:"创建时间"`
	UpdatedAt   *gtime.Time `orm:"updated_at"   description:"更新时间"`
	DeletedAt   *gtime.Time `orm:"deleted_at"   description:"软删除时间"`
}
