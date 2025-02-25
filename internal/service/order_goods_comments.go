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
	IOrderGoodsComments interface {
		Add(ctx context.Context, in model.AddOrderGoodsCommentsInput) (out model.AddOrderGoodsCommentsOutput, err error)
		Delete(ctx context.Context, in model.DelOrderGoodsCommentsInput) (out model.DelOrderGoodsCommentsOutput, err error)
	}
)

var (
	localOrderGoodsComments IOrderGoodsComments
)

func OrderGoodsComments() IOrderGoodsComments {
	if localOrderGoodsComments == nil {
		panic("implement not found for interface IOrderGoodsComments, forgot register?")
	}
	return localOrderGoodsComments
}

func RegisterOrderGoodsComments(i IOrderGoodsComments) {
	localOrderGoodsComments = i
}
