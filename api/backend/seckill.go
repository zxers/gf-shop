package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type SecKillActivityReq struct {
	g.Meta       `path:"/admin/add" tags:"Admin" method:"post" summary:"You first admin api"`
	Title        string      `json:"title" v:"required#用户名不能为空" dc:"用户名"`
	GoodsId      uint        `json:"goodsId"    v:"required#密码不能为空" dc:"密码"`
	SecKillPrice float64     `json:"sec_kill_price"    v:"required#密码不能为空" dc:"密码"`
	TotalStock   uint        `json:"total_stock"    v:"required#密码不能为空" dc:"密码"`
	StartTime    *gtime.Time `json:"start_time"    dc:"角色ids"`
	EndTime      *gtime.Time `json:"end_time"    dc:"是否超级管理员"`
}

type SecKillActivityRes struct {
	Id int `json:"id"`
}
