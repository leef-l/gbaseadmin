package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// RechargePlan API

// RechargePlanCreateReq 创建充值方案表请求
type RechargePlanCreateReq struct {
	g.Meta `path:"/recharge_plan/create" method:"post" tags:"充值方案表" summary:"创建充值方案表"`
	Title string `json:"title" v:"required#方案名称不能为空" dc:"方案名称"`
	Amount int64 `json:"amount" v:"required#充值金额（分）不能为空" dc:"充值金额（分）"`
	GiftAmount int64 `json:"giftAmount"  dc:"赠送金额（分）"`
	CoverImage string `json:"coverImage"  dc:"方案封面图"`
	Sort int `json:"sort"  dc:"排序（升序）"`
	Status int `json:"status"  dc:"状态"`
}

// RechargePlanCreateRes 创建充值方案表响应
type RechargePlanCreateRes struct {
	g.Meta `mime:"application/json"`
}

// RechargePlanUpdateReq 更新充值方案表请求
type RechargePlanUpdateReq struct {
	g.Meta `path:"/recharge_plan/update" method:"put" tags:"充值方案表" summary:"更新充值方案表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"充值方案表ID"`
	Title string `json:"title" dc:"方案名称"`
	Amount int64 `json:"amount" dc:"充值金额（分）"`
	GiftAmount int64 `json:"giftAmount" dc:"赠送金额（分）"`
	CoverImage string `json:"coverImage" dc:"方案封面图"`
	Sort int `json:"sort" dc:"排序（升序）"`
	Status int `json:"status" dc:"状态"`
}

// RechargePlanUpdateRes 更新充值方案表响应
type RechargePlanUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// RechargePlanDeleteReq 删除充值方案表请求
type RechargePlanDeleteReq struct {
	g.Meta `path:"/recharge_plan/delete" method:"delete" tags:"充值方案表" summary:"删除充值方案表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"充值方案表ID"`
}

// RechargePlanDeleteRes 删除充值方案表响应
type RechargePlanDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// RechargePlanDetailReq 获取充值方案表详情请求
type RechargePlanDetailReq struct {
	g.Meta `path:"/recharge_plan/detail" method:"get" tags:"充值方案表" summary:"获取充值方案表详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"充值方案表ID"`
}

// RechargePlanDetailRes 获取充值方案表详情响应
type RechargePlanDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.RechargePlanDetailOutput
}

// RechargePlanListReq 获取充值方案表列表请求
type RechargePlanListReq struct {
	g.Meta   `path:"/recharge_plan/list" method:"get" tags:"充值方案表" summary:"获取充值方案表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	Status int `json:"status" dc:"状态"`
}

// RechargePlanListRes 获取充值方案表列表响应
type RechargePlanListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.RechargePlanListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

