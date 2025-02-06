// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SeckillActivity is the golang structure for table seckill_activity.
type SeckillActivity struct {
	Id           uint        `json:"id"           orm:"id"            description:"活动ID"`
	Title        string      `json:"title"        orm:"title"         description:"活动标题"`
	GoodsId      uint        `json:"goodsId"      orm:"goods_id"      description:"商品ID"`
	SeckillPrice float64     `json:"seckillPrice" orm:"seckill_price" description:"秒杀价"`
	TotalStock   uint        `json:"totalStock"   orm:"total_stock"   description:"总库存"`
	StartTime    *gtime.Time `json:"startTime"    orm:"start_time"    description:"开始时间"`
	EndTime      *gtime.Time `json:"endTime"      orm:"end_time"      description:"结束时间"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:""`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:""`
}
