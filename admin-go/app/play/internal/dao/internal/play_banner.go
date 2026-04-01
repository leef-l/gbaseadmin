// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayBannerDao is the data access object for the table play_banner.
type PlayBannerDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PlayBannerColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PlayBannerColumns defines and stores column names for the table play_banner.
type PlayBannerColumns struct {
	Id        string // Banner ID
	Title     string // Banner标题
	Image     string // 图片URL
	LinkType  string // 跳转类型: 1内页 2外链 3活动页 4商品页 5陪玩师页 6唤醒App
	LinkValue string // 跳转值(页面路径/URL/业务ID/App Scheme)
	Sort      string // 排序(越大越前)
	Status    string // 状态: 0禁用 1启用
	StartTime string // 生效开始时间
	EndTime   string // 生效结束时间
	Remark    string // 备注
	CreatedBy string // 创建人
	DeptId    string // 部门ID
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// playBannerColumns holds the columns for the table play_banner.
var playBannerColumns = PlayBannerColumns{
	Id:        "id",
	Title:     "title",
	Image:     "image",
	LinkType:  "link_type",
	LinkValue: "link_value",
	Sort:      "sort",
	Status:    "status",
	StartTime: "start_time",
	EndTime:   "end_time",
	Remark:    "remark",
	CreatedBy: "created_by",
	DeptId:    "dept_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewPlayBannerDao creates and returns a new DAO object for table data access.
func NewPlayBannerDao(handlers ...gdb.ModelHandler) *PlayBannerDao {
	return &PlayBannerDao{
		group:    "default",
		table:    "play_banner",
		columns:  playBannerColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayBannerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayBannerDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayBannerDao) Columns() PlayBannerColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayBannerDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayBannerDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayBannerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
