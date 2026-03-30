package playapi

import (
	"context"

	v1 "gbaseadmin/api/playapi/v1"
)

type sMessage struct{}

// List 我的消息列表（暂无DB表，返回空列表）
func (s *sMessage) List(ctx context.Context, memberID int64, msgType string, page, pageSize int) (list []v1.MessageItem, total int, err error) {
	list = make([]v1.MessageItem, 0)
	total = 0
	return
}

// Read 标记单条消息已读（暂无DB表，直接返回成功）
func (s *sMessage) Read(ctx context.Context, memberID int64, id string) error {
	return nil
}

// ReadAll 标记全部消息已读（暂无DB表，直接返回成功）
func (s *sMessage) ReadAll(ctx context.Context, memberID int64) error {
	return nil
}
