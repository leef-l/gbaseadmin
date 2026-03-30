package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// Goods API

// GoodsCreateReq 创建商品表请求
type GoodsCreateReq struct {
	g.Meta `path:"/goods/create" method:"post" tags:"商品表" summary:"创建商品表"`
	CategoryID snowflake.JsonInt64 `json:"categoryID" v:"required#分类ID不能为空" dc:"分类ID"`
	CoachID snowflake.JsonInt64 `json:"coachID" v:"required#陪玩师ID不能为空" dc:"陪玩师ID"`
	Title string `json:"title" v:"required#商品名称不能为空" dc:"商品名称"`
	CoverImage string `json:"coverImage"  dc:"商品封面图"`
	DescContent string `json:"descContent"  dc:"商品详情描述"`
	Price int64 `json:"price"  dc:"单价（分）"`
	Unit string `json:"unit"  dc:"计量单位（如"`
	SalesNum int `json:"salesNum"  dc:"销量"`
	Sort int `json:"sort"  dc:"排序（升序）"`
	Status int `json:"status"  dc:"状态"`
}

// GoodsCreateRes 创建商品表响应
type GoodsCreateRes struct {
	g.Meta `mime:"application/json"`
}

// GoodsUpdateReq 更新商品表请求
type GoodsUpdateReq struct {
	g.Meta `path:"/goods/update" method:"put" tags:"商品表" summary:"更新商品表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"商品表ID"`
	CategoryID snowflake.JsonInt64 `json:"categoryID" dc:"分类ID"`
	CoachID snowflake.JsonInt64 `json:"coachID" dc:"陪玩师ID"`
	Title string `json:"title" dc:"商品名称"`
	CoverImage string `json:"coverImage" dc:"商品封面图"`
	DescContent string `json:"descContent" dc:"商品详情描述"`
	Price int64 `json:"price" dc:"单价（分）"`
	Unit string `json:"unit" dc:"计量单位（如"`
	SalesNum int `json:"salesNum" dc:"销量"`
	Sort int `json:"sort" dc:"排序（升序）"`
	Status int `json:"status" dc:"状态"`
}

// GoodsUpdateRes 更新商品表响应
type GoodsUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// GoodsDeleteReq 删除商品表请求
type GoodsDeleteReq struct {
	g.Meta `path:"/goods/delete" method:"delete" tags:"商品表" summary:"删除商品表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"商品表ID"`
}

// GoodsDeleteRes 删除商品表响应
type GoodsDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// GoodsDetailReq 获取商品表详情请求
type GoodsDetailReq struct {
	g.Meta `path:"/goods/detail" method:"get" tags:"商品表" summary:"获取商品表详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"商品表ID"`
}

// GoodsDetailRes 获取商品表详情响应
type GoodsDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.GoodsDetailOutput
}

// GoodsListReq 获取商品表列表请求
type GoodsListReq struct {
	g.Meta   `path:"/goods/list" method:"get" tags:"商品表" summary:"获取商品表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	Status int `json:"status" dc:"状态"`
}

// GoodsListRes 获取商品表列表响应
type GoodsListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.GoodsListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

