// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayOauth is the golang structure of table play_oauth for DAO operations like Where/Data.
type PlayOauth struct {
	g.Meta       `orm:"table:play_oauth, do:true"`
	Id           any         // 记录ID（Snowflake）
	MemberId     any         // 会员ID
	Provider     any         // 第三方平台:1=微信,2=支付宝
	OpenId       any         // 第三方OpenID
	UnionId      any         // 第三方UnionID
	Nickname     any         // 第三方昵称
	Avatar       any         // 第三方头像
	AccessToken  any         // 访问令牌
	RefreshToken any         // 刷新令牌
	ExpireAt     *gtime.Time // 令牌过期时间
	CreatedBy    any         // 创建人ID
	DeptId       any         // 所属部门ID
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	DeletedAt    *gtime.Time // 软删除时间
}
