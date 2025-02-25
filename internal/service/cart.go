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
	ICart interface {
		Add(ctx context.Context, in model.AddCartInput) (out model.AddCartOutput, err error)
		Delete(ctx context.Context, in model.DeleteCartInput) (out model.DeleteCartOutput, err error)
		List(ctx context.Context, in model.ListCartInput) (out *model.ListCartOutput, err error)
	}
)

var (
	localCart ICart
)

func Cart() ICart {
	if localCart == nil {
		panic("implement not found for interface ICart, forgot register?")
	}
	return localCart
}

func RegisterCart(i ICart) {
	localCart = i
}
