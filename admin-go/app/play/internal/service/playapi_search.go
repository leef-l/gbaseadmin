package service

import (
	"context"

	v1 "gbaseadmin/app/play/api/playapi/v1"
)

type IPlayapiSearch interface {
	Search(ctx context.Context, keyword, searchType string, page, pageSize int) (coachList []v1.SearchCoachItem, coachTotal int, goodsList []v1.SearchGoodsItem, goodsTotal int, err error)
}

var localPlayapiSearch IPlayapiSearch

func PlayapiSearch() IPlayapiSearch {
	return localPlayapiSearch
}

func RegisterPlayapiSearch(i IPlayapiSearch) {
	localPlayapiSearch = i
}
