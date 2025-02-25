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
	IAddress interface {
		Add(ctx context.Context, in model.AddAddressInput) (out *model.AddAddressOutput, err error)
		Update(ctx context.Context, in model.UpdateAddressInput) (err error)
		Delete(ctx context.Context, id int) (err error)
		Page(ctx context.Context, in model.PageAddressInput) (out *model.PageAddressOutput, err error)
		// 客户端获取省市县区地址
		GetCityList(ctx context.Context) (out *model.CityAddressListOutput, err error)
	}
)

var (
	localAddress IAddress
)

func Address() IAddress {
	if localAddress == nil {
		panic("implement not found for interface IAddress, forgot register?")
	}
	return localAddress
}

func RegisterAddress(i IAddress) {
	localAddress = i
}
