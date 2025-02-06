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
	ISeckill interface {
		Create(ctx context.Context, in model.AddSecKillActivityInput) (out model.AddSecKillActivityOutput, err error)
	}
)

var (
	localSeckill ISeckill
)

func Seckill() ISeckill {
	if localSeckill == nil {
		panic("implement not found for interface ISeckill, forgot register?")
	}
	return localSeckill
}

func RegisterSeckill(i ISeckill) {
	localSeckill = i
}
