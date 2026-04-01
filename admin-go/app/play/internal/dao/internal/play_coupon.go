// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayCouponDao is the data access object for the table play_coupon.
type PlayCouponDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PlayCouponColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PlayCouponColumns defines and stores column names for the table play_coupon.
type PlayCouponColumns struct {
	Id           string // 优惠券ID（Snowflake）
	Title        string // 优惠券名称
	Type         string // 优惠券类型:1=满减券,2=折扣券,3=无门槛券
	IsNewMember  string // 是否新人专享:0=否,1=是
	FaceValue    string // 面值（分，满减/无门槛时为抵扣额，折扣时为折扣值如 85=8.5折）
	MinAmount    string // 最低消费金额（分，0表示无门槛）
	TotalNum     string // 发放总量（0表示不限）
	UsedNum      string // 已使用数量
	ClaimNum     string // 已领取数量
	PerLimit     string // 每人限领张数
	ValidStartAt string // 有效期开始时间
	ValidEndAt   string // 有效期结束时间
	Sort         string // 排序（升序）
	Status       string // 状态:0=关闭,1=开启
	CreatedBy    string // 创建人ID
	DeptId       string // 所属部门ID
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
	DeletedAt    string // 软删除时间
}

// playCouponColumns holds the columns for the table play_coupon.
var playCouponColumns = PlayCouponColumns{
	Id:           "id",
	Title:        "title",
	Type:         "type",
	IsNewMember:  "is_new_member",
	FaceValue:    "face_value",
	MinAmount:    "min_amount",
	TotalNum:     "total_num",
	UsedNum:      "used_num",
	ClaimNum:     "claim_num",
	PerLimit:     "per_limit",
	ValidStartAt: "valid_start_at",
	ValidEndAt:   "valid_end_at",
	Sort:         "sort",
	Status:       "status",
	CreatedBy:    "created_by",
	DeptId:       "dept_id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

// NewPlayCouponDao creates and returns a new DAO object for table data access.
func NewPlayCouponDao(handlers ...gdb.ModelHandler) *PlayCouponDao {
	return &PlayCouponDao{
		group:    "default",
		table:    "play_coupon",
		columns:  playCouponColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayCouponDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayCouponDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayCouponDao) Columns() PlayCouponColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayCouponDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayCouponDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayCouponDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
