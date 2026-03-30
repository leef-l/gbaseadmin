package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type IOauth interface {
	Create(ctx context.Context, in *model.OauthCreateInput) error
	Update(ctx context.Context, in *model.OauthUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.OauthDetailOutput, err error)
	List(ctx context.Context, in *model.OauthListInput) (list []*model.OauthListOutput, total int, err error)
}

var localOauth IOauth

func Oauth() IOauth {
	return localOauth
}

func RegisterOauth(i IOauth) {
	localOauth = i
}
