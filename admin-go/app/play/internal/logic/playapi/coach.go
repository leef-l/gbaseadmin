package playapi

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	v1 "gbaseadmin/api/playapi/v1"
	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/service"
	"gbaseadmin/utility/snowflake"
)

type sPlayapiCoach struct{}

func init() {
	service.RegisterPlayapiCoach(&sPlayapiCoach{})
}

func (s *sPlayapiCoach) List(ctx context.Context, req *v1.CoachListReq) (list []v1.CoachListItem, total int, err error) {
	cc := dao.PlayCoach.Columns()
	mc := dao.PlayMember.Columns()
	m := dao.PlayCoach.Ctx(ctx).As("c").
		LeftJoin(dao.PlayMember.Table()+" m", fmt.Sprintf("m.%s = c.%s", mc.Id, cc.MemberId)).
		Where("c."+cc.Status, 1)

	if req.CategoryID != "" {
		gc := dao.PlayGoods.Columns()
		m = m.WhereIn("c."+cc.Id,
			dao.PlayGoods.Ctx(ctx).Fields(gc.CoachId).Where(gc.CategoryId, req.CategoryID).Where(gc.Status, 1),
		)
	}
	if req.Keyword != "" {
		m = m.Where("m."+mc.Nickname+" LIKE ? OR c."+cc.Intro+" LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}
	if req.Gender != nil {
		m = m.Where("m."+mc.Gender, *req.Gender)
	}
	if req.OnlineOnly != nil && *req.OnlineOnly == 1 {
		m = m.Where("c."+cc.IsOnline, 1)
	}

	switch req.SortBy {
	case "score":
		m = m.OrderDesc("c." + cc.TotalScore)
	case "orderCount":
		m = m.OrderDesc("c." + cc.TotalOrders)
	case "price_asc":
		// min price from goods subquery
		m = m.OrderAsc("(SELECT MIN(price) FROM " + dao.PlayGoods.Table() + " WHERE coach_id=c.id AND status=1)")
	case "price_desc":
		m = m.OrderDesc("(SELECT MIN(price) FROM " + dao.PlayGoods.Table() + " WHERE coach_id=c.id AND status=1)")
	default:
		m = m.OrderDesc("c." + cc.IsOnline).OrderDesc("c." + cc.TotalScore)
	}

	total, err = m.Count()
	if err != nil {
		return
	}

	records, err := m.Fields(
		"c."+cc.Id+" as coach_id",
		"m."+mc.Nickname,
		"m."+mc.Avatar,
		"m."+mc.Gender,
		"c."+cc.IsOnline,
		"c."+cc.TotalScore,
		"c."+cc.ScoreNum,
		"c."+cc.TotalOrders,
		"c."+cc.Intro,
	).Page(req.Page, req.PageSize).All()
	if err != nil {
		return
	}

	list = make([]v1.CoachListItem, 0, len(records))
	for _, r := range records {
		score := float64(0)
		sn := r["score_num"].Int()
		ts := r["total_score"].Float64()
		if sn > 0 {
			score = ts / float64(sn) / 100
		}
		item := v1.CoachListItem{
			CoachID:    r["coach_id"].String(),
			Nickname:   r["nickname"].String(),
			Avatar:     r["avatar"].String(),
			Gender:     r["gender"].Int(),
			IsOnline:   r["is_online"].Int(),
			Score:      score,
			OrderCount: r["total_orders"].Int(),
			Intro:      r["intro"].String(),
		}
		// 查最低价
		minPrice, _ := dao.PlayGoods.Ctx(ctx).
			Where(dao.PlayGoods.Columns().CoachId, r["coach_id"]).
			Where(dao.PlayGoods.Columns().Status, 1).
			Min(dao.PlayGoods.Columns().Price)
		item.MinPrice = int64(minPrice)
		list = append(list, item)
	}
	return
}

func (s *sPlayapiCoach) Detail(ctx context.Context, req *v1.CoachDetailReq) (res *v1.CoachDetailRes, err error) {
	cc := dao.PlayCoach.Columns()
	mc := dao.PlayMember.Columns()

	coach, err := dao.PlayCoach.Ctx(ctx).Where(cc.Id, req.CoachID).Where(cc.Status, 1).One()
	if err != nil {
		return
	}
	if coach.IsEmpty() {
		err = gerror.New("陪玩师不存在")
		return
	}

	member, err := dao.PlayMember.Ctx(ctx).Where(mc.Id, coach[cc.MemberId]).One()
	if err != nil {
		return
	}

	scoreNum := coach[cc.ScoreNum].Int()
	score := float64(0)
	if scoreNum > 0 {
		score = coach[cc.TotalScore].Float64() / float64(scoreNum) / 100
	}

	res = &v1.CoachDetailRes{
		CoachID:    req.CoachID,
		Nickname:   member[mc.Nickname].String(),
		Avatar:     member[mc.Avatar].String(),
		Gender:     member[mc.Gender].Int(),
		IsOnline:   coach[cc.IsOnline].Int(),
		Score:      score,
		OrderCount: coach[cc.TotalOrders].Int(),
		Intro:      coach[cc.Intro].String(),
	}

	// 商品列表
	gc := dao.PlayGoods.Columns()
	goodsRecords, err := dao.PlayGoods.Ctx(ctx).
		Where(gc.CoachId, req.CoachID).Where(gc.Status, 1).
		OrderAsc(gc.Sort).All()
	if err != nil {
		return
	}
	res.GoodsList = make([]v1.CoachGoodsItem, 0, len(goodsRecords))
	for _, gr := range goodsRecords {
		catName := ""
		catVal, _ := dao.PlayCategory.Ctx(ctx).Where(dao.PlayCategory.Columns().Id, gr[gc.CategoryId]).Value(dao.PlayCategory.Columns().Title)
		if catVal != nil {
			catName = catVal.String()
		}
		res.GoodsList = append(res.GoodsList, v1.CoachGoodsItem{
			GoodsID:      gr[gc.Id].String(),
			Title:        gr[gc.Title].String(),
			CategoryName: catName,
			Price:        gr[gc.Price].Int64(),
			Unit:         gr[gc.Unit].String(),
			Description:  gr[gc.DescContent].String(),
			Status:       gr[gc.Status].Int(),
		})
	}

	// 评价
	rc := dao.PlayReview.Columns()
	res.ReviewCount, err = dao.PlayReview.Ctx(ctx).Where(rc.CoachId, req.CoachID).Where(rc.Status, 1).Count()
	if err != nil {
		return
	}
	reviewRecords, err := dao.PlayReview.Ctx(ctx).As("r").
		LeftJoin(dao.PlayMember.Table()+" m", fmt.Sprintf("m.%s = r.%s", mc.Id, rc.MemberId)).
		Where("r."+rc.CoachId, req.CoachID).Where("r."+rc.Status, 1).
		OrderDesc("r." + rc.CreatedAt).Limit(3).
		Fields("r.*, m."+mc.Nickname+", m."+mc.Avatar).All()
	if err != nil {
		return
	}
	res.RecentReviews = make([]v1.ReviewBriefItem, 0, len(reviewRecords))
	for _, rr := range reviewRecords {
		item := v1.ReviewBriefItem{
			Score:       rr[rc.Score].Float64() / 100,
			Content:     rr[rc.ReviewContent].String(),
			IsAnonymous: rr[rc.IsAnonymous].Int(),
			CreatedAt:   rr[rc.CreatedAt].String(),
		}
		if item.IsAnonymous == 0 {
			item.Nickname = rr[mc.Nickname].String()
			item.Avatar = rr[mc.Avatar].String()
		} else {
			item.Nickname = "匿名用户"
		}
		res.RecentReviews = append(res.RecentReviews, item)
	}
	return
}

func (s *sPlayapiCoach) Apply(ctx context.Context, req *v1.CoachApplyReq) error {
	r := g.RequestFromCtx(ctx)
	memberID := r.GetCtxVar("jwt_member_id").Int64()

	// 检查是否已是陪玩师
	mc := dao.PlayMember.Columns()
	isCoach, err := dao.PlayMember.Ctx(ctx).Where(mc.Id, memberID).Value(mc.IsCoach)
	if err != nil {
		return err
	}
	if isCoach.Int() == 1 {
		return gerror.New("您已经是陪玩师")
	}

	// 检查是否有待审核的申请
	ac := dao.PlayCoachApply.Columns()
	count, err := dao.PlayCoachApply.Ctx(ctx).
		Where(ac.MemberId, memberID).Where(ac.AuditStatus, 0).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("您已有待审核的申请")
	}

	_, err = dao.PlayCoachApply.Ctx(ctx).Data(g.Map{
		ac.Id:               snowflake.Generate(),
		ac.MemberId:         memberID,
		ac.RealName:         req.RealName,
		ac.IdCard:           req.IdCard,
		ac.IdCardFrontImage: req.IdCardFrontImage,
		ac.IdCardBackImage:  req.IdCardBackImage,
		ac.SkillDesc:        req.SkillDesc,
		ac.AuditStatus:      0,
		ac.CreatedAt:        gtime.Now(),
		ac.UpdatedAt:        gtime.Now(),
	}).Insert()
	return err
}

func (s *sPlayapiCoach) ApplyStatus(ctx context.Context, req *v1.CoachApplyStatusReq) (res *v1.CoachApplyStatusRes, err error) {
	r := g.RequestFromCtx(ctx)
	memberID := r.GetCtxVar("jwt_member_id").Int64()

	ac := dao.PlayCoachApply.Columns()
	record, err := dao.PlayCoachApply.Ctx(ctx).
		Where(ac.MemberId, memberID).OrderDesc(ac.CreatedAt).One()
	if err != nil {
		return
	}

	res = &v1.CoachApplyStatusRes{}
	if record.IsEmpty() {
		return
	}
	res.HasApply = true
	res.AuditStatus = record[ac.AuditStatus].Int()
	res.AuditRemark = record[ac.AuditRemark].String()
	return
}

func (s *sPlayapiCoach) SetOnline(ctx context.Context, req *v1.CoachOnlineReq) error {
	r := g.RequestFromCtx(ctx)
	coachID := r.GetCtxVar("jwt_coach_id").Int64()
	cc := dao.PlayCoach.Columns()
	_, err := dao.PlayCoach.Ctx(ctx).Where(cc.Id, coachID).Data(g.Map{
		cc.IsOnline:  req.IsOnline,
		cc.UpdatedAt: gtime.Now(),
	}).Update()
	return err
}

func (s *sPlayapiCoach) MyGoods(ctx context.Context, req *v1.CoachMyGoodsReq) (list []v1.CoachGoodsItem, total int, err error) {
	r := g.RequestFromCtx(ctx)
	coachID := r.GetCtxVar("jwt_coach_id").Int64()
	gc := dao.PlayGoods.Columns()
	m := dao.PlayGoods.Ctx(ctx).Where(gc.CoachId, coachID)
	if req.Status != nil {
		m = m.Where(gc.Status, *req.Status)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	records, err := m.OrderAsc(gc.Sort).OrderDesc(gc.CreatedAt).Page(req.Page, req.PageSize).All()
	if err != nil {
		return
	}
	list = make([]v1.CoachGoodsItem, 0, len(records))
	for _, gr := range records {
		catName := ""
		catVal, _ := dao.PlayCategory.Ctx(ctx).Where(dao.PlayCategory.Columns().Id, gr[gc.CategoryId]).Value(dao.PlayCategory.Columns().Title)
		if catVal != nil {
			catName = catVal.String()
		}
		list = append(list, v1.CoachGoodsItem{
			GoodsID:      gr[gc.Id].String(),
			Title:        gr[gc.Title].String(),
			CategoryName: catName,
			Price:        gr[gc.Price].Int64(),
			Unit:         gr[gc.Unit].String(),
			Description:  gr[gc.DescContent].String(),
			Status:       gr[gc.Status].Int(),
		})
	}
	return
}

func (s *sPlayapiCoach) GoodsCreate(ctx context.Context, req *v1.CoachGoodsCreateReq) error {
	r := g.RequestFromCtx(ctx)
	coachID := r.GetCtxVar("jwt_coach_id").Int64()

	// 校验分类
	catC := dao.PlayCategory.Columns()
	catCount, err := dao.PlayCategory.Ctx(ctx).Where(catC.Id, req.CategoryID).Where(catC.Status, 1).Count()
	if err != nil {
		return err
	}
	if catCount == 0 {
		return gerror.New("分类不存在或已禁用")
	}

	gc := dao.PlayGoods.Columns()
	_, err = dao.PlayGoods.Ctx(ctx).Data(g.Map{
		gc.Id:          snowflake.Generate(),
		gc.CoachId:     coachID,
		gc.CategoryId:  req.CategoryID,
		gc.Title:       req.Title,
		gc.DescContent: req.Description,
		gc.Price:       req.Price,
		gc.Unit:        req.Unit,
		gc.CoverImage:  req.CoverImage,
		gc.Sort:        req.Sort,
		gc.Status:      1,
		gc.CreatedAt:   gtime.Now(),
		gc.UpdatedAt:   gtime.Now(),
	}).Insert()
	return err
}

func (s *sPlayapiCoach) GoodsUpdate(ctx context.Context, req *v1.CoachGoodsUpdateReq) error {
	r := g.RequestFromCtx(ctx)
	coachID := r.GetCtxVar("jwt_coach_id").Int64()
	gc := dao.PlayGoods.Columns()

	count, err := dao.PlayGoods.Ctx(ctx).Where(gc.Id, req.GoodsID).Where(gc.CoachId, coachID).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("商品不存在")
	}

	if req.CategoryID != "" {
		catC := dao.PlayCategory.Columns()
		catCount, e := dao.PlayCategory.Ctx(ctx).Where(catC.Id, req.CategoryID).Where(catC.Status, 1).Count()
		if e != nil {
			return e
		}
		if catCount == 0 {
			return gerror.New("分类不存在或已禁用")
		}
	}

	data := g.Map{gc.UpdatedAt: gtime.Now()}
	if req.CategoryID != "" {
		data[gc.CategoryId] = req.CategoryID
	}
	if req.Title != "" {
		data[gc.Title] = req.Title
	}
	if req.Description != "" {
		data[gc.DescContent] = req.Description
	}
	if req.Price != nil {
		data[gc.Price] = *req.Price
	}
	if req.Unit != "" {
		data[gc.Unit] = req.Unit
	}
	if req.CoverImage != "" {
		data[gc.CoverImage] = req.CoverImage
	}
	if req.Sort != nil {
		data[gc.Sort] = *req.Sort
	}

	_, err = dao.PlayGoods.Ctx(ctx).Where(gc.Id, req.GoodsID).Data(data).Update()
	return err
}

func (s *sPlayapiCoach) GoodsStatus(ctx context.Context, req *v1.CoachGoodsStatusReq) error {
	r := g.RequestFromCtx(ctx)
	coachID := r.GetCtxVar("jwt_coach_id").Int64()
	gc := dao.PlayGoods.Columns()

	count, err := dao.PlayGoods.Ctx(ctx).Where(gc.Id, req.GoodsID).Where(gc.CoachId, coachID).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("商品不存在")
	}

	_, err = dao.PlayGoods.Ctx(ctx).Where(gc.Id, req.GoodsID).Data(g.Map{
		gc.Status:    req.Status,
		gc.UpdatedAt: gtime.Now(),
	}).Update()
	return err
}

