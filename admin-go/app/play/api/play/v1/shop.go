package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// Shop API

// ShopCreateReq 创建店铺表请求
type ShopCreateReq struct {
	g.Meta `path:"/shop/create" method:"post" tags:"店铺表" summary:"创建店铺表"`
	Title string `json:"title" v:"required#店铺名称不能为空" dc:"店铺名称"`
	LogoImage string `json:"logoImage"  dc:"店铺LOGO"`
	CoverImage string `json:"coverImage"  dc:"封面图"`
	ContactName string `json:"contactName"  dc:"联系人姓名"`
	ContactPhone string `json:"contactPhone"  dc:"联系电话"`
	Intro string `json:"intro"  dc:"店铺简介"`
	CommissionRate int `json:"commissionRate"  dc:"店铺抽成比例（百分比，如 10 表示 10%）"`
	CoachNum int `json:"coachNum"  dc:"陪玩师数量"`
	Sort int `json:"sort"  dc:"排序（升序）"`
	Status int `json:"status"  dc:"状态"`
}

// ShopCreateRes 创建店铺表响应
type ShopCreateRes struct {
	g.Meta `mime:"application/json"`
}

// ShopUpdateReq 更新店铺表请求
type ShopUpdateReq struct {
	g.Meta `path:"/shop/update" method:"put" tags:"店铺表" summary:"更新店铺表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"店铺表ID"`
	Title string `json:"title" dc:"店铺名称"`
	LogoImage string `json:"logoImage" dc:"店铺LOGO"`
	CoverImage string `json:"coverImage" dc:"封面图"`
	ContactName string `json:"contactName" dc:"联系人姓名"`
	ContactPhone string `json:"contactPhone" dc:"联系电话"`
	Intro string `json:"intro" dc:"店铺简介"`
	CommissionRate int `json:"commissionRate" dc:"店铺抽成比例（百分比，如 10 表示 10%）"`
	CoachNum int `json:"coachNum" dc:"陪玩师数量"`
	Sort int `json:"sort" dc:"排序（升序）"`
	Status int `json:"status" dc:"状态"`
}

// ShopUpdateRes 更新店铺表响应
type ShopUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// ShopDeleteReq 删除店铺表请求
type ShopDeleteReq struct {
	g.Meta `path:"/shop/delete" method:"delete" tags:"店铺表" summary:"删除店铺表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"店铺表ID"`
}

// ShopDeleteRes 删除店铺表响应
type ShopDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// ShopDetailReq 获取店铺表详情请求
type ShopDetailReq struct {
	g.Meta `path:"/shop/detail" method:"get" tags:"店铺表" summary:"获取店铺表详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"店铺表ID"`
}

// ShopDetailRes 获取店铺表详情响应
type ShopDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.ShopDetailOutput
}

// ShopListReq 获取店铺表列表请求
type ShopListReq struct {
	g.Meta   `path:"/shop/list" method:"get" tags:"店铺表" summary:"获取店铺表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	Status int `json:"status" dc:"状态"`
}

// ShopListRes 获取店铺表列表响应
type ShopListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.ShopListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

