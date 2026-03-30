// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayCategoryDao is the data access object for the table play_category.
type PlayCategoryDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  PlayCategoryColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// PlayCategoryColumns defines and stores column names for the table play_category.
type PlayCategoryColumns struct {
	Id         string // 分类ID（Snowflake）
	ParentId   string // 上级分类ID，0 表示顶级分类
	Title      string // 分类名称
	Icon       string // 分类图标
	CoverImage string // 分类封面图
	Sort       string // 排序（升序）
	Status     string // 状态:0=关闭,1=开启
	CreatedBy  string // 创建人ID
	DeptId     string // 所属部门ID
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
	DeletedAt  string // 软删除时间
}

// playCategoryColumns holds the columns for the table play_category.
var playCategoryColumns = PlayCategoryColumns{
	Id:         "id",
	ParentId:   "parent_id",
	Title:      "title",
	Icon:       "icon",
	CoverImage: "cover_image",
	Sort:       "sort",
	Status:     "status",
	CreatedBy:  "created_by",
	DeptId:     "dept_id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// NewPlayCategoryDao creates and returns a new DAO object for table data access.
func NewPlayCategoryDao(handlers ...gdb.ModelHandler) *PlayCategoryDao {
	return &PlayCategoryDao{
		group:    "default",
		table:    "play_category",
		columns:  playCategoryColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayCategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayCategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayCategoryDao) Columns() PlayCategoryColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayCategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayCategoryDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayCategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
