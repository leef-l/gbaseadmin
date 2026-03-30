package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type IMemberLevel interface {
	Create(ctx context.Context, in *model.MemberLevelCreateInput) error
	Update(ctx context.Context, in *model.MemberLevelUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.MemberLevelDetailOutput, err error)
	List(ctx context.Context, in *model.MemberLevelListInput) (list []*model.MemberLevelListOutput, total int, err error)
}

var localMemberLevel IMemberLevel

func MemberLevel() IMemberLevel {
	return localMemberLevel
}

func RegisterMemberLevel(i IMemberLevel) {
	localMemberLevel = i
}
