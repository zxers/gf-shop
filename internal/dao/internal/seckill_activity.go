// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SeckillActivityDao is the data access object for the table seckill_activity.
type SeckillActivityDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of the current DAO.
	columns SeckillActivityColumns // columns contains all the column names of Table for convenient usage.
}

// SeckillActivityColumns defines and stores column names for the table seckill_activity.
type SeckillActivityColumns struct {
	Id           string // 活动ID
	Title        string // 活动标题
	GoodsId      string // 商品ID
	SeckillPrice string // 秒杀价
	TotalStock   string // 总库存
	StartTime    string // 开始时间
	EndTime      string // 结束时间
	CreatedAt    string //
	UpdatedAt    string //
}

// seckillActivityColumns holds the columns for the table seckill_activity.
var seckillActivityColumns = SeckillActivityColumns{
	Id:           "id",
	Title:        "title",
	GoodsId:      "goods_id",
	SeckillPrice: "seckill_price",
	TotalStock:   "total_stock",
	StartTime:    "start_time",
	EndTime:      "end_time",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewSeckillActivityDao creates and returns a new DAO object for table data access.
func NewSeckillActivityDao() *SeckillActivityDao {
	return &SeckillActivityDao{
		group:   "default",
		table:   "seckill_activity",
		columns: seckillActivityColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SeckillActivityDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SeckillActivityDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SeckillActivityDao) Columns() SeckillActivityColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SeckillActivityDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SeckillActivityDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SeckillActivityDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
