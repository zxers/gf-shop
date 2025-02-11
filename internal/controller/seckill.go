package controller

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
	"goframe-shop-v2/api/backend"
	"goframe-shop-v2/api/frontend"
	"goframe-shop-v2/internal/consts"
	"goframe-shop-v2/internal/model"
	"goframe-shop-v2/internal/service"
)

var Seckill = cSeckill{}

type cSeckill struct{}

func (a *cSeckill) Create(ctx context.Context, req *backend.SecKillActivityReq) (res *backend.SecKillActivityRes, err error) {
	out, err := service.Seckill().Create(ctx, model.AddSecKillActivityInput{
		Title:        req.Title,
		GoodsId:      req.GoodsId,
		SecKillPrice: req.SecKillPrice,
		TotalStock:   req.TotalStock,
	})
	if err != nil {
		return nil, err
	}
	return &backend.SecKillActivityRes{Id: int(out.Id)}, nil
}

func (a *cSeckill) SecKill(ctx context.Context, req *frontend.SecKillActivityReq) (res *frontend.SecKillActivityRes, err error) {
	sa, err := service.Seckill().Detail(ctx, model.SecKillActivityDetailInput{
		Id: req.GoodsId,
	})
	if err != nil {
		return nil, err
	}

	if sa.TotalStock <= 0 {
		return nil, errors.New("库存不足")
	}
	// 扣减库存
	err = service.Seckill().Update(ctx, model.SecKillActivityUpdateInput{
		Id:         uint(req.ActivityId),
		GoodsId:    req.GoodsId,
		TotalStock: sa.TotalStock - 1,
	})

	if err != nil {
		return nil, err
	}

	// 创建订单
	_, err = service.Order().Add(ctx, model.OrderAddInput{
		UserId: gconv.Uint(ctx.Value(consts.CtxUserId)),
		Number: uuid.NewString(),
	})

	if err != nil {
		return nil, err
	}

	return &frontend.SecKillActivityRes{}, nil
}
