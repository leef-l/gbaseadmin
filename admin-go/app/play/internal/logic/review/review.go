package review

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
	"gbaseadmin/utility/snowflake"
)

func init() {
	service.RegisterReview(New())
}

func New() *sReview {
	return &sReview{}
}

type sReview struct{}

// Create 创建评价表
func (s *sReview) Create(ctx context.Context, in *model.ReviewCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayReview.Ctx(ctx).Data(g.Map{
		dao.PlayReview.Columns().Id:        id,
		dao.PlayReview.Columns().OrderId: in.OrderID,
		dao.PlayReview.Columns().MemberId: in.MemberID,
		dao.PlayReview.Columns().CoachId: in.CoachID,
		dao.PlayReview.Columns().Score: in.Score,
		dao.PlayReview.Columns().ReviewContent: in.ReviewContent,
		dao.PlayReview.Columns().ReviewImage: in.ReviewImage,
		dao.PlayReview.Columns().ReplyContent: in.ReplyContent,
		dao.PlayReview.Columns().ReplyAt: in.ReplyAt,
		dao.PlayReview.Columns().IsAnonymous: in.IsAnonymous,
		dao.PlayReview.Columns().Status: in.Status,
		dao.PlayReview.Columns().CreatedAt: gtime.Now(),
		dao.PlayReview.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新评价表
func (s *sReview) Update(ctx context.Context, in *model.ReviewUpdateInput) error {
	data := g.Map{
		dao.PlayReview.Columns().OrderId: in.OrderID,
		dao.PlayReview.Columns().MemberId: in.MemberID,
		dao.PlayReview.Columns().CoachId: in.CoachID,
		dao.PlayReview.Columns().Score: in.Score,
		dao.PlayReview.Columns().ReviewContent: in.ReviewContent,
		dao.PlayReview.Columns().ReviewImage: in.ReviewImage,
		dao.PlayReview.Columns().ReplyContent: in.ReplyContent,
		dao.PlayReview.Columns().ReplyAt: in.ReplyAt,
		dao.PlayReview.Columns().IsAnonymous: in.IsAnonymous,
		dao.PlayReview.Columns().Status: in.Status,
		dao.PlayReview.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayReview.Ctx(ctx).Where(dao.PlayReview.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除评价表
func (s *sReview) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayReview.Ctx(ctx).Where(dao.PlayReview.Columns().Id, id).Data(g.Map{
		dao.PlayReview.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取评价表详情
func (s *sReview) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.ReviewDetailOutput, err error) {
	out = &model.ReviewDetailOutput{}
	err = dao.PlayReview.Ctx(ctx).Where(dao.PlayReview.Columns().Id, id).Where(dao.PlayReview.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	return
}

// List 获取评价表列表
func (s *sReview) List(ctx context.Context, in *model.ReviewListInput) (list []*model.ReviewListOutput, total int, err error) {
	m := dao.PlayReview.Ctx(ctx).Where(dao.PlayReview.Columns().DeletedAt, nil)
	if in.IsAnonymous > 0 {
		m = m.Where(dao.PlayReview.Columns().IsAnonymous, in.IsAnonymous)
	}
	if in.Status > 0 {
		m = m.Where(dao.PlayReview.Columns().Status, in.Status)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayReview.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	return
}

