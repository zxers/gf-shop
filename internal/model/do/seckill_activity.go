// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SeckillActivity is the golang structure of table seckill_activity for DAO operations like Where/Data.
type SeckillActivity struct {
	g.Meta       `orm:"table:seckill_activity, do:true"`
	Id           interface{} // 活动ID
	Title        interface{} // 活动标题
	GoodsId      interface{} // 商品ID
	SeckillPrice interface{} // 秒杀价
	TotalStock   interface{} // 总库存
	StartTime    *gtime.Time // 开始时间
	EndTime      *gtime.Time // 结束时间
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}
