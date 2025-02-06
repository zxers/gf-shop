// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UntitledTable24Dao is the data access object for the table untitled_table_24.
type UntitledTable24Dao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of the current DAO.
	columns UntitledTable24Columns // columns contains all the column names of Table for convenient usage.
}

// UntitledTable24Columns defines and stores column names for the table untitled_table_24.
type UntitledTable24Columns struct {
	Id string //
}

// untitledTable24Columns holds the columns for the table untitled_table_24.
var untitledTable24Columns = UntitledTable24Columns{
	Id: "id",
}

// NewUntitledTable24Dao creates and returns a new DAO object for table data access.
func NewUntitledTable24Dao() *UntitledTable24Dao {
	return &UntitledTable24Dao{
		group:   "default",
		table:   "untitled_table_24",
		columns: untitledTable24Columns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UntitledTable24Dao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UntitledTable24Dao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UntitledTable24Dao) Columns() UntitledTable24Columns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UntitledTable24Dao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UntitledTable24Dao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UntitledTable24Dao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
