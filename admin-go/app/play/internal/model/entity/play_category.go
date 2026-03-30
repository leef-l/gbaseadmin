// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayCategory is the golang structure for table play_category.
type PlayCategory struct {
	Id         uint64      `orm:"id"          description:"分类ID（Snowflake）"` // 分类ID（Snowflake）
	ParentId   uint64      `orm:"parent_id"   description:"上级分类ID，0 表示顶级分类"` // 上级分类ID，0 表示顶级分类
	Title      string      `orm:"title"       description:"分类名称"`            // 分类名称
	Icon       string      `orm:"icon"        description:"分类图标"`            // 分类图标
	CoverImage string      `orm:"cover_image" description:"分类封面图"`           // 分类封面图
	Sort       int         `orm:"sort"        description:"排序（升序）"`          // 排序（升序）
	Status     int         `orm:"status"      description:"状态:0=关闭,1=开启"`    // 状态:0=关闭,1=开启
	CreatedBy  uint64      `orm:"created_by"  description:"创建人ID"`           // 创建人ID
	DeptId     uint64      `orm:"dept_id"     description:"所属部门ID"`          // 所属部门ID
	CreatedAt  *gtime.Time `orm:"created_at"  description:"创建时间"`            // 创建时间
	UpdatedAt  *gtime.Time `orm:"updated_at"  description:"更新时间"`            // 更新时间
	DeletedAt  *gtime.Time `orm:"deleted_at"  description:"软删除时间"`           // 软删除时间
}
