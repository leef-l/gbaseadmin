package order

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
	service.RegisterOrder(New())
}

func New() *sOrder {
	return &sOrder{}
}

type sOrder struct{}

// Create 创建è®¢å•è¡¨
func (s *sOrder) Create(ctx context.Context, in *model.OrderCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayOrder.Ctx(ctx).Data(g.Map{
		dao.PlayOrder.Columns().Id:        id,
		dao.PlayOrder.Columns().OrderNo: in.OrderNo,
		dao.PlayOrder.Columns().MemberId: in.MemberID,
		dao.PlayOrder.Columns().CoachId: in.CoachID,
		dao.PlayOrder.Columns().ShopId: in.ShopID,
		dao.PlayOrder.Columns().GoodsId: in.GoodsID,
		dao.PlayOrder.Columns().GoodsTitle: in.GoodsTitle,
		dao.PlayOrder.Columns().GoodsPrice: in.GoodsPrice,
		dao.PlayOrder.Columns().Quantity: in.Quantity,
		dao.PlayOrder.Columns().TotalAmount: in.TotalAmount,
		dao.PlayOrder.Columns().DiscountAmount: in.DiscountAmount,
		dao.PlayOrder.Columns().CouponAmount: in.CouponAmount,
		dao.PlayOrder.Columns().PayAmount: in.PayAmount,
		dao.PlayOrder.Columns().CouponMemberId: in.CouponMemberID,
		dao.PlayOrder.Columns().PayType: in.PayType,
		dao.PlayOrder.Columns().OrderStatus: in.OrderStatus,
		dao.PlayOrder.Columns().PayAt: in.PayAt,
		dao.PlayOrder.Columns().StartAt: in.StartAt,
		dao.PlayOrder.Columns().FinishAt: in.FinishAt,
		dao.PlayOrder.Columns().CancelAt: in.CancelAt,
		dao.PlayOrder.Columns().CancelReason: in.CancelReason,
		dao.PlayOrder.Columns().Remark: in.Remark,
		dao.PlayOrder.Columns().CreatedAt: gtime.Now(),
		dao.PlayOrder.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新è®¢å•è¡¨
func (s *sOrder) Update(ctx context.Context, in *model.OrderUpdateInput) error {
	data := g.Map{
		dao.PlayOrder.Columns().OrderNo: in.OrderNo,
		dao.PlayOrder.Columns().MemberId: in.MemberID,
		dao.PlayOrder.Columns().CoachId: in.CoachID,
		dao.PlayOrder.Columns().ShopId: in.ShopID,
		dao.PlayOrder.Columns().GoodsId: in.GoodsID,
		dao.PlayOrder.Columns().GoodsTitle: in.GoodsTitle,
		dao.PlayOrder.Columns().GoodsPrice: in.GoodsPrice,
		dao.PlayOrder.Columns().Quantity: in.Quantity,
		dao.PlayOrder.Columns().TotalAmount: in.TotalAmount,
		dao.PlayOrder.Columns().DiscountAmount: in.DiscountAmount,
		dao.PlayOrder.Columns().CouponAmount: in.CouponAmount,
		dao.PlayOrder.Columns().PayAmount: in.PayAmount,
		dao.PlayOrder.Columns().CouponMemberId: in.CouponMemberID,
		dao.PlayOrder.Columns().PayType: in.PayType,
		dao.PlayOrder.Columns().OrderStatus: in.OrderStatus,
		dao.PlayOrder.Columns().PayAt: in.PayAt,
		dao.PlayOrder.Columns().StartAt: in.StartAt,
		dao.PlayOrder.Columns().FinishAt: in.FinishAt,
		dao.PlayOrder.Columns().CancelAt: in.CancelAt,
		dao.PlayOrder.Columns().CancelReason: in.CancelReason,
		dao.PlayOrder.Columns().Remark: in.Remark,
		dao.PlayOrder.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayOrder.Ctx(ctx).Where(dao.PlayOrder.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除è®¢å•è¡¨
func (s *sOrder) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayOrder.Ctx(ctx).Where(dao.PlayOrder.Columns().Id, id).Data(g.Map{
		dao.PlayOrder.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取è®¢å•è¡¨详情
func (s *sOrder) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.OrderDetailOutput, err error) {
	out = &model.OrderDetailOutput{}
	err = dao.PlayOrder.Ctx(ctx).Where(dao.PlayOrder.Columns().Id, id).Where(dao.PlayOrder.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	// 查询åº—é“ºID关联显示
	if out.ShopID != 0 {
		val, _ := g.DB().Ctx(ctx).Model("play_shop").Where("id", out.ShopID).Where("deleted_at", nil).Value("title")
		out.ShopTitle = val.String()
	}
	// 查询å•†å“ID关联显示
	if out.GoodsID != 0 {
		val, _ := g.DB().Ctx(ctx).Model("play_goods").Where("id", out.GoodsID).Where("deleted_at", nil).Value("title")
		out.GoodsTitle = val.String()
	}
	return
}

// List 获取è®¢å•è¡¨列表
func (s *sOrder) List(ctx context.Context, in *model.OrderListInput) (list []*model.OrderListOutput, total int, err error) {
	m := dao.PlayOrder.Ctx(ctx).Where(dao.PlayOrder.Columns().DeletedAt, nil)
	if in.PayType > 0 {
		m = m.Where(dao.PlayOrder.Columns().PayType, in.PayType)
	}
	if in.OrderStatus > 0 {
		m = m.Where(dao.PlayOrder.Columns().OrderStatus, in.OrderStatus)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayOrder.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	// 填充关联显示字段
	for _, item := range list {
		if item.ShopID != 0 {
			val, _ := g.DB().Ctx(ctx).Model("play_shop").Where("id", item.ShopID).Where("deleted_at", nil).Value("title")
			item.ShopTitle = val.String()
		}
		if item.GoodsID != 0 {
			val, _ := g.DB().Ctx(ctx).Model("play_goods").Where("id", item.GoodsID).Where("deleted_at", nil).Value("title")
			item.GoodsTitle = val.String()
		}
	}
	return
}

