// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayGoodsDao is the data access object for the table play_goods.
type PlayGoodsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PlayGoodsColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PlayGoodsColumns defines and stores column names for the table play_goods.
type PlayGoodsColumns struct {
	Id          string // 商品ID（Snowflake）
	CategoryId  string // 分类ID
	CoachId     string // 陪玩师ID
	Title       string // 商品名称
	CoverImage  string // 商品封面图
	DescContent string // 商品详情描述
	Price       string // 单价（分）
	Unit        string // 计量单位（如：局、小时、把）
	SalesNum    string // 销量
	Sort        string // 排序（升序）
	Status      string // 状态:0=下架,1=上架
	CreatedBy   string // 创建人ID
	DeptId      string // 所属部门ID
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	DeletedAt   string // 软删除时间
}

// playGoodsColumns holds the columns for the table play_goods.
var playGoodsColumns = PlayGoodsColumns{
	Id:          "id",
	CategoryId:  "category_id",
	CoachId:     "coach_id",
	Title:       "title",
	CoverImage:  "cover_image",
	DescContent: "desc_content",
	Price:       "price",
	Unit:        "unit",
	SalesNum:    "sales_num",
	Sort:        "sort",
	Status:      "status",
	CreatedBy:   "created_by",
	DeptId:      "dept_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewPlayGoodsDao creates and returns a new DAO object for table data access.
func NewPlayGoodsDao(handlers ...gdb.ModelHandler) *PlayGoodsDao {
	return &PlayGoodsDao{
		group:    "default",
		table:    "play_goods",
		columns:  playGoodsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayGoodsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayGoodsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayGoodsDao) Columns() PlayGoodsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayGoodsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayGoodsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayGoodsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
