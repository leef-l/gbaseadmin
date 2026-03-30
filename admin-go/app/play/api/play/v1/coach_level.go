package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// CoachLevel API

// CoachLevelCreateReq 创建陪玩师等级表请求
type CoachLevelCreateReq struct {
	g.Meta `path:"/coach_level/create" method:"post" tags:"陪玩师等级表" summary:"创建陪玩师等级表"`
	Title string `json:"title" v:"required#等级名称不能为空" dc:"等级名称"`
	Level int `json:"level"  dc:"等级"`
	Icon string `json:"icon"  dc:"等级图标"`
	MinOrders int `json:"minOrders"  dc:"所需最低接单数"`
	MinScore int `json:"minScore"  dc:"所需最低评分（乘100存储，如 450=4.50分）"`
	CommissionRate int `json:"commissionRate"  dc:"平台抽成比例（百分比，如 20 表示 20%）"`
	Sort int `json:"sort"  dc:"排序（升序）"`
	Status int `json:"status"  dc:"状态"`
}

// CoachLevelCreateRes 创建陪玩师等级表响应
type CoachLevelCreateRes struct {
	g.Meta `mime:"application/json"`
}

// CoachLevelUpdateReq 更新陪玩师等级表请求
type CoachLevelUpdateReq struct {
	g.Meta `path:"/coach_level/update" method:"put" tags:"陪玩师等级表" summary:"更新陪玩师等级表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"陪玩师等级表ID"`
	Title string `json:"title" dc:"等级名称"`
	Level int `json:"level" dc:"等级"`
	Icon string `json:"icon" dc:"等级图标"`
	MinOrders int `json:"minOrders" dc:"所需最低接单数"`
	MinScore int `json:"minScore" dc:"所需最低评分（乘100存储，如 450=4.50分）"`
	CommissionRate int `json:"commissionRate" dc:"平台抽成比例（百分比，如 20 表示 20%）"`
	Sort int `json:"sort" dc:"排序（升序）"`
	Status int `json:"status" dc:"状态"`
}

// CoachLevelUpdateRes 更新陪玩师等级表响应
type CoachLevelUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// CoachLevelDeleteReq 删除陪玩师等级表请求
type CoachLevelDeleteReq struct {
	g.Meta `path:"/coach_level/delete" method:"delete" tags:"陪玩师等级表" summary:"删除陪玩师等级表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"陪玩师等级表ID"`
}

// CoachLevelDeleteRes 删除陪玩师等级表响应
type CoachLevelDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// CoachLevelDetailReq 获取陪玩师等级表详情请求
type CoachLevelDetailReq struct {
	g.Meta `path:"/coach_level/detail" method:"get" tags:"陪玩师等级表" summary:"获取陪玩师等级表详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"陪玩师等级表ID"`
}

// CoachLevelDetailRes 获取陪玩师等级表详情响应
type CoachLevelDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.CoachLevelDetailOutput
}

// CoachLevelListReq 获取陪玩师等级表列表请求
type CoachLevelListReq struct {
	g.Meta   `path:"/coach_level/list" method:"get" tags:"陪玩师等级表" summary:"获取陪玩师等级表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	Level int `json:"level" dc:"等级"`
	Status int `json:"status" dc:"状态"`
}

// CoachLevelListRes 获取陪玩师等级表列表响应
type CoachLevelListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.CoachLevelListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

