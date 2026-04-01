package playapi

import (
	"context"
	"strconv"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	v1 "gbaseadmin/app/play/api/playapi/v1"
	"gbaseadmin/app/play/internal/dao"
)

type sMessage struct{}

// msgTypeToInt 将前端类型字符串转为数据库整数值
// system=1, order=2, activity=3
func msgTypeToInt(msgType string) int {
	switch msgType {
	case "system":
		return 1
	case "order":
		return 2
	case "activity":
		return 3
	default:
		return 0
	}
}

// msgTypeToString 将数据库整数值转为前端类型字符串
func msgTypeToString(msgType int) string {
	switch msgType {
	case 1:
		return "system"
	case 2:
		return "order"
	case 3:
		return "activity"
	default:
		return "system"
	}
}

// List 我的消息列表
func (s *sMessage) List(ctx context.Context, memberID int64, msgType string, page, pageSize int) (list []v1.MessageItem, total int, err error) {
	cols := dao.PlayMessage.Columns()
	m := dao.PlayMessage.Ctx(ctx).
		Where(cols.MemberId, memberID).
		Where(cols.Status, 1).
		Where(cols.DeletedAt, nil)

	// 按消息类型筛选
	if msgType != "" {
		typeInt := msgTypeToInt(msgType)
		if typeInt == 0 {
			return nil, 0, gerror.Newf("无效的消息类型: %s，支持 system/order/activity", msgType)
		}
		m = m.Where(cols.MsgType, typeInt)
	}

	total, err = m.Count()
	if err != nil {
		return
	}
	if total == 0 {
		list = make([]v1.MessageItem, 0)
		return
	}

	var records []struct {
		Id        uint64      `json:"id"`
		Title     string      `json:"title"`
		Content   string      `json:"content"`
		MsgType   int         `json:"msg_type"`
		IsRead    int         `json:"is_read"`
		CreatedAt *gtime.Time `json:"created_at"`
	}
	err = m.Page(page, pageSize).
		OrderDesc(cols.CreatedAt).
		Scan(&records)
	if err != nil {
		return
	}

	list = make([]v1.MessageItem, 0, len(records))
	for _, r := range records {
		item := v1.MessageItem{
			Id:     strconv.FormatUint(r.Id, 10),
			Type:   msgTypeToString(r.MsgType),
			Title:  r.Title,
			Desc:   r.Content,
			Unread: r.IsRead == 0,
		}
		if r.CreatedAt != nil {
			item.Time = r.CreatedAt.String()
		}
		list = append(list, item)
	}
	return
}

// Read 标记单条消息已读
func (s *sMessage) Read(ctx context.Context, memberID int64, id string) error {
	msgID, err := strconv.ParseUint(id, 10, 64)
	if err != nil || msgID == 0 {
		return gerror.New("消息ID无效")
	}

	cols := dao.PlayMessage.Columns()
	// 验证消息属于当前会员
	count, err := dao.PlayMessage.Ctx(ctx).
		Where(cols.Id, msgID).
		Where(cols.MemberId, memberID).
		Where(cols.DeletedAt, nil).
		Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("消息不存在")
	}

	_, err = dao.PlayMessage.Ctx(ctx).
		Where(cols.Id, msgID).
		Where(cols.MemberId, memberID).
		Where(cols.IsRead, 0).
		Data(g.Map{
			cols.IsRead:    1,
			cols.UpdatedAt: gtime.Now(),
		}).
		Update()
	return err
}

// ReadAll 标记该会员所有未读消息已读
func (s *sMessage) ReadAll(ctx context.Context, memberID int64) error {
	cols := dao.PlayMessage.Columns()
	_, err := dao.PlayMessage.Ctx(ctx).
		Where(cols.MemberId, memberID).
		Where(cols.IsRead, 0).
		Where(cols.DeletedAt, nil).
		Data(g.Map{
			cols.IsRead:    1,
			cols.UpdatedAt: gtime.Now(),
		}).
		Update()
	return err
}
