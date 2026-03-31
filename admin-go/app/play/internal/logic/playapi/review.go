package playapi

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	v1 "gbaseadmin/app/play/api/playapi/v1"
	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/service"
	"gbaseadmin/utility/snowflake"
)

type sPlayapiReview struct{}

func init() {
	service.RegisterPlayapiReview(&sPlayapiReview{})
}

func (s *sPlayapiReview) Create(ctx context.Context, req *v1.ReviewCreateReq) error {
	r := g.RequestFromCtx(ctx)
	memberID := r.GetCtxVar("jwt_member_id").Int64()

	oc := dao.PlayOrder.Columns()
	order, err := dao.PlayOrder.Ctx(ctx).Where(oc.Id, req.OrderID).Where(oc.MemberId, memberID).One()
	if err != nil {
		return err
	}
	if order.IsEmpty() {
		return gerror.New("订单不存在")
	}
	if order[oc.OrderStatus].Int() != 3 {
		return gerror.New("订单未完成，不可评价")
	}

	// 检查是否已评价
	rc := dao.PlayReview.Columns()
	cnt, err := dao.PlayReview.Ctx(ctx).Where(rc.OrderId, req.OrderID).Where(rc.MemberId, memberID).Count()
	if err != nil {
		return err
	}
	if cnt > 0 {
		return gerror.New("该订单已评价")
	}

	coachID := order[oc.CoachId].Int64()
	scoreInt := int(req.Score * 100)

	// 插入评价
	_, err = dao.PlayReview.Ctx(ctx).Data(g.Map{
		rc.Id:            snowflake.Generate(),
		rc.OrderId:       req.OrderID,
		rc.MemberId:      memberID,
		rc.CoachId:       coachID,
		rc.Score:         scoreInt,
		rc.ReviewContent: req.Content,
		rc.ReviewImage:   req.Images,
		rc.IsAnonymous:   req.IsAnonymous,
		rc.Status:        1,
		rc.CreatedAt:     gtime.Now(),
		rc.UpdatedAt:     gtime.Now(),
	}).Insert()
	if err != nil {
		return err
	}

	// 更新陪玩师评分统计
	cc := dao.PlayCoach.Columns()
	_, err = dao.PlayCoach.Ctx(ctx).Where(cc.Id, coachID).Data(g.Map{
		cc.TotalScore: gdb.Raw(fmt.Sprintf("%s + %d", cc.TotalScore, scoreInt)),
		cc.ScoreNum:   gdb.Raw(cc.ScoreNum + " + 1"),
		cc.UpdatedAt:  gtime.Now(),
	}).Update()
	return err
}

func (s *sPlayapiReview) List(ctx context.Context, req *v1.ReviewListReq) (list []v1.ReviewListItem, total int, err error) {
	rc := dao.PlayReview.Columns()
	mc := dao.PlayMember.Columns()

	m := dao.PlayReview.Ctx(ctx).As("r").
		LeftJoin(dao.PlayMember.Table()+" m", fmt.Sprintf("m.%s = r.%s", mc.Id, rc.MemberId)).
		Where("r."+rc.CoachId, req.CoachID).Where("r."+rc.Status, 1)

	total, err = m.Count()
	if err != nil {
		return
	}

	records, err := m.Fields("r.*, m."+mc.Nickname+", m."+mc.Avatar).
		OrderDesc("r." + rc.CreatedAt).Page(req.Page, req.PageSize).All()
	if err != nil {
		return
	}

	list = make([]v1.ReviewListItem, 0, len(records))
	for _, rr := range records {
		item := v1.ReviewListItem{
			ReviewID:     rr[rc.Id].String(),
			Score:        rr[rc.Score].Float64() / 100,
			Content:      rr[rc.ReviewContent].String(),
			Images:       rr[rc.ReviewImage].String(),
			IsAnonymous:  rr[rc.IsAnonymous].Int(),
			ReplyContent: rr[rc.ReplyContent].String(),
			CreatedAt:    rr[rc.CreatedAt].String(),
		}
		if item.IsAnonymous == 0 {
			item.Nickname = rr[mc.Nickname].String()
			item.Avatar = rr[mc.Avatar].String()
		} else {
			item.Nickname = "匿名用户"
		}
		list = append(list, item)
	}
	return
}

func (s *sPlayapiReview) Reply(ctx context.Context, req *v1.ReviewReplyReq) error {
	r := g.RequestFromCtx(ctx)
	coachID := r.GetCtxVar("jwt_coach_id").Int64()
	rc := dao.PlayReview.Columns()

	review, err := dao.PlayReview.Ctx(ctx).Where(rc.Id, req.ReviewID).Where(rc.CoachId, coachID).One()
	if err != nil {
		return err
	}
	if review.IsEmpty() {
		return gerror.New("评价不存在")
	}
	if review[rc.ReplyContent].String() != "" {
		return gerror.New("已回复过该评价")
	}

	_, err = dao.PlayReview.Ctx(ctx).Where(rc.Id, req.ReviewID).Data(g.Map{
		rc.ReplyContent: req.Reply,
		rc.ReplyAt:      gtime.Now(),
		rc.UpdatedAt:    gtime.Now(),
	}).Update()
	return err
}
