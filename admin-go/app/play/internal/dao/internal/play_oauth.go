// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayOauthDao is the data access object for the table play_oauth.
type PlayOauthDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PlayOauthColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PlayOauthColumns defines and stores column names for the table play_oauth.
type PlayOauthColumns struct {
	Id           string // 记录ID（Snowflake）
	MemberId     string // 会员ID
	Provider     string // 第三方平台:1=微信,2=支付宝
	OpenId       string // 第三方OpenID
	UnionId      string // 第三方UnionID
	Nickname     string // 第三方昵称
	Avatar       string // 第三方头像
	AccessToken  string // 访问令牌
	RefreshToken string // 刷新令牌
	ExpireAt     string // 令牌过期时间
	CreatedBy    string // 创建人ID
	DeptId       string // 所属部门ID
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
	DeletedAt    string // 软删除时间
}

// playOauthColumns holds the columns for the table play_oauth.
var playOauthColumns = PlayOauthColumns{
	Id:           "id",
	MemberId:     "member_id",
	Provider:     "provider",
	OpenId:       "open_id",
	UnionId:      "union_id",
	Nickname:     "nickname",
	Avatar:       "avatar",
	AccessToken:  "access_token",
	RefreshToken: "refresh_token",
	ExpireAt:     "expire_at",
	CreatedBy:    "created_by",
	DeptId:       "dept_id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

// NewPlayOauthDao creates and returns a new DAO object for table data access.
func NewPlayOauthDao(handlers ...gdb.ModelHandler) *PlayOauthDao {
	return &PlayOauthDao{
		group:    "default",
		table:    "play_oauth",
		columns:  playOauthColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayOauthDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayOauthDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayOauthDao) Columns() PlayOauthColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayOauthDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayOauthDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayOauthDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
