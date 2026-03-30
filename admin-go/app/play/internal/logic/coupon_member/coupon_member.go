package coupon_member

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
	service.RegisterCouponMember(New())
}

func New() *sCouponMember {
	return &sCouponMember{}
}

type sCouponMember struct{}

// Create 创建ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨
func (s *sCouponMember) Create(ctx context.Context, in *model.CouponMemberCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayCouponMember.Ctx(ctx).Data(g.Map{
		dao.PlayCouponMember.Columns().Id:        id,
		dao.PlayCouponMember.Columns().CouponId: in.CouponID,
		dao.PlayCouponMember.Columns().MemberId: in.MemberID,
		dao.PlayCouponMember.Columns().OrderId: in.OrderID,
		dao.PlayCouponMember.Columns().UseStatus: in.UseStatus,
		dao.PlayCouponMember.Columns().ClaimAt: in.ClaimAt,
		dao.PlayCouponMember.Columns().UseAt: in.UseAt,
		dao.PlayCouponMember.Columns().ExpireAt: in.ExpireAt,
		dao.PlayCouponMember.Columns().CreatedAt: gtime.Now(),
		dao.PlayCouponMember.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨
func (s *sCouponMember) Update(ctx context.Context, in *model.CouponMemberUpdateInput) error {
	data := g.Map{
		dao.PlayCouponMember.Columns().CouponId: in.CouponID,
		dao.PlayCouponMember.Columns().MemberId: in.MemberID,
		dao.PlayCouponMember.Columns().OrderId: in.OrderID,
		dao.PlayCouponMember.Columns().UseStatus: in.UseStatus,
		dao.PlayCouponMember.Columns().ClaimAt: in.ClaimAt,
		dao.PlayCouponMember.Columns().UseAt: in.UseAt,
		dao.PlayCouponMember.Columns().ExpireAt: in.ExpireAt,
		dao.PlayCouponMember.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayCouponMember.Ctx(ctx).Where(dao.PlayCouponMember.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨
func (s *sCouponMember) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayCouponMember.Ctx(ctx).Where(dao.PlayCouponMember.Columns().Id, id).Data(g.Map{
		dao.PlayCouponMember.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨详情
func (s *sCouponMember) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.CouponMemberDetailOutput, err error) {
	out = &model.CouponMemberDetailOutput{}
	err = dao.PlayCouponMember.Ctx(ctx).Where(dao.PlayCouponMember.Columns().Id, id).Where(dao.PlayCouponMember.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	// 查询ä¼˜æƒ åˆ¸æ¨¡æ¿ID关联显示
	if out.CouponID != 0 {
		val, _ := g.DB().Ctx(ctx).Model("play_coupon").Where("id", out.CouponID).Where("deleted_at", nil).Value("title")
		out.CouponTitle = val.String()
	}
	return
}

// List 获取ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨列表
func (s *sCouponMember) List(ctx context.Context, in *model.CouponMemberListInput) (list []*model.CouponMemberListOutput, total int, err error) {
	m := dao.PlayCouponMember.Ctx(ctx).Where(dao.PlayCouponMember.Columns().DeletedAt, nil)
	if in.UseStatus > 0 {
		m = m.Where(dao.PlayCouponMember.Columns().UseStatus, in.UseStatus)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayCouponMember.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	// 填充关联显示字段
	for _, item := range list {
		if item.CouponID != 0 {
			val, _ := g.DB().Ctx(ctx).Model("play_coupon").Where("id", item.CouponID).Where("deleted_at", nil).Value("title")
			item.CouponTitle = val.String()
		}
	}
	return
}

