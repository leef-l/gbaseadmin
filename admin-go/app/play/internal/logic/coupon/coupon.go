package coupon

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
	service.RegisterCoupon(New())
}

func New() *sCoupon {
	return &sCoupon{}
}

type sCoupon struct{}

// Create 创建ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨
func (s *sCoupon) Create(ctx context.Context, in *model.CouponCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayCoupon.Ctx(ctx).Data(g.Map{
		dao.PlayCoupon.Columns().Id:        id,
		dao.PlayCoupon.Columns().Title: in.Title,
		dao.PlayCoupon.Columns().Type: in.Type,
		dao.PlayCoupon.Columns().IsNewMember: in.IsNewMember,
		dao.PlayCoupon.Columns().FaceValue: in.FaceValue,
		dao.PlayCoupon.Columns().MinAmount: in.MinAmount,
		dao.PlayCoupon.Columns().TotalNum: in.TotalNum,
		dao.PlayCoupon.Columns().UsedNum: in.UsedNum,
		dao.PlayCoupon.Columns().ClaimNum: in.ClaimNum,
		dao.PlayCoupon.Columns().PerLimit: in.PerLimit,
		dao.PlayCoupon.Columns().ValidStartAt: in.ValidStartAt,
		dao.PlayCoupon.Columns().ValidEndAt: in.ValidEndAt,
		dao.PlayCoupon.Columns().Sort: in.Sort,
		dao.PlayCoupon.Columns().Status: in.Status,
		dao.PlayCoupon.Columns().CreatedAt: gtime.Now(),
		dao.PlayCoupon.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨
func (s *sCoupon) Update(ctx context.Context, in *model.CouponUpdateInput) error {
	data := g.Map{
		dao.PlayCoupon.Columns().Title: in.Title,
		dao.PlayCoupon.Columns().Type: in.Type,
		dao.PlayCoupon.Columns().IsNewMember: in.IsNewMember,
		dao.PlayCoupon.Columns().FaceValue: in.FaceValue,
		dao.PlayCoupon.Columns().MinAmount: in.MinAmount,
		dao.PlayCoupon.Columns().TotalNum: in.TotalNum,
		dao.PlayCoupon.Columns().UsedNum: in.UsedNum,
		dao.PlayCoupon.Columns().ClaimNum: in.ClaimNum,
		dao.PlayCoupon.Columns().PerLimit: in.PerLimit,
		dao.PlayCoupon.Columns().ValidStartAt: in.ValidStartAt,
		dao.PlayCoupon.Columns().ValidEndAt: in.ValidEndAt,
		dao.PlayCoupon.Columns().Sort: in.Sort,
		dao.PlayCoupon.Columns().Status: in.Status,
		dao.PlayCoupon.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayCoupon.Ctx(ctx).Where(dao.PlayCoupon.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨
func (s *sCoupon) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayCoupon.Ctx(ctx).Where(dao.PlayCoupon.Columns().Id, id).Data(g.Map{
		dao.PlayCoupon.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨详情
func (s *sCoupon) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.CouponDetailOutput, err error) {
	out = &model.CouponDetailOutput{}
	err = dao.PlayCoupon.Ctx(ctx).Where(dao.PlayCoupon.Columns().Id, id).Where(dao.PlayCoupon.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	return
}

// List 获取ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨列表
func (s *sCoupon) List(ctx context.Context, in *model.CouponListInput) (list []*model.CouponListOutput, total int, err error) {
	m := dao.PlayCoupon.Ctx(ctx).Where(dao.PlayCoupon.Columns().DeletedAt, nil)
	if in.Type > 0 {
		m = m.Where(dao.PlayCoupon.Columns().Type, in.Type)
	}
	if in.IsNewMember > 0 {
		m = m.Where(dao.PlayCoupon.Columns().IsNewMember, in.IsNewMember)
	}
	if in.Status > 0 {
		m = m.Where(dao.PlayCoupon.Columns().Status, in.Status)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayCoupon.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	return
}

