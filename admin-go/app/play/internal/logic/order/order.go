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

// Create 创建订单表
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

// Update 更新订单表
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

// Delete 软删除订单表
func (s *sOrder) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayOrder.Ctx(ctx).Where(dao.PlayOrder.Columns().Id, id).Data(g.Map{
		dao.PlayOrder.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取订单表详情
func (s *sOrder) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.OrderDetailOutput, err error) {
	out = &model.OrderDetailOutput{}
	err = dao.PlayOrder.Ctx(ctx).Where(dao.PlayOrder.Columns().Id, id).Where(dao.PlayOrder.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	// 查询店铺ID（0表示无店铺）关联显示
	if out.ShopID != 0 {
		val, err := g.DB().Ctx(ctx).Model("play_shop").Where("id", out.ShopID).Where("deleted_at", nil).Value("title")
		if err == nil {
			out.ShopTitle = val.String()
		}
	}
	return
}

// List 获取订单表列表
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
			val, err := g.DB().Ctx(ctx).Model("play_shop").Where("id", item.ShopID).Where("deleted_at", nil).Value("title")
			if err == nil {
				item.ShopTitle = val.String()
			}
		}
	}
	return
}

