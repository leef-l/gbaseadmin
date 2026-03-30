// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayCategory is the golang structure of table play_category for DAO operations like Where/Data.
type PlayCategory struct {
	g.Meta     `orm:"table:play_category, do:true"`
	Id         any         // 分类ID（Snowflake）
	ParentId   any         // 上级分类ID，0 表示顶级分类
	Title      any         // 分类名称
	Icon       any         // 分类图标
	CoverImage any         // 分类封面图
	Sort       any         // 排序（升序）
	Status     any         // 状态:0=关闭,1=开启
	CreatedBy  any         // 创建人ID
	DeptId     any         // 所属部门ID
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	DeletedAt  *gtime.Time // 软删除时间
}
