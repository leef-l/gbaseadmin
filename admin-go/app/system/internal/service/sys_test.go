package service

import (
	"context"
	"gbaseadmin/app/system/internal/model"
	"gbaseadmin/utility/snowflake"
)

type ISysTest interface {
	Create(ctx context.Context, in *model.SysTestCreateInput) error
	Update(ctx context.Context, in *model.SysTestUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.SysTestDetailOutput, err error)
	List(ctx context.Context, in *model.SysTestListInput) (list []*model.SysTestListOutput, total int, err error)
	Tree(ctx context.Context) (tree []*model.SysTestTreeOutput, err error)
}

var localSysTest ISysTest

func SysTest() ISysTest {
	return localSysTest
}

func RegisterSysTest(i ISysTest) {
	localSysTest = i
}
