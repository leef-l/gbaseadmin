package activity_step_log

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
	service.RegisterActivityStepLog(New())
}

func New() *sActivityStepLog {
	return &sActivityStepLog{}
}

type sActivityStepLog struct{}

// Create 创建活动步骤提交记录
func (s *sActivityStepLog) Create(ctx context.Context, in *model.ActivityStepLogCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayActivityStepLog.Ctx(ctx).Data(g.Map{
		dao.PlayActivityStepLog.Columns().Id:        id,
		dao.PlayActivityStepLog.Columns().ActivityId: in.ActivityID,
		dao.PlayActivityStepLog.Columns().StepId: in.StepID,
		dao.PlayActivityStepLog.Columns().JoinId: in.JoinID,
		dao.PlayActivityStepLog.Columns().MemberId: in.MemberID,
		dao.PlayActivityStepLog.Columns().StepType: in.StepType,
		dao.PlayActivityStepLog.Columns().SubmitText: in.SubmitText,
		dao.PlayActivityStepLog.Columns().SubmitImage: in.SubmitImage,
		dao.PlayActivityStepLog.Columns().AuditStatus: in.AuditStatus,
		dao.PlayActivityStepLog.Columns().AuditRemark: in.AuditRemark,
		dao.PlayActivityStepLog.Columns().AuditBy: in.AuditBy,
		dao.PlayActivityStepLog.Columns().AuditAt: in.AuditAt,
		dao.PlayActivityStepLog.Columns().CreatedAt: gtime.Now(),
		dao.PlayActivityStepLog.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新活动步骤提交记录
func (s *sActivityStepLog) Update(ctx context.Context, in *model.ActivityStepLogUpdateInput) error {
	data := g.Map{
		dao.PlayActivityStepLog.Columns().ActivityId: in.ActivityID,
		dao.PlayActivityStepLog.Columns().StepId: in.StepID,
		dao.PlayActivityStepLog.Columns().JoinId: in.JoinID,
		dao.PlayActivityStepLog.Columns().MemberId: in.MemberID,
		dao.PlayActivityStepLog.Columns().StepType: in.StepType,
		dao.PlayActivityStepLog.Columns().SubmitText: in.SubmitText,
		dao.PlayActivityStepLog.Columns().SubmitImage: in.SubmitImage,
		dao.PlayActivityStepLog.Columns().AuditStatus: in.AuditStatus,
		dao.PlayActivityStepLog.Columns().AuditRemark: in.AuditRemark,
		dao.PlayActivityStepLog.Columns().AuditBy: in.AuditBy,
		dao.PlayActivityStepLog.Columns().AuditAt: in.AuditAt,
		dao.PlayActivityStepLog.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayActivityStepLog.Ctx(ctx).Where(dao.PlayActivityStepLog.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除活动步骤提交记录
func (s *sActivityStepLog) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayActivityStepLog.Ctx(ctx).Where(dao.PlayActivityStepLog.Columns().Id, id).Data(g.Map{
		dao.PlayActivityStepLog.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取活动步骤提交记录详情
func (s *sActivityStepLog) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.ActivityStepLogDetailOutput, err error) {
	out = &model.ActivityStepLogDetailOutput{}
	err = dao.PlayActivityStepLog.Ctx(ctx).Where(dao.PlayActivityStepLog.Columns().Id, id).Where(dao.PlayActivityStepLog.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	// 查询活动ID关联显示
	if out.ActivityID != 0 {
		val, err := g.DB().Ctx(ctx).Model("play_activity").Where("id", out.ActivityID).Where("deleted_at", nil).Value("title")
		if err == nil {
			out.ActivityTitle = val.String()
		}
	}
	return
}

// List 获取活动步骤提交记录列表
func (s *sActivityStepLog) List(ctx context.Context, in *model.ActivityStepLogListInput) (list []*model.ActivityStepLogListOutput, total int, err error) {
	m := dao.PlayActivityStepLog.Ctx(ctx).Where(dao.PlayActivityStepLog.Columns().DeletedAt, nil)
	if in.StepType > 0 {
		m = m.Where(dao.PlayActivityStepLog.Columns().StepType, in.StepType)
	}
	if in.AuditStatus > 0 {
		m = m.Where(dao.PlayActivityStepLog.Columns().AuditStatus, in.AuditStatus)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayActivityStepLog.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	// 填充关联显示字段
	for _, item := range list {
		if item.ActivityID != 0 {
			val, err := g.DB().Ctx(ctx).Model("play_activity").Where("id", item.ActivityID).Where("deleted_at", nil).Value("title")
			if err == nil {
				item.ActivityTitle = val.String()
			}
		}
	}
	return
}

