package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type IBanner interface {
	Create(ctx context.Context, in *model.BannerCreateInput) error
	Update(ctx context.Context, in *model.BannerUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	BatchDelete(ctx context.Context, ids []snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.BannerDetailOutput, err error)
	List(ctx context.Context, in *model.BannerListInput) (list []*model.BannerListOutput, total int, err error)
	Export(ctx context.Context, in *model.BannerListInput) (list []*model.BannerListOutput, err error)
}

var localBanner IBanner

func Banner() IBanner {
	return localBanner
}

func RegisterBanner(i IBanner) {
	localBanner = i
}
