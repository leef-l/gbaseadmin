// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayMember is the golang structure for table play_member.
type PlayMember struct {
	Id            uint64      `orm:"id"              description:"会员ID（Snowflake）"` // 会员ID（Snowflake）
	Phone         string      `orm:"phone"           description:"手机号"`             // 手机号
	Password      string      `orm:"password"        description:"密码（bcrypt 加密）"`   // 密码（bcrypt 加密）
	Nickname      string      `orm:"nickname"        description:"昵称"`              // 昵称
	Avatar        string      `orm:"avatar"          description:"头像"`              // 头像
	Gender        int         `orm:"gender"          description:"性别:0=未知,1=男,2=女"` // 性别:0=未知,1=男,2=女
	MemberLevelId uint64      `orm:"member_level_id" description:"会员等级ID"`          // 会员等级ID
	VipExpireAt   *gtime.Time `orm:"vip_expire_at"   description:"VIP到期时间"`         // VIP到期时间
	Exp           int         `orm:"exp"             description:"经验值"`             // 经验值
	Balance       int64       `orm:"balance"         description:"账户余额（分）"`         // 账户余额（分）
	IsCoach       int         `orm:"is_coach"        description:"是否陪玩师:0=否,1=是"`   // 是否陪玩师:0=否,1=是
	Status        int         `orm:"status"          description:"状态:0=禁用,1=正常"`    // 状态:0=禁用,1=正常
	LastLoginAt   *gtime.Time `orm:"last_login_at"   description:"最后登录时间"`          // 最后登录时间
	CreatedBy     uint64      `orm:"created_by"      description:"创建人ID"`           // 创建人ID
	DeptId        uint64      `orm:"dept_id"         description:"所属部门ID"`          // 所属部门ID
	CreatedAt     *gtime.Time `orm:"created_at"      description:"创建时间"`            // 创建时间
	UpdatedAt     *gtime.Time `orm:"updated_at"      description:"更新时间"`            // 更新时间
	DeletedAt     *gtime.Time `orm:"deleted_at"      description:"软删除时间"`           // 软删除时间
}
