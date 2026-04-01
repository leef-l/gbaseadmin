package playapi

import (
	"context"
	"strconv"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	v1 "gbaseadmin/app/play/api/playapi/v1"
	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
	"gbaseadmin/utility/snowflake"
)

type sActivity struct{}

// List 活动列表
func (s *sActivity) List(ctx context.Context, page, pageSize int) (list []v1.ActivityListApiItem, total int, err error) {
	m := dao.PlayActivity.Ctx(ctx).
		Where(dao.PlayActivity.Columns().Status, 1).
		WhereGTE(dao.PlayActivity.Columns().EndAt, gtime.Now()).
		Where(dao.PlayActivity.Columns().DeletedAt, nil)

	total, err = m.Count()
	if err != nil {
		return
	}

	var records []struct {
		Id          uint64      `json:"id"`
		Title       string      `json:"title"`
		CoverImage  string      `json:"cover_image"`
		DescContent string      `json:"desc_content"`
		Type        int         `json:"type"`
		StartAt     *gtime.Time `json:"start_at"`
		EndAt       *gtime.Time `json:"end_at"`
		JoinNum     int         `json:"join_num"`
	}
	err = m.Page(page, pageSize).
		OrderDesc(dao.PlayActivity.Columns().StartAt).
		Scan(&records)
	if err != nil {
		return
	}
	list = make([]v1.ActivityListApiItem, 0, len(records))
	for _, r := range records {
		item := v1.ActivityListApiItem{
			ActivityID:  strconv.FormatUint(r.Id, 10),
			Title:       r.Title,
			Cover:       r.CoverImage,
			Description: r.DescContent,
			Type:        r.Type,
			JoinCount:   r.JoinNum,
		}
		if r.StartAt != nil {
			item.StartTime = r.StartAt.String()
		}
		if r.EndAt != nil {
			item.EndTime = r.EndAt.String()
		}
		list = append(list, item)
	}
	return
}

