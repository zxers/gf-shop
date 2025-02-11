package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type AddSecKillActivityInput struct {
	Title        string
	GoodsId      uint
	SecKillPrice float64
	TotalStock   uint
	StartTime    *gtime.Time
	EndTime      *gtime.Time
}

type AddSecKillActivityOutput struct {
	Id uint
}

type SecKillActivityUpdateInput struct {
	Id           uint
	Title        string
	GoodsId      uint
	SecKillPrice float64
	TotalStock   uint
	StartTime    *gtime.Time
	EndTime      *gtime.Time
}

type SecKillActivityDetailInput struct {
	Id uint
}

type SecKillActivityDetailOutput struct {
	Id           uint
	Title        string
	GoodsId      uint
	SecKillPrice float64
	TotalStock   uint
	StartTime    *gtime.Time
	EndTime      *gtime.Time
}
