package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// Coach API

// CoachCreateReq 创建陪玩师表请求
type CoachCreateReq struct {
	g.Meta `path:"/coach/create" method:"post" tags:"陪玩师表" summary:"创建陪玩师表"`
	MemberID snowflake.JsonInt64 `json:"memberID" v:"required#关联会员ID不能为空" dc:"关联会员ID"`
	CoachLevelID snowflake.JsonInt64 `json:"coachLevelID"  dc:"陪玩师等级ID"`
	ShopID snowflake.JsonInt64 `json:"shopID"  dc:"所属店铺ID（0表示无店铺）"`
	RealName string `json:"realName" v:"required#真实姓名不能为空" dc:"真实姓名"`
	Intro string `json:"intro"  dc:"个人简介"`
	CoverImage string `json:"coverImage"  dc:"封面图"`
	TotalOrders int `json:"totalOrders"  dc:"总接单数"`
	TotalScore int `json:"totalScore"  dc:"总评分（乘100，如 500=5.00）"`
	ScoreNum int `json:"scoreNum"  dc:"评分人数"`
	IncomeTotal int64 `json:"incomeTotal"  dc:"累计收入（分）"`
	IncomeBalance int64 `json:"incomeBalance"  dc:"可提现余额（分）"`
	IsOnline int `json:"isOnline"  dc:"是否在线"`
	Sort int `json:"sort"  dc:"排序（升序）"`
	Status int `json:"status"  dc:"状态"`
}

// CoachCreateRes 创建陪玩师表响应
type CoachCreateRes struct {
	g.Meta `mime:"application/json"`
}

// CoachUpdateReq 更新陪玩师表请求
type CoachUpdateReq struct {
	g.Meta `path:"/coach/update" method:"put" tags:"陪玩师表" summary:"更新陪玩师表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"陪玩师表ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" dc:"关联会员ID"`
	CoachLevelID snowflake.JsonInt64 `json:"coachLevelID" dc:"陪玩师等级ID"`
	ShopID snowflake.JsonInt64 `json:"shopID" dc:"所属店铺ID（0表示无店铺）"`
	RealName string `json:"realName" dc:"真实姓名"`
	Intro string `json:"intro" dc:"个人简介"`
	CoverImage string `json:"coverImage" dc:"封面图"`
	TotalOrders int `json:"totalOrders" dc:"总接单数"`
	TotalScore int `json:"totalScore" dc:"总评分（乘100，如 500=5.00）"`
	ScoreNum int `json:"scoreNum" dc:"评分人数"`
	IncomeTotal int64 `json:"incomeTotal" dc:"累计收入（分）"`
	IncomeBalance int64 `json:"incomeBalance" dc:"可提现余额（分）"`
	IsOnline int `json:"isOnline" dc:"是否在线"`
	Sort int `json:"sort" dc:"排序（升序）"`
	Status int `json:"status" dc:"状态"`
}

// CoachUpdateRes 更新陪玩师表响应
type CoachUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// CoachDeleteReq 删除陪玩师表请求
type CoachDeleteReq struct {
	g.Meta `path:"/coach/delete" method:"delete" tags:"陪玩师表" summary:"删除陪玩师表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"陪玩师表ID"`
}

// CoachDeleteRes 删除陪玩师表响应
type CoachDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// CoachDetailReq 获取陪玩师表详情请求
type CoachDetailReq struct {
	g.Meta `path:"/coach/detail" method:"get" tags:"陪玩师表" summary:"获取陪玩师表详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"陪玩师表ID"`
}

// CoachDetailRes 获取陪玩师表详情响应
type CoachDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.CoachDetailOutput
}

// CoachListReq 获取陪玩师表列表请求
type CoachListReq struct {
	g.Meta   `path:"/coach/list" method:"get" tags:"陪玩师表" summary:"获取陪玩师表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	IsOnline int `json:"isOnline" dc:"是否在线"`
	Status int `json:"status" dc:"状态"`
}

// CoachListRes 获取陪玩师表列表响应
type CoachListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.CoachListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

