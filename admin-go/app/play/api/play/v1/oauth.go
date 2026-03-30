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

// OauthCreateReq 创建ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨请求
type OauthCreateReq struct {
	g.Meta `path:"/oauth/create" method:"post" tags:"ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨" summary:"创建ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨"`
	MemberID snowflake.JsonInt64 `json:"memberID" v:"required#ä¼šå‘˜ID不能为空" dc:"ä¼šå‘˜ID"`
	Provider int `json:"provider" v:"required#ç¬¬ä¸‰æ–¹å¹³å°不能为空" dc:"ç¬¬ä¸‰æ–¹å¹³å°"`
	OpenID snowflake.JsonInt64 `json:"openID" v:"required#ç¬¬ä¸‰æ–¹OpenID不能为空" dc:"ç¬¬ä¸‰æ–¹OpenID"`
	UnionID snowflake.JsonInt64 `json:"unionID"  dc:"ç¬¬ä¸‰æ–¹UnionID"`
	Nickname string `json:"nickname"  dc:"ç¬¬ä¸‰æ–¹æ˜µç§°"`
	Avatar string `json:"avatar"  dc:"ç¬¬ä¸‰æ–¹å¤´åƒ"`
	AccessToken string `json:"accessToken"  dc:"è®¿é—®ä»¤ç‰Œ"`
	RefreshToken string `json:"refreshToken"  dc:"åˆ·æ–°ä»¤ç‰Œ"`
	ExpireAt *gtime.Time `json:"expireAt"  dc:"ä»¤ç‰Œè¿‡æœŸæ—¶é—´"`
}

// OauthCreateRes 创建ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨响应
type OauthCreateRes struct {
	g.Meta `mime:"application/json"`
}

// OauthUpdateReq 更新ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨请求
type OauthUpdateReq struct {
	g.Meta `path:"/oauth/update" method:"put" tags:"ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨" summary:"更新ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" dc:"ä¼šå‘˜ID"`
	Provider int `json:"provider" dc:"ç¬¬ä¸‰æ–¹å¹³å°"`
	OpenID snowflake.JsonInt64 `json:"openID" dc:"ç¬¬ä¸‰æ–¹OpenID"`
	UnionID snowflake.JsonInt64 `json:"unionID" dc:"ç¬¬ä¸‰æ–¹UnionID"`
	Nickname string `json:"nickname" dc:"ç¬¬ä¸‰æ–¹æ˜µç§°"`
	Avatar string `json:"avatar" dc:"ç¬¬ä¸‰æ–¹å¤´åƒ"`
	AccessToken string `json:"accessToken" dc:"è®¿é—®ä»¤ç‰Œ"`
	RefreshToken string `json:"refreshToken" dc:"åˆ·æ–°ä»¤ç‰Œ"`
	ExpireAt *gtime.Time `json:"expireAt" dc:"ä»¤ç‰Œè¿‡æœŸæ—¶é—´"`
}

// OauthUpdateRes 更新ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨响应
type OauthUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// OauthDeleteReq 删除ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨请求
type OauthDeleteReq struct {
	g.Meta `path:"/oauth/delete" method:"delete" tags:"ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨" summary:"删除ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨ID"`
}

// OauthDeleteRes 删除ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨响应
type OauthDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// OauthDetailReq 获取ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨详情请求
type OauthDetailReq struct {
	g.Meta `path:"/oauth/detail" method:"get" tags:"ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨" summary:"获取ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨ID"`
}

// OauthDetailRes 获取ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨详情响应
type OauthDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.OauthDetailOutput
}

// OauthListReq 获取ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨列表请求
type OauthListReq struct {
	g.Meta   `path:"/oauth/list" method:"get" tags:"ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨" summary:"获取ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	Provider int `json:"provider" dc:"ç¬¬ä¸‰æ–¹å¹³å°"`
}

// OauthListRes 获取ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨列表响应
type OauthListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.OauthListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