// Detail 活动详情
func (s *sActivity) Detail(ctx context.Context, activityID string, memberID int64) (out *v1.ActivityDetailApiRes, err error) {
	aid, _ := strconv.ParseUint(activityID, 10, 64)
	act, err := dao.PlayActivity.Ctx(ctx).
		Where(dao.PlayActivity.Columns().Id, aid).
		Where(dao.PlayActivity.Columns().Status, 1).
		Where(dao.PlayActivity.Columns().DeletedAt, nil).
		One()
	if err != nil {
		return
	}
	if act.IsEmpty() {
		err = gerror.New("活动不存在")
		return
	}
	out = &v1.ActivityDetailApiRes{
		ActivityID: activityID,
		Title:      act[dao.PlayActivity.Columns().Title].String(),
		Cover:      act[dao.PlayActivity.Columns().CoverImage].String(),
		Content:    act[dao.PlayActivity.Columns().DescContent].String(),
		Type:       act[dao.PlayActivity.Columns().Type].Int(),
		JoinCount:  act[dao.PlayActivity.Columns().JoinNum].Int(),
	}
	if act[dao.PlayActivity.Columns().StartAt].GTime() != nil {
		out.StartTime = act[dao.PlayActivity.Columns().StartAt].String()
	}
	if act[dao.PlayActivity.Columns().EndAt].GTime() != nil {
		out.EndTime = act[dao.PlayActivity.Columns().EndAt].String()
	}
	// 查询步骤
	var steps []struct {
		Id          uint64 `json:"id"`
		StepNum     int    `json:"step_num"`
		Title       string `json:"title"`
		DescContent string `json:"desc_content"`
		StepType    int    `json:"step_type"`
		ExampleText string `json:"example_text"`
		StepImage   string `json:"step_image"`
		IsRequired  int    `json:"is_required"`
	}
	err = dao.PlayActivityStep.Ctx(ctx).
		Where(dao.PlayActivityStep.Columns().ActivityId, aid).
		Where(dao.PlayActivityStep.Columns().DeletedAt, nil).
		OrderAsc(dao.PlayActivityStep.Columns().StepNum).
		Scan(&steps)
	if err != nil {
		return
	}
	out.Steps = make([]v1.ActivityStepApiItem, 0, len(steps))
	for _, st := range steps {
		out.Steps = append(out.Steps, v1.ActivityStepApiItem{
			StepID:      strconv.FormatUint(st.Id, 10),
			StepNo:      st.StepNum,
			Title:       st.Title,
			Description: st.DescContent,
			StepType:    st.StepType,
			ExampleText: st.ExampleText,
			StepImage:   st.StepImage,
			IsRequired:  st.IsRequired,
		})
	}
	// 查询奖励
	var rewards []struct {
		Id          uint64 `json:"id"`
		RewardType  int    `json:"reward_type"`
		RewardValue int64  `json:"reward_value"`
		RewardName  string `json:"reward_name"`
	}
	err = dao.PlayActivityReward.Ctx(ctx).
		Where(dao.PlayActivityReward.Columns().ActivityId, aid).
		Where(dao.PlayActivityReward.Columns().DeletedAt, nil).
		OrderAsc(dao.PlayActivityReward.Columns().Sort).
		Scan(&rewards)
	if err != nil {
		return
	}
	out.Rewards = make([]v1.ActivityRewardApiItem, 0, len(rewards))
	for _, rw := range rewards {
		out.Rewards = append(out.Rewards, v1.ActivityRewardApiItem{
			RewardID:    strconv.FormatUint(rw.Id, 10),
			RewardName:  rw.RewardName,
			RewardType:  rw.RewardType,
			RewardValue: rw.RewardValue,
		})
	}

	// 查询当前用户报名状态及进度
	if memberID > 0 {
		// 优先取未完成的报名（进行中），没有则取最近一条
		join, e := dao.PlayActivityJoin.Ctx(ctx).
			Where(dao.PlayActivityJoin.Columns().ActivityId, aid).
			Where(dao.PlayActivityJoin.Columns().MemberId, memberID).
			Where(dao.PlayActivityJoin.Columns().DeletedAt, nil).
			WhereLT(dao.PlayActivityJoin.Columns().JoinStatus, 2).
			OrderDesc(dao.PlayActivityJoin.Columns().CreatedAt).
			One()
		if e == nil && join.IsEmpty() {
			// 没有进行中的，取最近完成的
			join, e = dao.PlayActivityJoin.Ctx(ctx).
				Where(dao.PlayActivityJoin.Columns().ActivityId, aid).
				Where(dao.PlayActivityJoin.Columns().MemberId, memberID).
				Where(dao.PlayActivityJoin.Columns().DeletedAt, nil).
				OrderDesc(dao.PlayActivityJoin.Columns().CreatedAt).
				One()
		}
		if e == nil && !join.IsEmpty() {
			out.HasJoined = true
			joinStatus := join[dao.PlayActivityJoin.Columns().JoinStatus].Int()
			progress := &v1.ActivityProgressInfo{
				CurrentStep: join[dao.PlayActivityJoin.Columns().CurrentStep].Int(),
				IsCompleted: joinStatus >= 2,
				IsRewarded:  joinStatus >= 3,
			}
			if join[dao.PlayActivityJoin.Columns().CreatedAt].GTime() != nil {
				progress.JoinedAt = join[dao.PlayActivityJoin.Columns().CreatedAt].String()
			}
			out.MyProgress = progress
		}
	}

	return
}

// Join 报名参与活动
func (s *sActivity) Join(ctx context.Context, memberID int64, activityID string) error {
	aid, _ := strconv.ParseUint(activityID, 10, 64)
	act, err := dao.PlayActivity.Ctx(ctx).
		Where(dao.PlayActivity.Columns().Id, aid).
		Where(dao.PlayActivity.Columns().Status, 1).
		Where(dao.PlayActivity.Columns().DeletedAt, nil).
		One()
	if err != nil {
		return err
	}
	if act.IsEmpty() {
		return gerror.New("活动不存在")
	}
	now := gtime.Now()
	if act[dao.PlayActivity.Columns().StartAt].GTime() != nil && now.Before(act[dao.PlayActivity.Columns().StartAt].GTime()) {
		return gerror.New("活动未开始")
	}
	if act[dao.PlayActivity.Columns().EndAt].GTime() != nil && now.After(act[dao.PlayActivity.Columns().EndAt].GTime()) {
		return gerror.New("活动已结束")
	}
	// 检查是否有未完成的报名（join_status < 2）
	cnt, err := dao.PlayActivityJoin.Ctx(ctx).
		Where(dao.PlayActivityJoin.Columns().ActivityId, aid).
		Where(dao.PlayActivityJoin.Columns().MemberId, memberID).
		Where(dao.PlayActivityJoin.Columns().DeletedAt, nil).
		WhereLT(dao.PlayActivityJoin.Columns().JoinStatus, 2).
		Count()
	if err != nil {
		return err
	}
	if cnt > 0 {
		return gerror.New("您有未完成的报名，请先完成或取消后再报名")
	}
	return dao.PlayActivityJoin.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		id := snowflake.Generate()
		_, err := dao.PlayActivityJoin.Ctx(ctx).Data(g.Map{
			dao.PlayActivityJoin.Columns().Id:          id,
			dao.PlayActivityJoin.Columns().ActivityId:  aid,
			dao.PlayActivityJoin.Columns().MemberId:    memberID,
			dao.PlayActivityJoin.Columns().JoinStatus:  0,
			dao.PlayActivityJoin.Columns().CurrentStep: 0,
			dao.PlayActivityJoin.Columns().CreatedAt:   gtime.Now(),
			dao.PlayActivityJoin.Columns().UpdatedAt:   gtime.Now(),
		}).Insert()
		if err != nil {
			return err
		}
		_, err = dao.PlayActivity.Ctx(ctx).
			Where(dao.PlayActivity.Columns().Id, aid).
			Increment(dao.PlayActivity.Columns().JoinNum, 1)
		return err
	})
}

