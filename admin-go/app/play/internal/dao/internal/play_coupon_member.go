// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayCouponMemberDao is the data access object for the table play_coupon_member.
type PlayCouponMemberDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  PlayCouponMemberColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// PlayCouponMemberColumns defines and stores column names for the table play_coupon_member.
type PlayCouponMemberColumns struct {
	Id        string // 记录ID（Snowflake）
	CouponId  string // 优惠券模板ID
	MemberId  string // 会员ID
	OrderId   string // 使用的订单ID（0表示未使用）
	UseStatus string // 使用状态:0=未使用,1=已使用,2=已过期
	ClaimAt   string // 领取时间
	UseAt     string // 使用时间
	ExpireAt  string // 过期时间
	CreatedBy string // 创建人ID
	DeptId    string // 所属部门ID
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 软删除时间
}

// playCouponMemberColumns holds the columns for the table play_coupon_member.
var playCouponMemberColumns = PlayCouponMemberColumns{
	Id:        "id",
	CouponId:  "coupon_id",
	MemberId:  "member_id",
	OrderId:   "order_id",
	UseStatus: "use_status",
	ClaimAt:   "claim_at",
	UseAt:     "use_at",
	ExpireAt:  "expire_at",
	CreatedBy: "created_by",
	DeptId:    "dept_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewPlayCouponMemberDao creates and returns a new DAO object for table data access.
func NewPlayCouponMemberDao(handlers ...gdb.ModelHandler) *PlayCouponMemberDao {
	return &PlayCouponMemberDao{
		group:    "default",
		table:    "play_coupon_member",
		columns:  playCouponMemberColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayCouponMemberDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayCouponMemberDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayCouponMemberDao) Columns() PlayCouponMemberColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayCouponMemberDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayCouponMemberDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayCouponMemberDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
