package message

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

var Message = cMessage{}

type cMessage struct{}

// Create 创建会员消息
func (c *cMessage) Create(ctx context.Context, req *v1.MessageCreateReq) (res *v1.MessageCreateRes, err error) {
	err = service.Message().Create(ctx, &model.MessageCreateInput{
		MemberID: req.MemberID,
		Title: req.Title,
		Content: req.Content,
		MsgType: req.MsgType,
		BizID: req.BizID,
		IsRead: req.IsRead,
		Status: req.Status,
	})
	return
}

// Update 更新会员消息
func (c *cMessage) Update(ctx context.Context, req *v1.MessageUpdateReq) (res *v1.MessageUpdateRes, err error) {
	err = service.Message().Update(ctx, &model.MessageUpdateInput{
		ID: req.ID,
		MemberID: req.MemberID,
		Title: req.Title,
		Content: req.Content,
		MsgType: req.MsgType,
		BizID: req.BizID,
		IsRead: req.IsRead,
		Status: req.Status,
	})
	return
}

// Delete 删除会员消息
func (c *cMessage) Delete(ctx context.Context, req *v1.MessageDeleteReq) (res *v1.MessageDeleteRes, err error) {
	err = service.Message().Delete(ctx, req.ID)
	return
}

// BatchDelete 批量删除会员消息
func (c *cMessage) BatchDelete(ctx context.Context, req *v1.MessageBatchDeleteReq) (res *v1.MessageBatchDeleteRes, err error) {
	err = service.Message().BatchDelete(ctx, req.IDs)
	return
}

// Detail 获取会员消息详情
func (c *cMessage) Detail(ctx context.Context, req *v1.MessageDetailReq) (res *v1.MessageDetailRes, err error) {
	res = &v1.MessageDetailRes{}
	res.MessageDetailOutput, err = service.Message().Detail(ctx, req.ID)
	return
}

// List 获取会员消息列表
func (c *cMessage) List(ctx context.Context, req *v1.MessageListReq) (res *v1.MessageListRes, err error) {
	res = &v1.MessageListRes{}
	res.List, res.Total, err = service.Message().List(ctx, &model.MessageListInput{
		PageNum:   req.PageNum,
		PageSize:  req.PageSize,
		OrderBy:   req.OrderBy,
		OrderDir:  req.OrderDir,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Title: req.Title,
	})
	return
}
// Export 导出会员消息
func (c *cMessage) Export(ctx context.Context, req *v1.MessageExportReq) (res *v1.MessageExportRes, err error) {
	list, err := service.Message().Export(ctx, &model.MessageListInput{
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Title: req.Title,
	})
	if err != nil {
		return
	}
	// CSV 导出
	r := g.RequestFromCtx(ctx)
	r.Response.Header().Set("Content-Type", "text/csv; charset=utf-8")
	r.Response.Header().Set("Content-Disposition", `attachment; filename="message.csv"`)
	r.Response.Write("\xEF\xBB\xBF") // UTF-8 BOM
	// 表头
	r.Response.Writeln("接收者会员ID,消息标题,消息内容,消息类型 1=系统通知 2=订单消息 3=活动消息,关联业务ID,是否已读 0=未读 1=已读,状态 1=正常 0=禁用,创建时间")
	// 数据行
	for _, item := range list {
		r.Response.Writefln("%v,%v,%v,%v,%v,%v,%v,%v",
			item.MemberNickname,
			 item.Title,
			 item.Content,
			 item.MsgType,
			 item.BizID,
			 item.IsRead,
			 item.Status,
			item.CreatedAt,
		)
	}
	return
}

