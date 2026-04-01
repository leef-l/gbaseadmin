// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayMember is the golang structure of table play_member for DAO operations like Where/Data.
type PlayMember struct {
	g.Meta        `orm:"table:play_member, do:true"`
	Id            any         // 会员ID（Snowflake）
	Phone         any         // 手机号
	Password      any         // 密码（bcrypt 加密）
	Nickname      any         // 昵称
	Avatar        any         // 头像
	Gender        any         // 性别:0=未知,1=男,2=女
	MemberLevelId any         // 会员等级ID
	VipExpireAt   *gtime.Time // VIP到期时间
	Exp           any         // 经验值
	Balance       any         // 账户余额（分）
	IsCoach       any         // 是否陪玩师:0=否,1=是
	Status        any         // 状态:0=禁用,1=正常
	LastLoginAt   *gtime.Time // 最后登录时间
	CreatedBy     any         // 创建人ID
	DeptId        any         // 所属部门ID
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 软删除时间
}
