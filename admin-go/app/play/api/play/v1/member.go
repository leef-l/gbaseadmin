package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// Member API

// MemberCreateReq 创建会员表请求
type MemberCreateReq struct {
	g.Meta `path:"/member/create" method:"post" tags:"会员表" summary:"创建会员表"`
	Phone string `json:"phone" v:"required#手机号不能为空" dc:"手机号"`
	Password string `json:"password" v:"required#密码（bcrypt 加密）不能为空" dc:"密码（bcrypt 加密）"`
	Nickname string `json:"nickname"  dc:"昵称"`
	Avatar string `json:"avatar"  dc:"头像"`
	Gender int `json:"gender"  dc:"性别"`
	MemberLevelID snowflake.JsonInt64 `json:"memberLevelID"  dc:"会员等级ID"`
	Exp int `json:"exp"  dc:"经验值"`
	Balance int64 `json:"balance"  dc:"账户余额（分）"`
	IsCoach int `json:"isCoach"  dc:"是否陪玩师"`
	Status int `json:"status"  dc:"状态"`
	LastLoginAt *gtime.Time `json:"lastLoginAt"  dc:"最后登录时间"`
}

// MemberCreateRes 创建会员表响应
type MemberCreateRes struct {
	g.Meta `mime:"application/json"`
}

// MemberUpdateReq 更新会员表请求
type MemberUpdateReq struct {
	g.Meta `path:"/member/update" method:"put" tags:"会员表" summary:"更新会员表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"会员表ID"`
	Phone string `json:"phone" dc:"手机号"`
	Password string `json:"password" dc:"密码（bcrypt 加密）"`
	Nickname string `json:"nickname" dc:"昵称"`
	Avatar string `json:"avatar" dc:"头像"`
	Gender int `json:"gender" dc:"性别"`
	MemberLevelID snowflake.JsonInt64 `json:"memberLevelID" dc:"会员等级ID"`
	Exp int `json:"exp" dc:"经验值"`
	Balance int64 `json:"balance" dc:"账户余额（分）"`
	IsCoach int `json:"isCoach" dc:"是否陪玩师"`
	Status int `json:"status" dc:"状态"`
	LastLoginAt *gtime.Time `json:"lastLoginAt" dc:"最后登录时间"`
}

// MemberUpdateRes 更新会员表响应
type MemberUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// MemberDeleteReq 删除会员表请求
type MemberDeleteReq struct {
	g.Meta `path:"/member/delete" method:"delete" tags:"会员表" summary:"删除会员表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"会员表ID"`
}

// MemberDeleteRes 删除会员表响应
type MemberDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// MemberDetailReq 获取会员表详情请求
type MemberDetailReq struct {
	g.Meta `path:"/member/detail" method:"get" tags:"会员表" summary:"获取会员表详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"会员表ID"`
}

// MemberDetailRes 获取会员表详情响应
type MemberDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.MemberDetailOutput
}

// MemberListReq 获取会员表列表请求
type MemberListReq struct {
	g.Meta   `path:"/member/list" method:"get" tags:"会员表" summary:"获取会员表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	Gender int `json:"gender" dc:"性别"`
	IsCoach int `json:"isCoach" dc:"是否陪玩师"`
	Status int `json:"status" dc:"状态"`
}

// MemberListRes 获取会员表列表响应
type MemberListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.MemberListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

