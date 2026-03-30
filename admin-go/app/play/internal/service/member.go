package service

import (
	"context"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

type IMember interface {
	Create(ctx context.Context, in *model.MemberCreateInput) error
	Update(ctx context.Context, in *model.MemberUpdateInput) error
	Delete(ctx context.Context, id snowflake.JsonInt64) error
	Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.MemberDetailOutput, err error)
	List(ctx context.Context, in *model.MemberListInput) (list []*model.MemberListOutput, total int, err error)
}

var localMember IMember

func Member() IMember {
	return localMember
}

func RegisterMember(i IMember) {
	localMember = i
}
