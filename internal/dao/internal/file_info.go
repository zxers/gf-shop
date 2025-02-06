// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FileInfoDao is the data access object for the table file_info.
type FileInfoDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns FileInfoColumns // columns contains all the column names of Table for convenient usage.
}

// FileInfoColumns defines and stores column names for the table file_info.
type FileInfoColumns struct {
	Id        string //
	Name      string // 图片名称
	Src       string // 本地文件存储路径
	Url       string // URL地址
	UserId    string // 用户id
	CreatedAt string //
	UpdatedAt string //
}

// fileInfoColumns holds the columns for the table file_info.
var fileInfoColumns = FileInfoColumns{
	Id:        "id",
	Name:      "name",
	Src:       "src",
	Url:       "url",
	UserId:    "user_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewFileInfoDao creates and returns a new DAO object for table data access.
func NewFileInfoDao() *FileInfoDao {
	return &FileInfoDao{
		group:   "default",
		table:   "file_info",
		columns: fileInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FileInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FileInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FileInfoDao) Columns() FileInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FileInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FileInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FileInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