// CompleteStep 完成活动步骤
func (s *sActivity) CompleteStep(ctx context.Context, memberID int64, activityID, stepID, imageUrl, submitText string) (currentStep int, isCompleted bool, err error) {
	aid, _ := strconv.ParseUint(activityID, 10, 64)
	sid, _ := strconv.ParseUint(stepID, 10, 64)
	// 查询参与记录
	join, err := dao.PlayActivityJoin.Ctx(ctx).
		Where(dao.PlayActivityJoin.Columns().ActivityId, aid).
		Where(dao.PlayActivityJoin.Columns().MemberId, memberID).
		Where(dao.PlayActivityJoin.Columns().DeletedAt, nil).
		One()
	if err != nil {
		return
	}
	if join.IsEmpty() {
		err = gerror.New("您未参与该活动")
		return
	}
	if join[dao.PlayActivityJoin.Columns().JoinStatus].Int() >= 2 {
		err = gerror.New("活动已完成")
		return
	}
	// 查询步骤
	step, err := dao.PlayActivityStep.Ctx(ctx).
		Where(dao.PlayActivityStep.Columns().Id, sid).
		Where(dao.PlayActivityStep.Columns().ActivityId, aid).
		Where(dao.PlayActivityStep.Columns().DeletedAt, nil).
		One()
	if err != nil {
		return
	}
	if step.IsEmpty() {
		err = gerror.New("步骤不存在")
		return
	}
	curStep := join[dao.PlayActivityJoin.Columns().CurrentStep].Int()
	stepNum := step[dao.PlayActivityStep.Columns().StepNum].Int()
	if stepNum != curStep+1 {
		err = gerror.New("请按顺序完成步骤")
		return
	}

	// 查询总步骤数
	totalSteps, err := dao.PlayActivityStep.Ctx(ctx).
		Where(dao.PlayActivityStep.Columns().ActivityId, aid).
		Where(dao.PlayActivityStep.Columns().DeletedAt, nil).
		Count()
	if err != nil {
		return
	}
	currentStep = stepNum
	isCompleted = currentStep >= totalSteps
	joinUID := join[dao.PlayActivityJoin.Columns().Id].Uint64()
	data := g.Map{
		dao.PlayActivityJoin.Columns().CurrentStep: currentStep,
		dao.PlayActivityJoin.Columns().UpdatedAt:   gtime.Now(),
	}
	if isCompleted {
		data[dao.PlayActivityJoin.Columns().JoinStatus] = 2
		data[dao.PlayActivityJoin.Columns().FinishAt] = gtime.Now()
	} else if curStep == 0 {
		data[dao.PlayActivityJoin.Columns().JoinStatus] = 1
	}
	err = dao.PlayActivityJoin.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, e := dao.PlayActivityJoin.Ctx(ctx).
			Where(dao.PlayActivityJoin.Columns().Id, joinUID).
			Data(data).Update()
		if e != nil {
			return e
		}
		logID := snowflake.Generate()
		_, e = dao.PlayActivityStepLog.Ctx(ctx).Data(g.Map{
			dao.PlayActivityStepLog.Columns().Id:          logID,
			dao.PlayActivityStepLog.Columns().ActivityId:  aid,
			dao.PlayActivityStepLog.Columns().StepId:      sid,
			dao.PlayActivityStepLog.Columns().JoinId:      joinUID,
			dao.PlayActivityStepLog.Columns().MemberId:    memberID,
			dao.PlayActivityStepLog.Columns().StepType:    step[dao.PlayActivityStep.Columns().StepType].Int(),
			dao.PlayActivityStepLog.Columns().SubmitImage: imageUrl,
			dao.PlayActivityStepLog.Columns().SubmitText:  submitText,
			dao.PlayActivityStepLog.Columns().AuditStatus: 0,
			dao.PlayActivityStepLog.Columns().CreatedAt:   gtime.Now(),
			dao.PlayActivityStepLog.Columns().UpdatedAt:   gtime.Now(),
		}).Insert()
		return e
	})
	return
}

