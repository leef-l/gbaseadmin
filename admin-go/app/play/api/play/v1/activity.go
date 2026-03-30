package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// Activity API

// ActivityCreateReq 创建活动表请求
type ActivityCreateReq struct {
	g.Meta `path:"/activity/create" method:"post" tags:"活动表" summary:"创建活动表"`
	Title string `json:"title" v:"required#活动名称不能为空" dc:"活动名称"`
	CoverImage string `json:"coverImage"  dc:"活动封面图"`
	DescContent string `json:"descContent"  dc:"活动详情描述（富文本，支持图文混排）"`
	Type int `json:"type"  dc:"活动类型"`
	ConditionType int `json:"conditionType"  dc:"参与条件"`
	ConditionValue int64 `json:"conditionValue"  dc:"条件值（分/次，如充值满5000分、下单满3次）"`
	IsAutoReward int `json:"isAutoReward"  dc:"是否自动发奖"`
	StartAt *gtime.Time `json:"startAt" v:"required#活动开始时间不能为空" dc:"活动开始时间"`
	EndAt *gtime.Time `json:"endAt" v:"required#活动结束时间不能为空" dc:"活动结束时间"`
	MaxNum int `json:"maxNum"  dc:"参与人数上限（0表示不限）"`
	JoinNum int `json:"joinNum"  dc:"已参与人数"`
	Sort int `json:"sort"  dc:"排序（升序）"`
	Status int `json:"status"  dc:"状态"`
}

// ActivityCreateRes 创建活动表响应
type ActivityCreateRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityUpdateReq 更新活动表请求
type ActivityUpdateReq struct {
	g.Meta `path:"/activity/update" method:"put" tags:"活动表" summary:"更新活动表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"活动表ID"`
	Title string `json:"title" dc:"活动名称"`
	CoverImage string `json:"coverImage" dc:"活动封面图"`
	DescContent string `json:"descContent" dc:"活动详情描述（富文本，支持图文混排）"`
	Type int `json:"type" dc:"活动类型"`
	ConditionType int `json:"conditionType" dc:"参与条件"`
	ConditionValue int64 `json:"conditionValue" dc:"条件值（分/次，如充值满5000分、下单满3次）"`
	IsAutoReward int `json:"isAutoReward" dc:"是否自动发奖"`
	StartAt *gtime.Time `json:"startAt" dc:"活动开始时间"`
	EndAt *gtime.Time `json:"endAt" dc:"活动结束时间"`
	MaxNum int `json:"maxNum" dc:"参与人数上限（0表示不限）"`
	JoinNum int `json:"joinNum" dc:"已参与人数"`
	Sort int `json:"sort" dc:"排序（升序）"`
	Status int `json:"status" dc:"状态"`
}

// ActivityUpdateRes 更新活动表响应
type ActivityUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityDeleteReq 删除活动表请求
type ActivityDeleteReq struct {
	g.Meta `path:"/activity/delete" method:"delete" tags:"活动表" summary:"删除活动表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"活动表ID"`
}

// ActivityDeleteRes 删除活动表响应
type ActivityDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityDetailReq 获取活动表详情请求
type ActivityDetailReq struct {
	g.Meta `path:"/activity/detail" method:"get" tags:"活动表" summary:"获取活动表详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"活动表ID"`
}

// ActivityDetailRes 获取活动表详情响应
type ActivityDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.ActivityDetailOutput
}

// ActivityListReq 获取活动表列表请求
type ActivityListReq struct {
	g.Meta   `path:"/activity/list" method:"get" tags:"活动表" summary:"获取活动表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	Type int `json:"type" dc:"活动类型"`
	ConditionType int `json:"conditionType" dc:"参与条件"`
	IsAutoReward int `json:"isAutoReward" dc:"是否自动发奖"`
	Status int `json:"status" dc:"状态"`
}

// ActivityListRes 获取活动表列表响应
type ActivityListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.ActivityListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

