package playapi

import (
	"context"
	"encoding/base64"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"

	v1 "gbaseadmin/api/playapi/v1"
	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/utility/jwt"
	"gbaseadmin/utility/snowflake"
)

type sAuth struct{}

func (s *sAuth) Login(ctx context.Context, req *v1.AuthLoginReq) (*v1.AuthLoginRes, error) {
	// 开发阶段：固定验证码 123456
	if req.Code != "123456" {
		return nil, fmt.Errorf("验证码错误或已过期")
	}

	// 查询会员是否存在
	memberColumns := dao.PlayMember.Columns()
	record, err := dao.PlayMember.Ctx(ctx).
		Where(memberColumns.Phone, req.Phone).
		Where(memberColumns.DeletedAt, nil).
		One()
	if err != nil {
		return nil, err
	}

	var (
		memberId int64
		isCoach  int
		coachId  int64
		isNew    bool
	)

	if record.IsEmpty() {
		// 自动注册
		isNew = true
		id := snowflake.Generate()
		memberId = int64(id)

		// 查询最低等级
		levelId, _ := dao.PlayMemberLevel.Ctx(ctx).
			Where(dao.PlayMemberLevel.Columns().DeletedAt, nil).
			Where(dao.PlayMemberLevel.Columns().Status, 1).
			OrderAsc(dao.PlayMemberLevel.Columns().Level).
			Value(dao.PlayMemberLevel.Columns().Id)

		phone4 := req.Phone
		if len(phone4) >= 4 {
			phone4 = phone4[len(phone4)-4:]
		}

		_, err = dao.PlayMember.Ctx(ctx).Data(g.Map{
			memberColumns.Id:            memberId,
			memberColumns.Phone:         req.Phone,
			memberColumns.Nickname:      "用户" + phone4,
			memberColumns.Avatar:        "",
			memberColumns.Gender:        0,
			memberColumns.MemberLevelId: levelId.Int64(),
			memberColumns.Exp:           0,
			memberColumns.Balance:       0,
			memberColumns.IsCoach:       0,
			memberColumns.Status:        1,
			memberColumns.LastLoginAt:   gtime.Now(),
			memberColumns.CreatedAt:     gtime.Now(),
			memberColumns.UpdatedAt:     gtime.Now(),
		}).Insert()
		if err != nil {
			return nil, err
		}
	} else {
		memberId = record[memberColumns.Id].Int64()
		isCoach = record[memberColumns.IsCoach].Int()
		status := record[memberColumns.Status].Int()
		if status == 0 {
			return nil, fmt.Errorf("账号已被禁用")
		}

		// 如果是陪玩师，查询 coach_id
		if isCoach == 1 {
			coachRecord, _ := dao.PlayCoach.Ctx(ctx).
				Where(dao.PlayCoach.Columns().MemberId, memberId).
				Where(dao.PlayCoach.Columns().DeletedAt, nil).
				One()
			if !coachRecord.IsEmpty() {
				coachId = coachRecord[dao.PlayCoach.Columns().Id].Int64()
			}
		}

		// 更新最后登录时间
		_, _ = dao.PlayMember.Ctx(ctx).
			Where(memberColumns.Id, memberId).
			Data(g.Map{memberColumns.LastLoginAt: gtime.Now()}).
			Update()
	}

	// 签发 JWT
	token, err := jwt.GenerateMemberToken(memberId, req.Phone, isCoach, coachId, "member")
	if err != nil {
		return nil, err
	}

	// 生成 RefreshToken（简单 base64 编码 memberId）
	refreshToken := base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(memberId, 10)))

	return &v1.AuthLoginRes{
		Token:        token,
		RefreshToken: refreshToken,
		ExpiresIn:    86400,
		IsNew:        isNew,
	}, nil
}

func (s *sAuth) SendCode(ctx context.Context, req *v1.AuthSendCodeReq) (*v1.AuthSendCodeRes, error) {
	// 生成6位随机验证码
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	// 开发阶段：仅打印日志，不实际发送短信
	glog.Infof(ctx, "发送验证码: phone=%s, scene=%s, code=%s", req.Phone, req.Scene, code)
	return &v1.AuthSendCodeRes{}, nil
}

func (s *sAuth) RefreshToken(ctx context.Context, req *v1.AuthRefreshTokenReq) (*v1.AuthRefreshTokenRes, error) {
	// 解析 RefreshToken（base64 编码的 memberId）
	decoded, err := base64.StdEncoding.DecodeString(req.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("刷新令牌无效")
	}
	memberId, err := strconv.ParseInt(string(decoded), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("刷新令牌无效")
	}

	// 查询会员信息
	memberColumns := dao.PlayMember.Columns()
	record, err := dao.PlayMember.Ctx(ctx).
		Where(memberColumns.Id, memberId).
		Where(memberColumns.DeletedAt, nil).
		One()
	if err != nil {
		return nil, err
	}
	if record.IsEmpty() {
		return nil, fmt.Errorf("刷新令牌无效")
	}
	if record[memberColumns.Status].Int() == 0 {
		return nil, fmt.Errorf("账号已被禁用")
	}

	phone := record[memberColumns.Phone].String()
	isCoach := record[memberColumns.IsCoach].Int()
	var coachId int64
	if isCoach == 1 {
		coachRecord, _ := dao.PlayCoach.Ctx(ctx).
			Where(dao.PlayCoach.Columns().MemberId, memberId).
			Where(dao.PlayCoach.Columns().DeletedAt, nil).
			One()
		if !coachRecord.IsEmpty() {
			coachId = coachRecord[dao.PlayCoach.Columns().Id].Int64()
		}
	}

	// 重新签发 JWT
	token, err := jwt.GenerateMemberToken(memberId, phone, isCoach, coachId, "member")
	if err != nil {
		return nil, err
	}

	// 生成新 RefreshToken
	newRefreshToken := base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(memberId, 10)))

	return &v1.AuthRefreshTokenRes{
		Token:        token,
		RefreshToken: newRefreshToken,
		ExpiresIn:    86400,
	}, nil
}

func (s *sAuth) WxLogin(ctx context.Context, req *v1.AuthWxLoginReq) (*v1.AuthWxLoginRes, error) {
	return nil, fmt.Errorf("微信登录暂未开通")
}

func (s *sAuth) AlipayLogin(ctx context.Context, req *v1.AuthAlipayLoginReq) (*v1.AuthAlipayLoginRes, error) {
	return nil, fmt.Errorf("支付宝登录暂未开通")
}
