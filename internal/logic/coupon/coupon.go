package position

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-shop-v2/internal/dao"
	"goframe-shop-v2/internal/model"
	"goframe-shop-v2/internal/model/entity"
	"goframe-shop-v2/internal/service"
)

type sCoupon struct{}

func init() {
	service.RegisterCoupon(New())
}

func New() *sCoupon {
	return &sCoupon{}
}

func (s *sCoupon) Create(ctx context.Context, in model.CouponCreateInput) (out model.CouponCreateOutput, err error) {
	lastInsertID, err := dao.CouponInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.CouponCreateOutput{CouponId: uint(lastInsertID)}, err
}

// Delete 删除
func (s *sCoupon) Delete(ctx context.Context, id uint) (err error) {
	_, err = dao.CouponInfo.Ctx(ctx).Where(g.Map{
		dao.CouponInfo.Columns().Id: id,
	}).Delete()
	if err != nil {
		return err
	}
	return
}

// Update 修改
func (s *sCoupon) Update(ctx context.Context, in model.CouponUpdateInput) error {
	_, err := dao.CouponInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.CouponInfo.Columns().Id).
		Where(dao.CouponInfo.Columns().Id, in.Id).
		Update()
	return err
}

// GetList 查询分类列表
func (s *sCoupon) GetList(ctx context.Context, in model.CouponGetListInput) (out *model.CouponGetListOutput, err error) {
	//1.获得*gdb.Model对象，方面后续调用
	m := dao.CouponInfo.Ctx(ctx)
	//2. 实例化响应结构体
	out = &model.CouponGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	//3. 分页查询
	listModel := m.Page(in.Page, in.Size)
	//4. 再查询count，判断有无数据
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	//5. 延迟初始化list切片 确定有数据，再按期望大小初始化切片容量
	out.List = make([]model.CouponGetListOutputItem, 0, in.Size)
	//6. 把查询到的结果赋值到响应结构体中
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// GetList 查询分类全部信息-不翻页
func (s *sCoupon) GetListAll(ctx context.Context, in model.CouponGetListInput) (out *model.CouponGetListOutput, err error) {
	var (
		m = dao.CouponInfo.Ctx(ctx)
	)
	out = &model.CouponGetListOutput{}

	listModel := m
	// 排序方式
	listModel = listModel.OrderDesc(dao.CouponInfo.Columns().Price)

	// 执行查询
	var list []*entity.CouponInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
