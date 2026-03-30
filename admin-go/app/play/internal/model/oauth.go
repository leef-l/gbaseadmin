package model

import (
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/utility/snowflake"
)

// Oauth DTO 模型

// OauthCreateInput 创建ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨输入
type OauthCreateInput struct {
	MemberID snowflake.JsonInt64 `json:"memberID"`
	Provider int `json:"provider"`
	OpenID snowflake.JsonInt64 `json:"openID"`
	UnionID snowflake.JsonInt64 `json:"unionID"`
	Nickname string `json:"nickname"`
	Avatar string `json:"avatar"`
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpireAt *gtime.Time `json:"expireAt"`
}

// OauthUpdateInput 更新ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨输入
type OauthUpdateInput struct {
	ID snowflake.JsonInt64 `json:"id"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	Provider int `json:"provider"`
	OpenID snowflake.JsonInt64 `json:"openID"`
	UnionID snowflake.JsonInt64 `json:"unionID"`
	Nickname string `json:"nickname"`
	Avatar string `json:"avatar"`
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpireAt *gtime.Time `json:"expireAt"`
}

// OauthDetailOutput ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨详情输出
type OauthDetailOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	Provider int `json:"provider"`
	OpenID snowflake.JsonInt64 `json:"openID"`
	UnionID snowflake.JsonInt64 `json:"unionID"`
	Nickname string `json:"nickname"`
	Avatar string `json:"avatar"`
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpireAt *gtime.Time `json:"expireAt"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// OauthListOutput ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨列表输出
type OauthListOutput struct {
	ID snowflake.JsonInt64 `json:"id"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
	Provider int `json:"provider"`
	OpenID snowflake.JsonInt64 `json:"openID"`
	UnionID snowflake.JsonInt64 `json:"unionID"`
	Nickname string `json:"nickname"`
	Avatar string `json:"avatar"`
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpireAt *gtime.Time `json:"expireAt"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// OauthListInput ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨列表查询输入
type OauthListInput struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Provider int `json:"provider"`
}

