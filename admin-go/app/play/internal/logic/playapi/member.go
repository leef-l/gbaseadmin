package playapi

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	v1 "gbaseadmin/api/playapi/v1"
	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/model/entity"
	"gbaseadmin/utility/jwt"
)

type sMember struct{}

func (s *sMember) Info(ctx context.Context, req *v1.MemberInfoReq) (*v1.MemberInfoRes, error) {
	memberId := g.RequestFromCtx(ctx).GetCtxVar("jwt_member_id").Int64()
	claims, _ := g.RequestFromCtx(ctx).GetCtxVar("jwt_member_claims").Val().(*jwt.MemberClaims)

	memberColumns := dao.PlayMember.Columns()
	record, err := dao.PlayMember.Ctx(ctx).
		Where(memberColumns.Id, memberId).
		Where(memberColumns.DeletedAt, nil).
		One()
	if err != nil {
		return nil, err
	}
	if record.IsEmpty() {
		return nil, fmt.Errorf("会员不存在")
	}

	// 查询会员等级信息
	var levelTitle, levelIcon string
	var discount int
	levelId := record[memberColumns.MemberLevelId].Int64()
	if levelId > 0 {
		levelColumns := dao.PlayMemberLevel.Columns()
		levelRecord, _ := dao.PlayMemberLevel.Ctx(ctx).
			Where(levelColumns.Id, levelId).
			Where(levelColumns.DeletedAt, nil).
			One()
		if !levelRecord.IsEmpty() {
			levelTitle = levelRecord[levelColumns.Title].String()
			levelIcon = levelRecord[levelColumns.Icon].String()
			discount = levelRecord[levelColumns.Discount].Int()
		}
	}

	// 手机号脱敏
	phone := record[memberColumns.Phone].String()
	if len(phone) >= 11 {
		phone = phone[:3] + "****" + phone[7:]
	}

	// 查询陪玩师ID
	var coachIdStr string
	isCoach := record[memberColumns.IsCoach].Int()
	if isCoach == 1 {
		coachRecord, _ := dao.PlayCoach.Ctx(ctx).
			Where(dao.PlayCoach.Columns().MemberId, memberId).
			Where(dao.PlayCoach.Columns().DeletedAt, nil).
			One()
		if !coachRecord.IsEmpty() {
			coachIdStr = coachRecord[dao.PlayCoach.Columns().Id].String()
		}
	}

	currentRole := "member"
	if claims != nil {
		currentRole = claims.CurrentRole
	}

	return &v1.MemberInfoRes{
		MemberID:    strconv.FormatInt(memberId, 10),
		Phone:       phone,
		Nickname:    record[memberColumns.Nickname].String(),
		Avatar:      record[memberColumns.Avatar].String(),
		Gender:      record[memberColumns.Gender].Int(),
		Balance:     record[memberColumns.Balance].Int64(),
		LevelTitle:  levelTitle,
		LevelIcon:   levelIcon,
		Discount:    discount,
		Exp:         record[memberColumns.Exp].Int(),
		IsCoach:     isCoach,
		CoachID:     coachIdStr,
		CurrentRole: currentRole,
		WxBound:     false,
		AlipayBound: false,
	}, nil
}