// ClaimReward 领取奖励（一次性发放该活动的所有奖励）
func (s *sActivity) ClaimReward(ctx context.Context, memberID int64, activityID string) error {
	aid, _ := strconv.ParseUint(activityID, 10, 64)
	// 查询参与记录
	join, err := dao.PlayActivityJoin.Ctx(ctx).
		Where(dao.PlayActivityJoin.Columns().ActivityId, aid).
		Where(dao.PlayActivityJoin.Columns().MemberId, memberID).
		Where(dao.PlayActivityJoin.Columns().DeletedAt, nil).
		One()
	if err != nil {
		return err
	}
	if join.IsEmpty() {
		return gerror.New("您未参与该活动")
	}
	if join[dao.PlayActivityJoin.Columns().JoinStatus].Int() < 2 {
		return gerror.New("请先完成活动")
	}
	if join[dao.PlayActivityJoin.Columns().JoinStatus].Int() == 3 {
		return gerror.New("您已领取过奖励")
	}
	// 查询该活动所有奖励
	var rewards []struct {
		Id          uint64 `json:"id"`
		RewardType  int    `json:"reward_type"`
		RewardValue int64  `json:"reward_value"`
		RewardName  string `json:"reward_name"`
	}
	err = dao.PlayActivityReward.Ctx(ctx).
		Where(dao.PlayActivityReward.Columns().ActivityId, aid).
		Where(dao.PlayActivityReward.Columns().DeletedAt, nil).
		Scan(&rewards)
	if err != nil {
		return err
	}
	if len(rewards) == 0 {
		return gerror.New("该活动暂无奖励")
	}

	joinID := join[dao.PlayActivityJoin.Columns().Id].Uint64()

	return dao.PlayActivityJoin.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, rw := range rewards {
			switch rw.RewardType {
			case 1: // 余额
				err := service.BalanceLogEnhance().AddLog(ctx, &model.AddBalanceLogInput{
					MemberID:     snowflake.JsonInt64(memberID),
					BizType:      4, // 活动赠送
					BizID:        snowflake.JsonInt64(rw.Id),
					ChangeAmount: rw.RewardValue,
					Remark:       "活动奖励-" + rw.RewardName,
				})
				if err != nil {
					return err
				}
			case 2: // 优惠券 - rewardValue 为优惠券ID
				couponID := uint64(rw.RewardValue)
				coupon, e := dao.PlayCoupon.Ctx(ctx).Where(dao.PlayCoupon.Columns().Id, couponID).One()
				if e != nil || coupon.IsEmpty() {
					g.Log().Warningf(ctx, "活动奖励优惠券不存在: couponID=%d, rewardID=%d", couponID, rw.Id)
					continue
				}
				cmID := snowflake.Generate()
				_, err := dao.PlayCouponMember.Ctx(ctx).Data(g.Map{
					dao.PlayCouponMember.Columns().Id:        cmID,
					dao.PlayCouponMember.Columns().CouponId:  couponID,
					dao.PlayCouponMember.Columns().MemberId:  memberID,
					dao.PlayCouponMember.Columns().UseStatus: 0,
					dao.PlayCouponMember.Columns().ClaimAt:   gtime.Now(),
					dao.PlayCouponMember.Columns().ExpireAt:  coupon[dao.PlayCoupon.Columns().ValidEndAt].GTime(),
					dao.PlayCouponMember.Columns().CreatedAt: gtime.Now(),
					dao.PlayCouponMember.Columns().UpdatedAt: gtime.Now(),
				}).Insert()
				if err != nil {
					return err
				}
			case 3: // 经验值
				_, err := dao.PlayMember.Ctx(ctx).
					Where(dao.PlayMember.Columns().Id, memberID).
					Increment(dao.PlayMember.Columns().Exp, rw.RewardValue)
				if err != nil {
					return err
				}
			case 4: // 会员天数（暂不支持，记录日志跳过）
				g.Log().Infof(ctx, "活动奖励-会员天数暂未实现: memberID=%d, rewardValue=%d天", memberID, rw.RewardValue)
			}
		}
		// 更新参与记录为已领奖
		_, err := dao.PlayActivityJoin.Ctx(ctx).
			Where(dao.PlayActivityJoin.Columns().Id, joinID).
			Data(g.Map{
				dao.PlayActivityJoin.Columns().JoinStatus: 3,
				dao.PlayActivityJoin.Columns().RewardAt:   gtime.Now(),
				dao.PlayActivityJoin.Columns().UpdatedAt:  gtime.Now(),
			}).Update()
		return err
	})
}

