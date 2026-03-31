// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UploadFileDao is the data access object for the table upload_file.
type UploadFileDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UploadFileColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UploadFileColumns defines and stores column names for the table upload_file.
type UploadFileColumns struct {
	Id        string // ID
	DirId     string // æ‰€å±žç›®å½•
	Name      string // æ–‡ä»¶åç§°
	Url       string // æ–‡ä»¶åœ°å€
	Ext       string // æ–‡ä»¶æ‰©å±•å
	Size      string // æ–‡ä»¶å¤§å°
	Mime      string // MIMEç±»åž‹
	Storage   string // å­˜å‚¨ç±»åž‹:1=æœ¬åœ°,2=é˜¿é‡Œäº‘OSS,3=è…¾è®¯äº‘COS
	IsImage   string // æ˜¯å¦å›¾ç‰‡:0=å¦,1=æ˜¯
	CreatedAt string // åˆ›å»ºæ—¶é—´
	UpdatedAt string // æ›´æ–°æ—¶é—´
	DeletedAt string // åˆ é™¤æ—¶é—´
	CreatedBy string // åˆ›å»ºäºº
	DeptId    string // éƒ¨é—¨ID
}

// uploadFileColumns holds the columns for the table upload_file.
var uploadFileColumns = UploadFileColumns{
	Id:        "id",
	DirId:     "dir_id",
	Name:      "name",
	Url:       "url",
	Ext:       "ext",
	Size:      "size",
	Mime:      "mime",
	Storage:   "storage",
	IsImage:   "is_image",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	CreatedBy: "created_by",
	DeptId:    "dept_id",
}

// NewUploadFileDao creates and returns a new DAO object for table data access.
func NewUploadFileDao(handlers ...gdb.ModelHandler) *UploadFileDao {
	return &UploadFileDao{
		group:    "default",
		table:    "upload_file",
		columns:  uploadFileColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UploadFileDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UploadFileDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UploadFileDao) Columns() UploadFileColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UploadFileDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UploadFileDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UploadFileDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
