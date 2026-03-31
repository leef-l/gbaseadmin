package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ========== 公开接口 ==========

type GoodsListReq struct {
	g.Meta     `path:"/goods/list" method:"get" tags:"C端商品" summary:"商品列表"`
	CategoryID string `json:"categoryId" dc:"分类ID筛选"`
	Keyword    string `json:"keyword" dc:"关键词搜索"`
	SortBy     string `json:"sortBy" v:"in:price_asc,price_desc,sales,newest#排序值不合法" dc:"排序"`
	Page       int    `json:"page" v:"min:1" dc:"页码" d:"1"`
	PageSize   int    `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}

type GoodsListItem struct {
	GoodsID     string  `json:"goodsId" dc:"商品ID"`
	Title       string  `json:"title" dc:"商品标题"`
	CoverImage  string  `json:"coverImage" dc:"商品封面图"`
	Price       int64   `json:"price" dc:"单价(分)"`
	Unit        string  `json:"unit" dc:"单位"`
	SalesNum    int     `json:"salesNum" dc:"销量"`
	CoachID     string  `json:"coachId" dc:"陪玩师ID"`
	CoachName   string  `json:"coachName" dc:"陪玩师昵称"`
	CoachAvatar string  `json:"coachAvatar" dc:"陪玩师头像"`
	CoachScore  float64 `json:"coachScore" dc:"陪玩师评分"`
	CoachOnline int     `json:"coachOnline" dc:"陪玩师在线状态"`
}

type GoodsListRes struct {
	g.Meta `mime:"application/json"`
	Total  int             `json:"total" dc:"总数"`
	List   []GoodsListItem `json:"list" dc:"商品列表"`
}

type GoodsDetailReq struct {
	g.Meta  `path:"/goods/detail" method:"get" tags:"C端商品" summary:"商品详情"`
	GoodsID string `json:"goodsId" v:"required#商品ID不能为空" dc:"商品ID"`
}

type GoodsDetailRes struct {
	g.Meta       `mime:"application/json"`
	GoodsID      string  `json:"goodsId" dc:"商品ID"`
	Title        string  `json:"title" dc:"商品标题"`
	Description  string  `json:"description" dc:"商品描述"`
	CoverImage   string  `json:"coverImage" dc:"商品封面图"`
	CategoryID   string  `json:"categoryId" dc:"分类ID"`
	CategoryName string  `json:"categoryName" dc:"分类名称"`
	Price        int64   `json:"price" dc:"单价(分)"`
	Unit         string  `json:"unit" dc:"单位"`
	SalesNum     int     `json:"salesNum" dc:"销量"`
	CoachID      string  `json:"coachId" dc:"陪玩师ID"`
	CoachName    string  `json:"coachName" dc:"陪玩师昵称"`
	CoachAvatar  string  `json:"coachAvatar" dc:"陪玩师头像"`
	CoachScore   float64 `json:"coachScore" dc:"陪玩师评分"`
	CoachOnline  int     `json:"coachOnline" dc:"陪玩师在线状态"`
}

type CategoryListReq struct {
	g.Meta `path:"/category/list" method:"get" tags:"C端商品" summary:"分类列表（树形）"`
}

type CategoryTreeItem struct {
	CategoryID string             `json:"categoryId" dc:"分类ID"`
	Name       string             `json:"name" dc:"分类名称"`
	Icon       string             `json:"icon" dc:"分类图标"`
	CoverImage string             `json:"coverImage" dc:"分类封面图"`
	Sort       int                `json:"sort" dc:"排序"`
	Children   []CategoryTreeItem `json:"children" dc:"子分类"`
}

type CategoryListRes struct {
	g.Meta `mime:"application/json"`
	List   []CategoryTreeItem `json:"list" dc:"分类树形列表"`
}
