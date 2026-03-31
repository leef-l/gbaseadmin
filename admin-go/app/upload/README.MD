# GBaseAdmin - Upload 应用

文件管理应用，提供文件上传、目录管理、上传配置等功能。

## 数据库表

| 表名 | 说明 |
|------|------|
| `upload_dir` | 文件目录（树形结构） |
| `upload_file` | 文件记录 |
| `upload_config` | 上传配置（本地/OSS/COS） |
| `upload_dir_rule` | 文件目录规则 |

## API 路由

所有接口前缀：`/api/upload/`，需要 Auth 认证。

| 模块 | 路径 | 说明 |
|------|------|------|
| 文件上传 | `POST /api/upload/uploader/upload` | 上传文件（multipart/form-data） |
| 文件目录 | `/api/upload/dir/list,create,update,delete,detail,tree` | 目录管理（树形） |
| 文件记录 | `/api/upload/file/list,create,update,delete,detail` | 文件管理 |
| 上传配置 | `/api/upload/config/list,create,update,delete,detail` | 配置管理 |
| 目录规则 | `/api/upload/dir_rule/list,create,update,delete,detail` | 目录规则管理 |

静态文件访问：`/upload/` 映射到 `resource/upload/` 目录。

## 上传流程

1. 前端调用 `POST /api/upload/uploader/upload`，发送 `file`（文件）+ `dirId`（可选目录ID）
2. 后端读取 `upload_config` 中 `is_default=1` 的默认配置，获取存储类型和大小限制
3. 验证文件大小，解析扩展名和 MIME 类型，判断是否为图片
4. 本地存储：保存到 `resource/upload/{日期}/{唯一文件名}`
5. 写入 `upload_file` 记录，返回文件信息（id, url, name, size, ext, mime, isImage）

## 存储类型

| 值 | 类型 | 配置字段 |
|----|------|---------|
| 1 | 本地存储 | `local_path` |
| 2 | 阿里云 OSS | `oss_endpoint`, `oss_bucket`, `oss_access_key`, `oss_secret_key` |
| 3 | 腾讯云 COS | `cos_region`, `cos_bucket`, `cos_secret_id`, `cos_secret_key` |

## 文件删除

删除文件时同时执行物理删除，根据 `storage` 字段判断存储类型：
- 本地：直接删除磁盘文件
- OSS/COS：TODO（待实现云端删除）

## 前端组件

本应用提供以下可复用组件：

| 组件 | 路径 | 说明 |
|------|------|------|
| ImageUpload | `components/upload/image-upload.vue` | 图片上传（picture-card 样式） |
| FileUpload | `components/upload/file-upload.vue` | 文件上传 |
| RichText | `components/tinymce/index.vue` | TinyMCE v8 富文本编辑器 |
| JsonEditor | `components/json-editor/index.vue` | JSON 编辑器（tree + code 双模式） |
| ImagePicker | `components/image-picker/` | 图片选择器（目录树 + 图片网格 + 勾选） |

这些组件已注册到 `adapter/component/index.ts`，可在 vben-form 中直接使用：

```typescript
{ component: 'ImageUpload', fieldName: 'avatar', label: '头像', componentProps: { maxCount: 1 } }
{ component: 'FileUpload', fieldName: 'attachment', label: '附件', componentProps: { maxCount: 3 } }
{ component: 'RichText', fieldName: 'content', label: '内容', formItemClass: 'col-span-full' }
{ component: 'JsonEditor', fieldName: 'config', label: '配置', formItemClass: 'col-span-full' }
```

## 前端页面

| 页面 | 说明 |
|------|------|
| 文件列表 | 左侧目录树 + 列表/区块双视图切换 + 图片预览 + 上传 |
| 上传配置 | storage RadioGroup 联动显示本地/OSS/COS 配置字段 |
| 文件目录 | 树形 CRUD |
| 目录规则 | 标准 CRUD |
