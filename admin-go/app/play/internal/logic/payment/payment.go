package payment

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
	service.RegisterPayment(New())
}

func New() *sPayment {
	return &sPayment{}
}

type sPayment struct{}

// Create 创建æ”¯ä»˜è®°å½•è¡¨
func (s *sPayment) Create(ctx context.Context, in *model.PaymentCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayPayment.Ctx(ctx).Data(g.Map{
		dao.PlayPayment.Columns().Id:        id,
		dao.PlayPayment.Columns().OrderId: in.OrderID,
		dao.PlayPayment.Columns().MemberId: in.MemberID,
		dao.PlayPayment.Columns().PaymentNo: in.PaymentNo,
		dao.PlayPayment.Columns().TradeNo: in.TradeNo,
		dao.PlayPayment.Columns().PayType: in.PayType,
		dao.PlayPayment.Columns().PayAmount: in.PayAmount,
		dao.PlayPayment.Columns().PayStatus: in.PayStatus,
		dao.PlayPayment.Columns().PayAt: in.PayAt,
		dao.PlayPayment.Columns().RefundAt: in.RefundAt,
		dao.PlayPayment.Columns().RefundAmount: in.RefundAmount,
		dao.PlayPayment.Columns().CallbackContent: in.CallbackContent,
		dao.PlayPayment.Columns().CreatedAt: gtime.Now(),
		dao.PlayPayment.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新æ”¯ä»˜è®°å½•è¡¨
func (s *sPayment) Update(ctx context.Context, in *model.PaymentUpdateInput) error {
	data := g.Map{
		dao.PlayPayment.Columns().OrderId: in.OrderID,
		dao.PlayPayment.Columns().MemberId: in.MemberID,
		dao.PlayPayment.Columns().PaymentNo: in.PaymentNo,
		dao.PlayPayment.Columns().TradeNo: in.TradeNo,
		dao.PlayPayment.Columns().PayType: in.PayType,
		dao.PlayPayment.Columns().PayAmount: in.PayAmount,
		dao.PlayPayment.Columns().PayStatus: in.PayStatus,
		dao.PlayPayment.Columns().PayAt: in.PayAt,
		dao.PlayPayment.Columns().RefundAt: in.RefundAt,
		dao.PlayPayment.Columns().RefundAmount: in.RefundAmount,
		dao.PlayPayment.Columns().CallbackContent: in.CallbackContent,
		dao.PlayPayment.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayPayment.Ctx(ctx).Where(dao.PlayPayment.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除æ”¯ä»˜è®°å½•è¡¨
func (s *sPayment) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayPayment.Ctx(ctx).Where(dao.PlayPayment.Columns().Id, id).Data(g.Map{
		dao.PlayPayment.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取æ”¯ä»˜è®°å½•è¡¨详情
func (s *sPayment) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.PaymentDetailOutput, err error) {
	out = &model.PaymentDetailOutput{}
	err = dao.PlayPayment.Ctx(ctx).Where(dao.PlayPayment.Columns().Id, id).Where(dao.PlayPayment.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	return
}

// List 获取æ”¯ä»˜è®°å½•è¡¨列表
func (s *sPayment) List(ctx context.Context, in *model.PaymentListInput) (list []*model.PaymentListOutput, total int, err error) {
	m := dao.PlayPayment.Ctx(ctx).Where(dao.PlayPayment.Columns().DeletedAt, nil)
	if in.PayType > 0 {
		m = m.Where(dao.PlayPayment.Columns().PayType, in.PayType)
	}
	if in.PayStatus > 0 {
		m = m.Where(dao.PlayPayment.Columns().PayStatus, in.PayStatus)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayPayment.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	return
}

