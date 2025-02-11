package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SecKillActivityReq struct {
	g.Meta     `path:"/seckill" tags:"seckill" method:"post" summary:"You first admin api"`
	ActivityId int  `json:"activity_id"    v:"required#密码不能为空" dc:"密码"`
	GoodsId    uint `json:"goods_id"    v:"required#密码不能为空" dc:"密码"`
}

type SecKillActivityRes struct {
}