func (s *sMember) Update(ctx context.Context, req *v1.MemberUpdateReq) (*v1.MemberUpdateRes, error) {
	memberId := g.RequestFromCtx(ctx).GetCtxVar("jwt_member_id").Int64()
	memberColumns := dao.PlayMember.Columns()

	data := g.Map{
		memberColumns.UpdatedAt: gtime.Now(),
	}
	if req.Nickname != "" {
		data[memberColumns.Nickname] = req.Nickname
	}
	if req.Avatar != "" {
		data[memberColumns.Avatar] = req.Avatar
	}
	if req.Gender != nil {
		data[memberColumns.Gender] = *req.Gender
	}

	_, err := dao.PlayMember.Ctx(ctx).
		Where(memberColumns.Id, memberId).
		Data(data).
		Update()
	if err != nil {
		return nil, err
	}

	// 如果是陪玩师，同步更新 play_coach
	isCoach := g.RequestFromCtx(ctx).GetCtxVar("jwt_is_coach").Int()
	if isCoach == 1 {
		coachData := g.Map{}
		if req.Nickname != "" {
			coachData["real_name"] = req.Nickname
		}
		if req.Avatar != "" {
			coachData["cover_image"] = req.Avatar
		}
		if len(coachData) > 0 {
			coachData[dao.PlayCoach.Columns().UpdatedAt] = gtime.Now()
			_, _ = dao.PlayCoach.Ctx(ctx).
				Where(dao.PlayCoach.Columns().MemberId, memberId).
				Data(coachData).
				Update()
		}
	}

	return &v1.MemberUpdateRes{}, nil
}

func (s *sMember) SwitchRole(ctx context.Context, req *v1.MemberSwitchRoleReq) (*v1.MemberSwitchRoleRes, error) {
	claims, _ := g.RequestFromCtx(ctx).GetCtxVar("jwt_member_claims").Val().(*jwt.MemberClaims)
	if claims == nil {
		return nil, fmt.Errorf("登录信息异常")
	}

	var coachId int64
	if req.Role == "coach" {
		if claims.IsCoach != 1 {
			return nil, fmt.Errorf("您还不是陪玩师")
		}
		coachRecord, _ := dao.PlayCoach.Ctx(ctx).
			Where(dao.PlayCoach.Columns().MemberId, claims.MemberID).
			Where(dao.PlayCoach.Columns().DeletedAt, nil).
			One()
		if !coachRecord.IsEmpty() {
			coachId = coachRecord[dao.PlayCoach.Columns().Id].Int64()
		}
	}

	token, err := jwt.GenerateMemberToken(claims.MemberID, claims.Phone, claims.IsCoach, coachId, req.Role)
	if err != nil {
		return nil, err
	}

	return &v1.MemberSwitchRoleRes{
		Token: token,
	}, nil
}

func (s *sMember) BalanceLog(ctx context.Context, req *v1.MemberBalanceLogReq) (*v1.MemberBalanceLogRes, error) {
	memberId := g.RequestFromCtx(ctx).GetCtxVar("jwt_member_id").Int64()
	columns := dao.PlayBalanceLog.Columns()

	m := dao.PlayBalanceLog.Ctx(ctx).
		Where(columns.MemberId, memberId).
		Where(columns.DeletedAt, nil)

	if req.Type == "income" {
		m = m.Where(fmt.Sprintf("%s>?", columns.ChangeAmount), 0)
	} else if req.Type == "expense" {
		m = m.Where(fmt.Sprintf("%s<?", columns.ChangeAmount), 0)
	}

	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	var entities []*entity.PlayBalanceLog
	err = m.Page(req.Page, req.PageSize).
		OrderDesc(columns.CreatedAt).
		Scan(&entities)
	if err != nil {
		return nil, err
	}

	// 业务类型映射
	bizTypeMap := map[int]string{
		1: "recharge",
		2: "pay",
		3: "refund",
		4: "income",
		5: "withdraw",
	}

	list := make([]v1.MemberBalanceLogItem, 0, len(entities))
	for _, e := range entities {
		typeName := bizTypeMap[e.BizType]
		if typeName == "" {
			typeName = "other"
		}
		createdAt := ""
		if e.CreatedAt != nil {
			createdAt = e.CreatedAt.String()
		}
		list = append(list, v1.MemberBalanceLogItem{
			ID:        strconv.FormatUint(e.Id, 10),
			Type:      typeName,
			Amount:    e.ChangeAmount,
			Balance:   e.AfterBalance,
			Remark:    e.Remark,
			CreatedAt: createdAt,
		})
	}

	return &v1.MemberBalanceLogRes{
		Total: total,
		List:  list,
	}, nil
}
