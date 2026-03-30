package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// RechargePlan API

// RechargePlanCreateReq 创建å……å€¼æ–¹æ¡ˆè¡¨请求
type RechargePlanCreateReq struct {
	g.Meta `path:"/recharge_plan/create" method:"post" tags:"å……å€¼æ–¹æ¡ˆè¡¨" summary:"创建å……å€¼æ–¹æ¡ˆè¡¨"`
	Title string `json:"title" v:"required#æ–¹æ¡ˆåç§°不能为空" dc:"æ–¹æ¡ˆåç§°"`
	Amount int64 `json:"amount" v:"required#å……å€¼é‡‘é¢ï¼ˆåˆ†ï¼‰不能为空" dc:"å……å€¼é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	GiftAmount int64 `json:"giftAmount"  dc:"èµ é€é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	CoverImage string `json:"coverImage"  dc:"æ–¹æ¡ˆå°é¢å›¾"`
	Sort int `json:"sort"  dc:"æŽ’åºï¼ˆå‡åºï¼‰"`
	Status int `json:"status"  dc:"çŠ¶æ€"`
}

// RechargePlanCreateRes 创建å……å€¼æ–¹æ¡ˆè¡¨响应
type RechargePlanCreateRes struct {
	g.Meta `mime:"application/json"`
}

// RechargePlanUpdateReq 更新å……å€¼æ–¹æ¡ˆè¡¨请求
type RechargePlanUpdateReq struct {
	g.Meta `path:"/recharge_plan/update" method:"put" tags:"å……å€¼æ–¹æ¡ˆè¡¨" summary:"更新å……å€¼æ–¹æ¡ˆè¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"å……å€¼æ–¹æ¡ˆè¡¨ID"`
	Title string `json:"title" dc:"æ–¹æ¡ˆåç§°"`
	Amount int64 `json:"amount" dc:"å……å€¼é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	GiftAmount int64 `json:"giftAmount" dc:"èµ é€é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	CoverImage string `json:"coverImage" dc:"æ–¹æ¡ˆå°é¢å›¾"`
	Sort int `json:"sort" dc:"æŽ’åºï¼ˆå‡åºï¼‰"`
	Status int `json:"status" dc:"çŠ¶æ€"`
}

// RechargePlanUpdateRes 更新å……å€¼æ–¹æ¡ˆè¡¨响应
type RechargePlanUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// RechargePlanDeleteReq 删除å……å€¼æ–¹æ¡ˆè¡¨请求
type RechargePlanDeleteReq struct {
	g.Meta `path:"/recharge_plan/delete" method:"delete" tags:"å……å€¼æ–¹æ¡ˆè¡¨" summary:"删除å……å€¼æ–¹æ¡ˆè¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"å……å€¼æ–¹æ¡ˆè¡¨ID"`
}

// RechargePlanDeleteRes 删除å……å€¼æ–¹æ¡ˆè¡¨响应
type RechargePlanDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// RechargePlanDetailReq 获取å……å€¼æ–¹æ¡ˆè¡¨详情请求
type RechargePlanDetailReq struct {
	g.Meta `path:"/recharge_plan/detail" method:"get" tags:"å……å€¼æ–¹æ¡ˆè¡¨" summary:"获取å……å€¼æ–¹æ¡ˆè¡¨详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"å……å€¼æ–¹æ¡ˆè¡¨ID"`
}

// RechargePlanDetailRes 获取å……å€¼æ–¹æ¡ˆè¡¨详情响应
type RechargePlanDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.RechargePlanDetailOutput
}

// RechargePlanListReq 获取å……å€¼æ–¹æ¡ˆè¡¨列表请求
type RechargePlanListReq struct {
	g.Meta   `path:"/recharge_plan/list" method:"get" tags:"å……å€¼æ–¹æ¡ˆè¡¨" summary:"获取å……å€¼æ–¹æ¡ˆè¡¨列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	Status int `json:"status" dc:"çŠ¶æ€"`
}

// RechargePlanListRes 获取å……å€¼æ–¹æ¡ˆè¡¨列表响应
type RechargePlanListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.RechargePlanListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

