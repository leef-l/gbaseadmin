package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// Category API

// CategoryCreateReq 创建商品分类表请求
type CategoryCreateReq struct {
	g.Meta `path:"/category/create" method:"post" tags:"商品分类表" summary:"创建商品分类表"`
	ParentID snowflake.JsonInt64 `json:"parentID"  dc:"上级分类ID，0 表示顶级分类"`
	Title string `json:"title" v:"required#分类名称不能为空" dc:"分类名称"`
	Icon string `json:"icon"  dc:"分类图标"`
	CoverImage string `json:"coverImage"  dc:"分类封面图"`
	Sort int `json:"sort"  dc:"排序（升序）"`
	Status int `json:"status"  dc:"状态"`
}

// CategoryCreateRes 创建商品分类表响应
type CategoryCreateRes struct {
	g.Meta `mime:"application/json"`
}

// CategoryUpdateReq 更新商品分类表请求
type CategoryUpdateReq struct {
	g.Meta `path:"/category/update" method:"put" tags:"商品分类表" summary:"更新商品分类表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"商品分类表ID"`
	ParentID snowflake.JsonInt64 `json:"parentID" dc:"上级分类ID，0 表示顶级分类"`
	Title string `json:"title" dc:"分类名称"`
	Icon string `json:"icon" dc:"分类图标"`
	CoverImage string `json:"coverImage" dc:"分类封面图"`
	Sort int `json:"sort" dc:"排序（升序）"`
	Status int `json:"status" dc:"状态"`
}

// CategoryUpdateRes 更新商品分类表响应
type CategoryUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// CategoryDeleteReq 删除商品分类表请求
type CategoryDeleteReq struct {
	g.Meta `path:"/category/delete" method:"delete" tags:"商品分类表" summary:"删除商品分类表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"商品分类表ID"`
}

// CategoryDeleteRes 删除商品分类表响应
type CategoryDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// CategoryDetailReq 获取商品分类表详情请求
type CategoryDetailReq struct {
	g.Meta `path:"/category/detail" method:"get" tags:"商品分类表" summary:"获取商品分类表详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"商品分类表ID"`
}

// CategoryDetailRes 获取商品分类表详情响应
type CategoryDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.CategoryDetailOutput
}

// CategoryListReq 获取商品分类表列表请求
type CategoryListReq struct {
	g.Meta   `path:"/category/list" method:"get" tags:"商品分类表" summary:"获取商品分类表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	Status int `json:"status" dc:"状态"`
}

// CategoryListRes 获取商品分类表列表响应
type CategoryListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.CategoryListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

// CategoryTreeReq 获取商品分类表树形结构请求
type CategoryTreeReq struct {
	g.Meta `path:"/category/tree" method:"get" tags:"商品分类表" summary:"获取商品分类表树形结构"`
	Status int `json:"status" dc:"状态"`
}

// CategoryTreeRes 获取商品分类表树形结构响应
type CategoryTreeRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.CategoryTreeOutput `json:"list" dc:"树形数据"`
}

