package activity_step

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
	service.RegisterActivityStep(New())
}

func New() *sActivityStep {
	return &sActivityStep{}
}

type sActivityStep struct{}

// Create 创建活动步骤表
func (s *sActivityStep) Create(ctx context.Context, in *model.ActivityStepCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayActivityStep.Ctx(ctx).Data(g.Map{
		dao.PlayActivityStep.Columns().Id:          id,
		dao.PlayActivityStep.Columns().ActivityId:   in.ActivityID,
		dao.PlayActivityStep.Columns().StepNum:      in.StepNum,
		dao.PlayActivityStep.Columns().Title:        in.Title,
		dao.PlayActivityStep.Columns().StepType:     in.StepType,
		dao.PlayActivityStep.Columns().ExampleText:  in.ExampleText,
		dao.PlayActivityStep.Columns().DescContent:  in.DescContent,
		dao.PlayActivityStep.Columns().StepImage:    in.StepImage,
		dao.PlayActivityStep.Columns().IsRequired:   in.IsRequired,
		dao.PlayActivityStep.Columns().Sort:         in.Sort,
		dao.PlayActivityStep.Columns().CreatedAt:    gtime.Now(),
		dao.PlayActivityStep.Columns().UpdatedAt:    gtime.Now(),
	}).Insert()
	return err
}

// Update 更新活动步骤表
func (s *sActivityStep) Update(ctx context.Context, in *model.ActivityStepUpdateInput) error {
	data := g.Map{
		dao.PlayActivityStep.Columns().ActivityId:   in.ActivityID,
		dao.PlayActivityStep.Columns().StepNum:      in.StepNum,
		dao.PlayActivityStep.Columns().Title:        in.Title,
		dao.PlayActivityStep.Columns().StepType:     in.StepType,
		dao.PlayActivityStep.Columns().ExampleText:  in.ExampleText,
		dao.PlayActivityStep.Columns().DescContent:  in.DescContent,
		dao.PlayActivityStep.Columns().StepImage:    in.StepImage,
		dao.PlayActivityStep.Columns().IsRequired:   in.IsRequired,
		dao.PlayActivityStep.Columns().Sort:         in.Sort,
		dao.PlayActivityStep.Columns().UpdatedAt:    gtime.Now(),
	}
	_, err := dao.PlayActivityStep.Ctx(ctx).Where(dao.PlayActivityStep.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除活动步骤表
func (s *sActivityStep) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayActivityStep.Ctx(ctx).Where(dao.PlayActivityStep.Columns().Id, id).Data(g.Map{
		dao.PlayActivityStep.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取活动步骤表详情
func (s *sActivityStep) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.ActivityStepDetailOutput, err error) {
	out = &model.ActivityStepDetailOutput{}
	err = dao.PlayActivityStep.Ctx(ctx).Where(dao.PlayActivityStep.Columns().Id, id).Where(dao.PlayActivityStep.Columns().DeletedAt, nil).Scan(out)
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

// List 获取活动步骤表列表
func (s *sActivityStep) List(ctx context.Context, in *model.ActivityStepListInput) (list []*model.ActivityStepListOutput, total int, err error) {
	m := dao.PlayActivityStep.Ctx(ctx).Where(dao.PlayActivityStep.Columns().DeletedAt, nil)
	if in.ActivityID != 0 {
		m = m.Where(dao.PlayActivityStep.Columns().ActivityId, in.ActivityID)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayActivityStep.Columns().Id).Scan(&list)
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

