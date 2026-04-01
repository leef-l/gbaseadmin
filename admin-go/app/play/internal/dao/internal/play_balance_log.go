// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayBalanceLogDao is the data access object for the table play_balance_log.
type PlayBalanceLogDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  PlayBalanceLogColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// PlayBalanceLogColumns defines and stores column names for the table play_balance_log.
type PlayBalanceLogColumns struct {
	Id            string // 流水ID（Snowflake）
	MemberId      string // 会员ID
	BizType       string // 业务类型:1=充值,2=消费,3=退款,4=活动赠送,5=提现
	BizId         string // 关联业务ID（订单ID/充值订单ID/活动ID）
	ChangeAmount  string // 变动金额（分，正数增加负数减少）
	BeforeBalance string // 变动前余额（分）
	AfterBalance  string // 变动后余额（分）
	Remark        string // 备注说明
	CreatedBy     string // 创建人ID
	DeptId        string // 所属部门ID
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	DeletedAt     string // 软删除时间
}

// playBalanceLogColumns holds the columns for the table play_balance_log.
var playBalanceLogColumns = PlayBalanceLogColumns{
	Id:            "id",
	MemberId:      "member_id",
	BizType:       "biz_type",
	BizId:         "biz_id",
	ChangeAmount:  "change_amount",
	BeforeBalance: "before_balance",
	AfterBalance:  "after_balance",
	Remark:        "remark",
	CreatedBy:     "created_by",
	DeptId:        "dept_id",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
}

// NewPlayBalanceLogDao creates and returns a new DAO object for table data access.
func NewPlayBalanceLogDao(handlers ...gdb.ModelHandler) *PlayBalanceLogDao {
	return &PlayBalanceLogDao{
		group:    "default",
		table:    "play_balance_log",
		columns:  playBalanceLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayBalanceLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayBalanceLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayBalanceLogDao) Columns() PlayBalanceLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayBalanceLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayBalanceLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayBalanceLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
