package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// Oauth API

// OauthCreateReq 创建第三方登录绑定表请求
type OauthCreateReq struct {
	g.Meta `path:"/oauth/create" method:"post" tags:"第三方登录绑定表" summary:"创建第三方登录绑定表"`
	MemberID snowflake.JsonInt64 `json:"memberID" v:"required#会员ID不能为空" dc:"会员ID"`
	Provider int `json:"provider" v:"required#第三方平台不能为空" dc:"第三方平台"`
	OpenID string `json:"openID" v:"required#第三方OpenID不能为空" dc:"第三方OpenID"`
	UnionID string `json:"unionID"  dc:"第三方UnionID"`
	Nickname string `json:"nickname"  dc:"第三方昵称"`
	Avatar string `json:"avatar"  dc:"第三方头像"`
	AccessToken string `json:"accessToken"  dc:"访问令牌"`
	RefreshToken string `json:"refreshToken"  dc:"刷新令牌"`
	ExpireAt *gtime.Time `json:"expireAt"  dc:"令牌过期时间"`
}

// OauthCreateRes 创建第三方登录绑定表响应
type OauthCreateRes struct {
	g.Meta `mime:"application/json"`
}

// OauthUpdateReq 更新第三方登录绑定表请求
type OauthUpdateReq struct {
	g.Meta `path:"/oauth/update" method:"put" tags:"第三方登录绑定表" summary:"更新第三方登录绑定表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"第三方登录绑定表ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" dc:"会员ID"`
	Provider int `json:"provider" dc:"第三方平台"`
	OpenID string `json:"openID" dc:"第三方OpenID"`
	UnionID string `json:"unionID" dc:"第三方UnionID"`
	Nickname string `json:"nickname" dc:"第三方昵称"`
	Avatar string `json:"avatar" dc:"第三方头像"`
	AccessToken string `json:"accessToken" dc:"访问令牌"`
	RefreshToken string `json:"refreshToken" dc:"刷新令牌"`
	ExpireAt *gtime.Time `json:"expireAt" dc:"令牌过期时间"`
}

// OauthUpdateRes 更新第三方登录绑定表响应
type OauthUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// OauthDeleteReq 删除第三方登录绑定表请求
type OauthDeleteReq struct {
	g.Meta `path:"/oauth/delete" method:"delete" tags:"第三方登录绑定表" summary:"删除第三方登录绑定表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"第三方登录绑定表ID"`
}

// OauthDeleteRes 删除第三方登录绑定表响应
type OauthDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// OauthDetailReq 获取第三方登录绑定表详情请求
type OauthDetailReq struct {
	g.Meta `path:"/oauth/detail" method:"get" tags:"第三方登录绑定表" summary:"获取第三方登录绑定表详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"第三方登录绑定表ID"`
}

// OauthDetailRes 获取第三方登录绑定表详情响应
type OauthDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.OauthDetailOutput
}

// OauthListReq 获取第三方登录绑定表列表请求
type OauthListReq struct {
	g.Meta   `path:"/oauth/list" method:"get" tags:"第三方登录绑定表" summary:"获取第三方登录绑定表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	Provider int `json:"provider" dc:"第三方平台"`
}

// OauthListRes 获取第三方登录绑定表列表响应
type OauthListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.OauthListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

