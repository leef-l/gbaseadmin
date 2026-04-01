// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayGoods is the golang structure for table play_goods.
type PlayGoods struct {
	Id          uint64      `orm:"id"           description:"商品ID（Snowflake）"` // 商品ID（Snowflake）
	CategoryId  uint64      `orm:"category_id"  description:"分类ID"`            // 分类ID
	CoachId     uint64      `orm:"coach_id"     description:"陪玩师ID"`           // 陪玩师ID
	Title       string      `orm:"title"        description:"商品名称"`            // 商品名称
	CoverImage  string      `orm:"cover_image"  description:"商品封面图"`           // 商品封面图
	DescContent string      `orm:"desc_content" description:"商品详情描述"`          // 商品详情描述
	Price       int64       `orm:"price"        description:"单价（分）"`           // 单价（分）
	Unit        string      `orm:"unit"         description:"计量单位（如：局、小时、把）"`  // 计量单位（如：局、小时、把）
	SalesNum    int         `orm:"sales_num"    description:"销量"`              // 销量
	Sort        int         `orm:"sort"         description:"排序（升序）"`          // 排序（升序）
	Status      int         `orm:"status"       description:"状态:0=下架,1=上架"`    // 状态:0=下架,1=上架
	CreatedBy   uint64      `orm:"created_by"   description:"创建人ID"`           // 创建人ID
	DeptId      uint64      `orm:"dept_id"      description:"所属部门ID"`          // 所属部门ID
	CreatedAt   *gtime.Time `orm:"created_at"   description:"创建时间"`            // 创建时间
	UpdatedAt   *gtime.Time `orm:"updated_at"   description:"更新时间"`            // 更新时间
	DeletedAt   *gtime.Time `orm:"deleted_at"   description:"软删除时间"`           // 软删除时间
}
