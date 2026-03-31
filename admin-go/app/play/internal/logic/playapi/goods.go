package playapi

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "gbaseadmin/app/play/api/playapi/v1"
	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/service"
)

type sPlayapiGoods struct{}

func init() {
	service.RegisterPlayapiGoods(&sPlayapiGoods{})
}

func (s *sPlayapiGoods) List(ctx context.Context, req *v1.GoodsListReq) (list []v1.GoodsListItem, total int, err error) {
	gc := dao.PlayGoods.Columns()
	cc := dao.PlayCoach.Columns()
	mc := dao.PlayMember.Columns()

	m := dao.PlayGoods.Ctx(ctx).As("g").
		LeftJoin(dao.PlayCoach.Table()+" c", fmt.Sprintf("c.%s = g.%s", cc.Id, gc.CoachId)).
		LeftJoin(dao.PlayMember.Table()+" m", fmt.Sprintf("m.%s = c.%s", mc.Id, cc.MemberId)).
		Where("g."+gc.Status, 1).
		Where("c."+cc.Status, 1)

	if req.CategoryID != "" {
		m = m.Where("g."+gc.CategoryId, req.CategoryID)
	}
	if req.Keyword != "" {
		m = m.Where("g."+gc.Title+" LIKE ? OR g."+gc.DescContent+" LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}

	switch req.SortBy {
	case "price_asc":
		m = m.OrderAsc("g." + gc.Price)
	case "price_desc":
		m = m.OrderDesc("g." + gc.Price)
	case "sales":
		m = m.OrderDesc("g." + gc.SalesNum)
	case "newest":
		m = m.OrderDesc("g." + gc.CreatedAt)
	default:
		m = m.OrderAsc("g." + gc.Sort).OrderDesc("g." + gc.SalesNum)
	}

	total, err = m.Count()
	if err != nil {
		return
	}

	records, err := m.Fields(
		"g.*",
		"c."+cc.IsOnline+" as coach_online",
		"c."+cc.TotalScore+" as coach_total_score",
		"c."+cc.ScoreNum+" as coach_score_num",
		"m."+mc.Nickname+" as coach_name",
		"m."+mc.Avatar+" as coach_avatar",
	).Page(req.Page, req.PageSize).All()
	if err != nil {
		return
	}

	list = make([]v1.GoodsListItem, 0, len(records))
	for _, r := range records {
		score := float64(0)
		sn := r["coach_score_num"].Int()
		if sn > 0 {
			score = r["coach_total_score"].Float64() / float64(sn) / 100
		}
		list = append(list, v1.GoodsListItem{
			GoodsID:     r[gc.Id].String(),
			Title:       r[gc.Title].String(),
			CoverImage:  r[gc.CoverImage].String(),
			Price:       r[gc.Price].Int64(),
			Unit:        r[gc.Unit].String(),
			SalesNum:    r[gc.SalesNum].Int(),
			CoachID:     r[gc.CoachId].String(),
			CoachName:   r["coach_name"].String(),
			CoachAvatar: r["coach_avatar"].String(),
			CoachScore:  score,
			CoachOnline: r["coach_online"].Int(),
		})
	}
	return
}

func (s *sPlayapiGoods) Detail(ctx context.Context, req *v1.GoodsDetailReq) (res *v1.GoodsDetailRes, err error) {
	gc := dao.PlayGoods.Columns()
	goods, err := dao.PlayGoods.Ctx(ctx).Where(gc.Id, req.GoodsID).Where(gc.Status, 1).One()
	if err != nil {
		return
	}
	if goods.IsEmpty() {
		err = gerror.New("商品不存在或已下架")
		return
	}

	cc := dao.PlayCoach.Columns()
	mc := dao.PlayMember.Columns()
	coach, _ := dao.PlayCoach.Ctx(ctx).Where(cc.Id, goods[gc.CoachId]).One()
	member, _ := dao.PlayMember.Ctx(ctx).Where(mc.Id, coach[cc.MemberId]).One()

	catName := ""
	catVal, _ := dao.PlayCategory.Ctx(ctx).Where(dao.PlayCategory.Columns().Id, goods[gc.CategoryId]).Value(dao.PlayCategory.Columns().Title)
	if catVal != nil {
		catName = catVal.String()
	}

	score := float64(0)
	if !coach.IsEmpty() && coach[cc.ScoreNum].Int() > 0 {
		score = coach[cc.TotalScore].Float64() / float64(coach[cc.ScoreNum].Int()) / 100
	}

	res = &v1.GoodsDetailRes{
		GoodsID:      req.GoodsID,
		Title:        goods[gc.Title].String(),
		Description:  goods[gc.DescContent].String(),
		CoverImage:   goods[gc.CoverImage].String(),
		CategoryID:   goods[gc.CategoryId].String(),
		CategoryName: catName,
		Price:        goods[gc.Price].Int64(),
		Unit:         goods[gc.Unit].String(),
		SalesNum:     goods[gc.SalesNum].Int(),
		CoachID:      goods[gc.CoachId].String(),
	}
	if !member.IsEmpty() {
		res.CoachName = member[mc.Nickname].String()
		res.CoachAvatar = member[mc.Avatar].String()
	}
	if !coach.IsEmpty() {
		res.CoachScore = score
		res.CoachOnline = coach[cc.IsOnline].Int()
	}
	return
}

func (s *sPlayapiGoods) CategoryList(ctx context.Context, req *v1.CategoryListReq) (list []v1.CategoryTreeItem, err error) {
	catC := dao.PlayCategory.Columns()
	records, err := dao.PlayCategory.Ctx(ctx).Where(catC.Status, 1).OrderAsc(catC.Sort).All()
	if err != nil {
		return
	}

	// 构建树形结构
	type catNode struct {
		item     v1.CategoryTreeItem
		parentID string
	}
	nodes := make([]catNode, 0, len(records))
	for _, r := range records {
		nodes = append(nodes, catNode{
			item: v1.CategoryTreeItem{
				CategoryID: r[catC.Id].String(),
				Name:       r[catC.Title].String(),
				Icon:       r[catC.Icon].String(),
				CoverImage: r[catC.CoverImage].String(),
				Sort:       r[catC.Sort].Int(),
				Children:   []v1.CategoryTreeItem{},
			},
			parentID: r[catC.ParentId].String(),
		})
	}

	childMap := make(map[string][]int)
	for i, n := range nodes {
		childMap[n.parentID] = append(childMap[n.parentID], i)
	}

	var buildTree func(parentID string) []v1.CategoryTreeItem
	buildTree = func(parentID string) []v1.CategoryTreeItem {
		indices := childMap[parentID]
		result := make([]v1.CategoryTreeItem, 0, len(indices))
		for _, idx := range indices {
			item := nodes[idx].item
			item.Children = buildTree(nodes[idx].item.CategoryID)
			result = append(result, item)
		}
		return result
	}

	list = buildTree("0")
	return
}
