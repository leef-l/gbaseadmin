package playapi

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "gbaseadmin/app/play/api/playapi/v1"
	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/model/entity"
)

type sPlayapiBanner struct{}

func (s *sPlayapiBanner) List(ctx context.Context, req *v1.BannerListReq) (list []v1.BannerListItem, err error) {
	bc := dao.PlayBanner.Columns()

	var entities []entity.PlayBanner
	err = dao.PlayBanner.Ctx(ctx).
		Where(bc.Status, 1).
		WhereNull(bc.DeletedAt).
		Where("("+bc.StartTime+" IS NULL OR "+bc.StartTime+" <= NOW())").
		Where("("+bc.EndTime+" IS NULL OR "+bc.EndTime+" >= NOW())").
		OrderDesc(bc.Sort).
		OrderDesc(bc.CreatedAt).
		Scan(&entities)
	if err != nil {
		return
	}

	list = make([]v1.BannerListItem, 0, len(entities))
	for _, e := range entities {
		list = append(list, v1.BannerListItem{
			BannerID:  gconv.String(e.Id),
			Title:     e.Title,
			Image:     e.Image,
			LinkType:  e.LinkType,
			LinkValue: e.LinkValue,
		})
	}
	return
}
