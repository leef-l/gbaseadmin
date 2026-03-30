package playapi

import (
	"context"

	v1 "gbaseadmin/api/playapi/v1"
	"gbaseadmin/app/play/internal/service"
)

var SearchPublic = &cSearchPublic{}

type cSearchPublic struct{}

// Search 综合搜索
func (c *cSearchPublic) Search(ctx context.Context, req *v1.SearchReq) (res *v1.SearchRes, err error) {
	res = &v1.SearchRes{}
	res.CoachList, res.CoachTotal, res.GoodsList, res.GoodsTotal, err = service.PlayapiSearch().Search(ctx, req.Keyword, req.Type, req.Page, req.PageSize)
	return
}
