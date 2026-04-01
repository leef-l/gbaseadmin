// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayOauth is the golang structure for table play_oauth.
type PlayOauth struct {
	Id           uint64      `orm:"id"            description:"记录ID（Snowflake）"`  // 记录ID（Snowflake）
	MemberId     uint64      `orm:"member_id"     description:"会员ID"`             // 会员ID
	Provider     int         `orm:"provider"      description:"第三方平台:1=微信,2=支付宝"` // 第三方平台:1=微信,2=支付宝
	OpenId       string      `orm:"open_id"       description:"第三方OpenID"`        // 第三方OpenID
	UnionId      string      `orm:"union_id"      description:"第三方UnionID"`       // 第三方UnionID
	Nickname     string      `orm:"nickname"      description:"第三方昵称"`            // 第三方昵称
	Avatar       string      `orm:"avatar"        description:"第三方头像"`            // 第三方头像
	AccessToken  string      `orm:"access_token"  description:"访问令牌"`             // 访问令牌
	RefreshToken string      `orm:"refresh_token" description:"刷新令牌"`             // 刷新令牌
	ExpireAt     *gtime.Time `orm:"expire_at"     description:"令牌过期时间"`           // 令牌过期时间
	CreatedBy    uint64      `orm:"created_by"    description:"创建人ID"`            // 创建人ID
	DeptId       uint64      `orm:"dept_id"       description:"所属部门ID"`           // 所属部门ID
	CreatedAt    *gtime.Time `orm:"created_at"    description:"创建时间"`             // 创建时间
	UpdatedAt    *gtime.Time `orm:"updated_at"    description:"更新时间"`             // 更新时间
	DeletedAt    *gtime.Time `orm:"deleted_at"    description:"软删除时间"`            // 软删除时间
}
