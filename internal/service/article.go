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
	IArticle interface {
		Create(ctx context.Context, in model.ArticleCreateInput) (out model.ArticleCreateOutput, err error)
		// Delete 删除
		Delete(ctx context.Context, in model.ArticleDeleteInput) (err error)
		// Update 修改
		Update(ctx context.Context, in model.ArticleUpdateInput) error
		// GetList 查询分类列表
		GetList(ctx context.Context, in model.ArticleGetListInput) (out *model.ArticleGetListOutput, err error)
		// 详情
		Detail(ctx context.Context, in model.ArticleDetailInput) (out *model.ArticleDetailOutput, err error)
	}
)

var (
	localArticle IArticle
)

func Article() IArticle {
	if localArticle == nil {
		panic("implement not found for interface IArticle, forgot register?")
	}
	return localArticle
}

func RegisterArticle(i IArticle) {
	localArticle = i
}