func (s *sPlayapiCoach) Income(ctx context.Context, req *v1.CoachIncomeReq) (res *v1.CoachIncomeRes, err error) {
	r := g.RequestFromCtx(ctx)
	coachID := r.GetCtxVar("jwt_coach_id").Int64()
	oc := dao.PlayOrder.Columns()

	res = &v1.CoachIncomeRes{}

	// 已完成订单基础查询
	base := dao.PlayOrder.Ctx(ctx).Where(oc.CoachId, coachID).Where(oc.OrderStatus, 3)

	// 今日收入
	todayIncome, _ := base.Clone().WhereGTE(oc.FinishAt, gtime.Now().Format("Y-m-d")).Sum(oc.PayAmount)
	res.TodayIncome = int64(todayIncome)

	// 本月收入
	monthStart := gtime.Now().Format("Y-m") + "-01"
	monthIncome, _ := base.Clone().WhereGTE(oc.FinishAt, monthStart).Sum(oc.PayAmount)
	res.MonthIncome = int64(monthIncome)

	// 累计总收入
	totalIncome, _ := base.Clone().Sum(oc.PayAmount)
	res.TotalIncome = int64(totalIncome)

	// 接单统计
	acceptBase := dao.PlayOrder.Ctx(ctx).Where(oc.CoachId, coachID).WhereIn(oc.OrderStatus, g.Slice{2, 3})
	res.TodayOrders, _ = acceptBase.Clone().WhereGTE(oc.CreatedAt, gtime.Now().Format("Y-m-d")).Count()
	res.TotalOrders, _ = acceptBase.Clone().Count()

	return
}

