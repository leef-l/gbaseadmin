package file

import (
	"context"
	"os"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/upload/internal/dao"
	"gbaseadmin/app/upload/internal/model"
	"gbaseadmin/app/upload/internal/service"
	"gbaseadmin/utility/snowflake"
)

func init() {
	service.RegisterFile(New())
}

func New() *sFile {
	return &sFile{}
}

type sFile struct{}

// Create 创建æ–‡ä»¶è®°å½•
func (s *sFile) Create(ctx context.Context, in *model.FileCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.UploadFile.Ctx(ctx).Data(g.Map{
		dao.UploadFile.Columns().Id:        id,
		dao.UploadFile.Columns().DirId: in.DirID,
		dao.UploadFile.Columns().Name: in.Name,
		dao.UploadFile.Columns().Url: in.URL,
		dao.UploadFile.Columns().Ext: in.Ext,
		dao.UploadFile.Columns().Size: in.Size,
		dao.UploadFile.Columns().Mime: in.Mime,
		dao.UploadFile.Columns().Storage: in.Storage,
		dao.UploadFile.Columns().IsImage: in.IsImage,
		dao.UploadFile.Columns().CreatedAt: gtime.Now(),
		dao.UploadFile.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新æ–‡ä»¶è®°å½•
func (s *sFile) Update(ctx context.Context, in *model.FileUpdateInput) error {
	data := g.Map{
		dao.UploadFile.Columns().DirId: in.DirID,
		dao.UploadFile.Columns().Name: in.Name,
		dao.UploadFile.Columns().Url: in.URL,
		dao.UploadFile.Columns().Ext: in.Ext,
		dao.UploadFile.Columns().Size: in.Size,
		dao.UploadFile.Columns().Mime: in.Mime,
		dao.UploadFile.Columns().Storage: in.Storage,
		dao.UploadFile.Columns().IsImage: in.IsImage,
		dao.UploadFile.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.UploadFile.Ctx(ctx).Where(dao.UploadFile.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 删除文件记录并物理删除文件
func (s *sFile) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	// 先查询文件信息，用于物理删除
	var fileInfo struct {
		Url     string `orm:"url"`
		Storage int    `orm:"storage"`
	}
	err := dao.UploadFile.Ctx(ctx).Where(dao.UploadFile.Columns().Id, id).
		Where(dao.UploadFile.Columns().DeletedAt, nil).Scan(&fileInfo)
	if err != nil {
		return err
	}

	// 软删除记录
	_, err = dao.UploadFile.Ctx(ctx).Where(dao.UploadFile.Columns().Id, id).Data(g.Map{
		dao.UploadFile.Columns().DeletedAt: gtime.Now(),
	}).Update()
	if err != nil {
		return err
	}

	// 物理删除文件
	if fileInfo.Url != "" {
		switch fileInfo.Storage {
		case 1: // 本地存储
			localPath := strings.TrimPrefix(fileInfo.Url, "/")
			_ = os.Remove(localPath)
		case 2: // 阿里云OSS — TODO
		case 3: // 腾讯云COS — TODO
		}
	}
	return nil
}

// Detail 获取æ–‡ä»¶è®°å½•详情
func (s *sFile) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.FileDetailOutput, err error) {
	out = &model.FileDetailOutput{}
	err = dao.UploadFile.Ctx(ctx).Where(dao.UploadFile.Columns().Id, id).Where(dao.UploadFile.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	// 查询æ‰€å±žç›®å½•关联显示
	if out.DirID != 0 {
		val, err := g.DB().Ctx(ctx).Model("upload_dir").Where("id", out.DirID).Where("deleted_at", nil).Value("name")
		if err == nil {
			out.DirName = val.String()
		}
	}
	return
}

// List 获取æ–‡ä»¶è®°å½•列表
func (s *sFile) List(ctx context.Context, in *model.FileListInput) (list []*model.FileListOutput, total int, err error) {
	m := dao.UploadFile.Ctx(ctx).Where(dao.UploadFile.Columns().DeletedAt, nil)
	if in.Storage > 0 {
		m = m.Where(dao.UploadFile.Columns().Storage, in.Storage)
	}
	if in.IsImage > 0 {
		m = m.Where(dao.UploadFile.Columns().IsImage, in.IsImage)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.UploadFile.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	// 填充关联显示字段
	for _, item := range list {
		if item.DirID != 0 {
			val, err := g.DB().Ctx(ctx).Model("upload_dir").Where("id", item.DirID).Where("deleted_at", nil).Value("name")
			if err == nil {
				item.DirName = val.String()
			}
		}
	}
	return
}

