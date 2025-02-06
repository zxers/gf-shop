package seckill

import (
	"context"
	"goframe-shop-v2/internal/dao"
	"goframe-shop-v2/internal/model"
	"goframe-shop-v2/internal/service"
)

type sSeckill struct{}

func init() {
	service.RegisterSeckill(New())
}

func New() *sSeckill {
	return &sSeckill{}
}

func (s *sSeckill) Create(ctx context.Context, in model.AddSecKillActivityInput) (out model.AddSecKillActivityOutput, err error) {
	//插入数据返回id
	lastInsertID, err := dao.SeckillActivity.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.AddSecKillActivityOutput{Id: uint(lastInsertID)}, err
}
