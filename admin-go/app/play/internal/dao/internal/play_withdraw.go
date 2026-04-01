// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayWithdrawDao is the data access object for the table play_withdraw.
type PlayWithdrawDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  PlayWithdrawColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// PlayWithdrawColumns defines and stores column names for the table play_withdraw.
type PlayWithdrawColumns struct {
	Id        string // 提现ID
	CoachId   string // 陪玩师ID
	MemberId  string // 会员ID
	Amount    string // 提现金额(分)
	Status    string // 状态 0=待审核 1=已打款 2=已拒绝
	Reason    string // 拒绝原因
	AuditedAt string // 审核时间
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// playWithdrawColumns holds the columns for the table play_withdraw.
var playWithdrawColumns = PlayWithdrawColumns{
	Id:        "id",
	CoachId:   "coach_id",
	MemberId:  "member_id",
	Amount:    "amount",
	Status:    "status",
	Reason:    "reason",
	AuditedAt: "audited_at",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewPlayWithdrawDao creates and returns a new DAO object for table data access.
func NewPlayWithdrawDao(handlers ...gdb.ModelHandler) *PlayWithdrawDao {
	return &PlayWithdrawDao{
		group:    "default",
		table:    "play_withdraw",
		columns:  playWithdrawColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayWithdrawDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayWithdrawDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayWithdrawDao) Columns() PlayWithdrawColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayWithdrawDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayWithdrawDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayWithdrawDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
