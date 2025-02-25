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
	IOrder interface {
		// 下单
		Add(ctx context.Context, in model.OrderAddInput) (out *model.OrderAddOutput, err error)
		List(ctx context.Context, in model.OrderListInput) (out *model.OrderListOutput, err error)
		Detail(ctx context.Context, in model.OrderDetailInput) (out *model.OrderDetailOutput, err error)
	}
)

var (
	localOrder IOrder
)

func Order() IOrder {
	if localOrder == nil {
		panic("implement not found for interface IOrder, forgot register?")
	}
	return localOrder
}

func RegisterOrder(i IOrder) {
	localOrder = i
}
