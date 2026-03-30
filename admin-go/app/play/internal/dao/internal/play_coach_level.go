// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayCoachLevelDao is the data access object for the table play_coach_level.
type PlayCoachLevelDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  PlayCoachLevelColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// PlayCoachLevelColumns defines and stores column names for the table play_coach_level.
type PlayCoachLevelColumns struct {
	Id             string // 等级ID（Snowflake）
	Title          string // 等级名称
	Level          string // 等级:1=青铜,2=白银,3=黄金,4=铂金,5=钻石
	Icon           string // 等级图标
	MinOrders      string // 所需最低接单数
	MinScore       string // 所需最低评分（乘100存储，如 450=4.50分）
	CommissionRate string // 平台抽成比例（百分比，如 20 表示 20%）
	Sort           string // 排序（升序）
	Status         string // 状态:0=关闭,1=开启
	CreatedBy      string // 创建人ID
	DeptId         string // 所属部门ID
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	DeletedAt      string // 软删除时间
}

// playCoachLevelColumns holds the columns for the table play_coach_level.
var playCoachLevelColumns = PlayCoachLevelColumns{
	Id:             "id",
	Title:          "title",
	Level:          "level",
	Icon:           "icon",
	MinOrders:      "min_orders",
	MinScore:       "min_score",
	CommissionRate: "commission_rate",
	Sort:           "sort",
	Status:         "status",
	CreatedBy:      "created_by",
	DeptId:         "dept_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewPlayCoachLevelDao creates and returns a new DAO object for table data access.
func NewPlayCoachLevelDao(handlers ...gdb.ModelHandler) *PlayCoachLevelDao {
	return &PlayCoachLevelDao{
		group:    "default",
		table:    "play_coach_level",
		columns:  playCoachLevelColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayCoachLevelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayCoachLevelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayCoachLevelDao) Columns() PlayCoachLevelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayCoachLevelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayCoachLevelDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayCoachLevelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