// MyJoins 我参与的活动列表
func (s *sActivity) MyJoins(ctx context.Context, memberID int64, page, pageSize int) (list []v1.ActivityMyJoinsItem, total int, err error) {
	m := dao.PlayActivityJoin.Ctx(ctx).
		Where(dao.PlayActivityJoin.Columns().MemberId, memberID).
		Where(dao.PlayActivityJoin.Columns().DeletedAt, nil)

	total, err = m.Count()
	if err != nil {
		return
	}

	var records []struct {
		Id          uint64      `json:"id"`
		ActivityId  uint64      `json:"activity_id"`
		JoinStatus  int         `json:"join_status"`
		CurrentStep int         `json:"current_step"`
		CreatedAt   *gtime.Time `json:"created_at"`
	}
	err = m.Page(page, pageSize).
		OrderDesc(dao.PlayActivityJoin.Columns().CreatedAt).
		Scan(&records)
	if err != nil {
		return
	}

	list = make([]v1.ActivityMyJoinsItem, 0, len(records))
	for _, r := range records {
		act, e := dao.PlayActivity.Ctx(ctx).
			Where(dao.PlayActivity.Columns().Id, r.ActivityId).
			One()
		if e != nil || act.IsEmpty() {
			continue
		}
		// 查询总步骤数
		totalSteps, _ := dao.PlayActivityStep.Ctx(ctx).
			Where(dao.PlayActivityStep.Columns().ActivityId, r.ActivityId).
			Where(dao.PlayActivityStep.Columns().DeletedAt, nil).
			Count()

		item := v1.ActivityMyJoinsItem{
			ActivityID:  strconv.FormatUint(r.ActivityId, 10),
			Title:       act[dao.PlayActivity.Columns().Title].String(),
			Cover:       act[dao.PlayActivity.Columns().CoverImage].String(),
			Type:        act[dao.PlayActivity.Columns().Type].Int(),
			CurrentStep: r.CurrentStep,
			TotalSteps:  totalSteps,
			JoinStatus:  r.JoinStatus,
		}
		if act[dao.PlayActivity.Columns().StartAt].GTime() != nil {
			item.StartTime = act[dao.PlayActivity.Columns().StartAt].String()
		}
		if act[dao.PlayActivity.Columns().EndAt].GTime() != nil {
			item.EndTime = act[dao.PlayActivity.Columns().EndAt].String()
		}
		if r.CreatedAt != nil {
			item.JoinedAt = r.CreatedAt.String()
		}
		list = append(list, item)
	}
	return
}

// Quit 取消报名
func (s *sActivity) Quit(ctx context.Context, memberID int64, activityID string) error {
	aid, _ := strconv.ParseUint(activityID, 10, 64)
	join, err := dao.PlayActivityJoin.Ctx(ctx).
		Where(dao.PlayActivityJoin.Columns().ActivityId, aid).
		Where(dao.PlayActivityJoin.Columns().MemberId, memberID).
		Where(dao.PlayActivityJoin.Columns().DeletedAt, nil).
		One()
	if err != nil {
		return err
	}
	if join.IsEmpty() {
		return gerror.New("您未参与该活动")
	}
	if join[dao.PlayActivityJoin.Columns().JoinStatus].Int() >= 2 {
		return gerror.New("活动已完成，无法取消")
	}
	return dao.PlayActivityJoin.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		joinID := join[dao.PlayActivityJoin.Columns().Id].Uint64()
		_, err := dao.PlayActivityJoin.Ctx(ctx).
			Where(dao.PlayActivityJoin.Columns().Id, joinID).
			Data(g.Map{
				dao.PlayActivityJoin.Columns().DeletedAt: gtime.Now(),
				dao.PlayActivityJoin.Columns().UpdatedAt: gtime.Now(),
			}).Update()
		if err != nil {
			return err
		}
		_, err = dao.PlayActivity.Ctx(ctx).
			Where(dao.PlayActivity.Columns().Id, aid).
			Decrement(dao.PlayActivity.Columns().JoinNum, 1)
		return err
	})
}
