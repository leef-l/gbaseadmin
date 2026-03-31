package playapi

import (
	"context"
	"strconv"

	v1 "gbaseadmin/app/play/api/playapi/v1"
	"gbaseadmin/app/play/internal/dao"
)

type sSearch struct{}

// Search 综合搜索
func (s *sSearch) Search(ctx context.Context, keyword, searchType string, page, pageSize int) (
	coachList []v1.SearchCoachItem, coachTotal int,
	goodsList []v1.SearchGoodsItem, goodsTotal int,
	err error,
) {
	like := "%" + keyword + "%"

	// 搜索陪玩师
	if searchType == "all" || searchType == "coach" {
		mc := dao.PlayCoach.Ctx(ctx).
			Where(dao.PlayCoach.Columns().Status, 1).
			Where(dao.PlayCoach.Columns().DeletedAt, nil)
		// 关联 play_member 获取 nickname
		mc = mc.InnerJoin("play_member", "play_member.id=play_coach.member_id")
		mc = mc.Where("play_member.nickname LIKE ? OR play_coach.intro LIKE ?", like, like)

		coachTotal, err = mc.Count()
		if err != nil {
			return
		}
		limit := pageSize
		if searchType == "all" {
			limit = 5
		}
		var records []struct {
			Id          uint64  `json:"id"`
			Nickname    string  `json:"nickname"`
			Avatar      string  `json:"avatar"`
			IsOnline    int     `json:"is_online"`
			TotalScore  int     `json:"total_score"`
			ScoreNum    int     `json:"score_num"`
			TotalOrders int     `json:"total_orders"`
		}
		err = mc.Fields("play_coach.id, play_member.nickname, play_member.avatar, play_coach.is_online, play_coach.total_score, play_coach.score_num, play_coach.total_orders").
			OrderDesc("play_coach.total_score").
			Limit(limit).
			Scan(&records)
		if err != nil {
			return
		}
		coachList = make([]v1.SearchCoachItem, 0, len(records))
		for _, r := range records {
			score := float64(0)
			if r.ScoreNum > 0 {
				score = float64(r.TotalScore) / float64(r.ScoreNum) / 100
			}
			coachList = append(coachList, v1.SearchCoachItem{
				CoachID:    strconv.FormatUint(r.Id, 10),
				Nickname:   r.Nickname,
				Avatar:     r.Avatar,
				IsOnline:   r.IsOnline,
				Score:      score,
				OrderCount: r.TotalOrders,
			})
		}
	}

	// 搜索商品
	if searchType == "all" || searchType == "goods" {
		mg := dao.PlayGoods.Ctx(ctx).
			Where(dao.PlayGoods.Columns().Status, 1).
			Where(dao.PlayGoods.Columns().DeletedAt, nil).
			Where("title LIKE ? OR desc_content LIKE ?", like, like)
		// 确保陪玩师有效
		mg = mg.InnerJoin("play_coach", "play_coach.id=play_goods.coach_id AND play_coach.status=1")
		mg = mg.InnerJoin("play_member", "play_member.id=play_coach.member_id")

		goodsTotal, err = mg.Count()
		if err != nil {
			return
		}
		limit := pageSize
		if searchType == "all" {
			limit = 10
		}
		var gRecords []struct {
			Id         uint64 `json:"id"`
			Title      string `json:"title"`
			CoverImage string `json:"cover_image"`
			Price      int64  `json:"price"`
			Unit       string `json:"unit"`
			CoachId    uint64 `json:"coach_id"`
			Nickname   string `json:"nickname"`
		}
		err = mg.Fields("play_goods.id, play_goods.title, play_goods.cover_image, play_goods.price, play_goods.unit, play_goods.coach_id, play_member.nickname").
			OrderDesc("play_goods.sales_num").
			Limit(limit).
			Scan(&gRecords)
		if err != nil {
			return
		}
		goodsList = make([]v1.SearchGoodsItem, 0, len(gRecords))
		for _, r := range gRecords {
			goodsList = append(goodsList, v1.SearchGoodsItem{
				GoodsID:    strconv.FormatUint(r.Id, 10),
				Title:      r.Title,
				CoverImage: r.CoverImage,
				Price:      r.Price,
				Unit:       r.Unit,
				CoachID:    strconv.FormatUint(r.CoachId, 10),
				CoachName:  r.Nickname,
			})
		}
	}

	if coachList == nil {
		coachList = []v1.SearchCoachItem{}
	}
	if goodsList == nil {
		goodsList = []v1.SearchGoodsItem{}
	}
	return
}
