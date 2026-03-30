// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayCoachDao is the data access object for the table play_coach.
type PlayCoachDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PlayCoachColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PlayCoachColumns defines and stores column names for the table play_coach.
type PlayCoachColumns struct {
	Id            string // 陪玩师ID（Snowflake）
	MemberId      string // 关联会员ID
	CoachLevelId  string // 陪玩师等级ID
	ShopId        string // 所属店铺ID（0表示无店铺）
	RealName      string // 真实姓名
	Intro         string // 个人简介
	CoverImage    string // 封面图
	TotalOrders   string // 总接单数
	TotalScore    string // 总评分（乘100，如 500=5.00）
	ScoreNum      string // 评分人数
	IncomeTotal   string // 累计收入（分）
	IncomeBalance string // 可提现余额（分）
	IsOnline      string // 是否在线:0=离线,1=在线
	Sort          string // 排序（升序）
	Status        string // 状态:0=禁用,1=正常
	CreatedBy     string // 创建人ID
	DeptId        string // 所属部门ID
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	DeletedAt     string // 软删除时间
}

// playCoachColumns holds the columns for the table play_coach.
var playCoachColumns = PlayCoachColumns{
	Id:            "id",
	MemberId:      "member_id",
	CoachLevelId:  "coach_level_id",
	ShopId:        "shop_id",
	RealName:      "real_name",
	Intro:         "intro",
	CoverImage:    "cover_image",
	TotalOrders:   "total_orders",
	TotalScore:    "total_score",
	ScoreNum:      "score_num",
	IncomeTotal:   "income_total",
	IncomeBalance: "income_balance",
	IsOnline:      "is_online",
	Sort:          "sort",
	Status:        "status",
	CreatedBy:     "created_by",
	DeptId:        "dept_id",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
}

// NewPlayCoachDao creates and returns a new DAO object for table data access.
func NewPlayCoachDao(handlers ...gdb.ModelHandler) *PlayCoachDao {
	return &PlayCoachDao{
		group:    "default",
		table:    "play_coach",
		columns:  playCoachColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayCoachDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayCoachDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayCoachDao) Columns() PlayCoachColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayCoachDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayCoachDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayCoachDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
