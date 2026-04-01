// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayGoods is the golang structure of table play_goods for DAO operations like Where/Data.
type PlayGoods struct {
	g.Meta      `orm:"table:play_goods, do:true"`
	Id          any         // 商品ID（Snowflake）
	CategoryId  any         // 分类ID
	CoachId     any         // 陪玩师ID
	Title       any         // 商品名称
	CoverImage  any         // 商品封面图
	DescContent any         // 商品详情描述
	Price       any         // 单价（分）
	Unit        any         // 计量单位（如：局、小时、把）
	SalesNum    any         // 销量
	Sort        any         // 排序（升序）
	Status      any         // 状态:0=下架,1=上架
	CreatedBy   any         // 创建人ID
	DeptId      any         // 所属部门ID
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 软删除时间
}