func (s *sPlayapiCoach) Orders(ctx context.Context, req *v1.CoachOrdersReq) (list []v1.OrderListItem, total int, err error) {
	r := g.RequestFromCtx(ctx)
	coachID := r.GetCtxVar("jwt_coach_id").Int64()
	oc := dao.PlayOrder.Columns()
	mc := dao.PlayMember.Columns()
	gc := dao.PlayGoods.Columns()

	m := dao.PlayOrder.Ctx(ctx).As("o").Where("o."+oc.CoachId, coachID)
	if req.Status != nil {
		m = m.Where("o."+oc.OrderStatus, *req.Status)
	}

	total, err = m.Count()
	if err != nil {
		return
	}

	records, err := m.OrderDesc("o." + oc.CreatedAt).Page(req.Page, req.PageSize).All()
	if err != nil {
		return
	}

	list = make([]v1.OrderListItem, 0, len(records))
	for _, o := range records {
		memberInfo, _ := dao.PlayMember.Ctx(ctx).Where(mc.Id, o[oc.MemberId]).One()
		goodsImg := ""
		if gid := o[oc.GoodsId]; gid != nil {
			imgVal, _ := dao.PlayGoods.Ctx(ctx).Where(gc.Id, gid).Value(gc.CoverImage)
			if imgVal != nil {
				goodsImg = imgVal.String()
			}
		}
		coachName := ""
		coachAvatar := ""
		if !memberInfo.IsEmpty() {
			coachName = memberInfo[mc.Nickname].String()
			coachAvatar = memberInfo[mc.Avatar].String()
		}
		list = append(list, v1.OrderListItem{
			OrderID:     o[oc.Id].String(),
			OrderNo:     o[oc.OrderNo].String(),
			GoodsTitle:  o[oc.GoodsTitle].String(),
			GoodsImage:  goodsImg,
			CoachID:     strconv.FormatInt(coachID, 10),
			CoachName:   coachName,
			CoachAvatar: coachAvatar,
			Quantity:    o[oc.Quantity].Int(),
			TotalAmount: o[oc.TotalAmount].Int64(),
			PayAmount:   o[oc.PayAmount].Int64(),
			OrderStatus: o[oc.OrderStatus].Int(),
			CreatedAt:   o[oc.CreatedAt].String(),
		})
	}
	return
}
