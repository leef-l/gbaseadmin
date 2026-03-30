// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayMemberDao is the data access object for the table play_member.
type PlayMemberDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PlayMemberColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PlayMemberColumns defines and stores column names for the table play_member.
type PlayMemberColumns struct {
	Id            string // 会员ID（Snowflake）
	Phone         string // 手机号
	Password      string // 密码（bcrypt 加密）
	Nickname      string // 昵称
	Avatar        string // 头像
	Gender        string // 性别:0=未知,1=男,2=女
	MemberLevelId string // 会员等级ID
	Exp           string // 经验值
	Balance       string // 账户余额（分）
	IsCoach       string // 是否陪玩师:0=否,1=是
	Status        string // 状态:0=禁用,1=正常
	LastLoginAt   string // 最后登录时间
	CreatedBy     string // 创建人ID
	DeptId        string // 所属部门ID
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	DeletedAt     string // 软删除时间
}

// playMemberColumns holds the columns for the table play_member.
var playMemberColumns = PlayMemberColumns{
	Id:            "id",
	Phone:         "phone",
	Password:      "password",
	Nickname:      "nickname",
	Avatar:        "avatar",
	Gender:        "gender",
	MemberLevelId: "member_level_id",
	Exp:           "exp",
	Balance:       "balance",
	IsCoach:       "is_coach",
	Status:        "status",
	LastLoginAt:   "last_login_at",
	CreatedBy:     "created_by",
	DeptId:        "dept_id",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
}

// NewPlayMemberDao creates and returns a new DAO object for table data access.
func NewPlayMemberDao(handlers ...gdb.ModelHandler) *PlayMemberDao {
	return &PlayMemberDao{
		group:    "default",
		table:    "play_member",
		columns:  playMemberColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayMemberDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayMemberDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayMemberDao) Columns() PlayMemberColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayMemberDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayMemberDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PlayMemberDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
