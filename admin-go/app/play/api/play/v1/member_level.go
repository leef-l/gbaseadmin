package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// MemberLevel API

// MemberLevelCreateReq 创建会员等级表请求
type MemberLevelCreateReq struct {
	g.Meta `path:"/member_level/create" method:"post" tags:"会员等级表" summary:"创建会员等级表"`
	Title string `json:"title" v:"required#等级名称不能为空" dc:"等级名称"`
	Level int `json:"level"  dc:"等级"`
	Icon string `json:"icon"  dc:"等级图标"`
	MinExp int `json:"minExp"  dc:"所需最低经验值"`
	Discount int `json:"discount"  dc:"折扣（百分比，如 90 表示九折）"`
	Sort int `json:"sort"  dc:"排序（升序）"`
	Status int `json:"status"  dc:"状态"`
}

// MemberLevelCreateRes 创建会员等级表响应
type MemberLevelCreateRes struct {
	g.Meta `mime:"application/json"`
}

// MemberLevelUpdateReq 更新会员等级表请求
type MemberLevelUpdateReq struct {
	g.Meta `path:"/member_level/update" method:"put" tags:"会员等级表" summary:"更新会员等级表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"会员等级表ID"`
	Title string `json:"title" dc:"等级名称"`
	Level int `json:"level" dc:"等级"`
	Icon string `json:"icon" dc:"等级图标"`
	MinExp int `json:"minExp" dc:"所需最低经验值"`
	Discount int `json:"discount" dc:"折扣（百分比，如 90 表示九折）"`
	Sort int `json:"sort" dc:"排序（升序）"`
	Status int `json:"status" dc:"状态"`
}

// MemberLevelUpdateRes 更新会员等级表响应
type MemberLevelUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// MemberLevelDeleteReq 删除会员等级表请求
type MemberLevelDeleteReq struct {
	g.Meta `path:"/member_level/delete" method:"delete" tags:"会员等级表" summary:"删除会员等级表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"会员等级表ID"`
}

// MemberLevelDeleteRes 删除会员等级表响应
type MemberLevelDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// MemberLevelDetailReq 获取会员等级表详情请求
type MemberLevelDetailReq struct {
	g.Meta `path:"/member_level/detail" method:"get" tags:"会员等级表" summary:"获取会员等级表详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"会员等级表ID"`
}

// MemberLevelDetailRes 获取会员等级表详情响应
type MemberLevelDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.MemberLevelDetailOutput
}

// MemberLevelListReq 获取会员等级表列表请求
type MemberLevelListReq struct {
	g.Meta   `path:"/member_level/list" method:"get" tags:"会员等级表" summary:"获取会员等级表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	Level int `json:"level" dc:"等级"`
	Status int `json:"status" dc:"状态"`
}

// MemberLevelListRes 获取会员等级表列表响应
type MemberLevelListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.MemberLevelListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

