package uploader

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/upload/internal/dao"
	"gbaseadmin/app/upload/internal/model"
	"gbaseadmin/app/upload/internal/service"
	"gbaseadmin/utility/snowflake"
)

func init() {
	service.RegisterUploader(&sUploader{})
}

type sUploader struct{}

var imageExts = map[string]bool{
	"jpg": true, "jpeg": true, "png": true, "gif": true,
	"webp": true, "svg": true, "bmp": true,
}

func (s *sUploader) Upload(ctx context.Context) (*model.UploadOutput, error) {
	r := g.RequestFromCtx(ctx)

	// 获取上传文件
	file := r.GetUploadFile("file")
	if file == nil {
		return nil, fmt.Errorf("请选择要上传的文件")
	}

	// 读取默认上传配置
	maxSize := int64(10 * 1024 * 1024) // 默认10MB
	storageType := 1                    // 默认本地
	localPath := "resource/upload"

	var configRecord map[string]interface{}
	err := dao.UploadConfig.Ctx(ctx).
		Where("is_default", 1).
		Where("status", 1).
		Where("deleted_at IS NULL").
		Scan(&configRecord)
	if err == nil && configRecord != nil {
		if ms, ok := configRecord["max_size"]; ok {
			if v, ok := ms.(int64); ok && v > 0 {
				maxSize = v * 1024 * 1024
			}
		}
		if st, ok := configRecord["storage"]; ok {
			if v, ok := st.(int64); ok {
				storageType = int(v)
			}
		}
		if lp, ok := configRecord["local_path"]; ok {
			if v, ok := lp.(string); ok && v != "" {
				localPath = v
			}
		}
	}

	// 验证文件大小
	if file.Size > maxSize {
		return nil, fmt.Errorf("文件大小超过限制（最大 %dMB）", maxSize/1024/1024)
	}

	// 解析文件信息
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != "" {
		ext = ext[1:] // 去掉点号
	}
	isImage := 0
	if imageExts[ext] {
		isImage = 1
	}

	// 获取目录ID
	dirId := r.Get("dirId").Int64()

	// 生成唯一文件名
	now := time.Now()
	dateDir := now.Format("2006-01-02")
	uniqueName := fmt.Sprintf("%d%d%04d.%s", now.UnixMilli(), now.UnixNano()%1000, rand.Intn(10000), ext)

	// 本地存储
	savePath := filepath.Join(localPath, dateDir)
	if err := os.MkdirAll(savePath, 0755); err != nil {
		return nil, fmt.Errorf("创建目录失败: %v", err)
	}

	file.Filename = uniqueName
	fullPath := filepath.Join(savePath, uniqueName)
	_, err = file.Save(savePath)
	if err != nil {
		return nil, fmt.Errorf("保存文件失败: %v", err)
	}

	// URL 路径：静态路由 /upload -> resource/upload，所以 URL 需要去掉 localPath 前缀
	relativePath := filepath.Join(dateDir, uniqueName)
	url := "/upload/" + strings.ReplaceAll(relativePath, "\\", "/")

	// 生成ID并写入数据库
	id := snowflake.Generate()
	_, err = dao.UploadFile.Ctx(ctx).Data(g.Map{
		"id":         id,
		"dir_id":     dirId,
		"name":       file.Filename,
		"url":        url,
		"ext":        ext,
		"size":       file.Size,
		"mime":       file.FileHeader.Header.Get("Content-Type"),
		"storage":    storageType,
		"is_image":   isImage,
		"created_at": gtime.Now(),
		"updated_at": gtime.Now(),
	}).Insert()
	if err != nil {
		os.Remove(fullPath)
		return nil, fmt.Errorf("保存文件记录失败: %v", err)
	}

	return &model.UploadOutput{
		ID:      snowflake.JsonInt64(id),
		URL:     url,
		Name:    file.Filename,
		Size:    file.Size,
		Ext:     ext,
		Mime:    file.FileHeader.Header.Get("Content-Type"),
		IsImage: isImage,
	}, nil
}
