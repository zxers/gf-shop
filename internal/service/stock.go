// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goframe-shop-v2/internal/model"
)

type (
	IStock interface {
		DecrementWithSql(ctx context.Context, in model.DecStockInput) (err error)
		DecrementWithRedis(ctx context.Context, in model.DecStockInput) error
	}
)

var (
	localStock IStock
)

func Stock() IStock {
	if localStock == nil {
		panic("implement not found for interface IStock, forgot register?")
	}
	return localStock
}

func RegisterStock(i IStock) {
	localStock = i
}
