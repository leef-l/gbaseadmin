# Windows 一键部署使用文档

## 概述

`deploy.ps1` 是 GBaseAdmin 项目的 Windows 一键部署脚本，在本地完成编译打包，通过 SSH 上传到服务器并自动重启服务，避免在配置较低的服务器上执行编译。

## 部署流程

```
Windows 本地                                服务器 (pw.easytestdev.online)
┌─────────────────────┐                    ┌─────────────────────────────┐
│ 1. Go 交叉编译       │                    │                             │
│    system/play/upload│                    │                             │
│ 2. pnpm build:antd  │    SCP 上传        │ 4. 解压到部署目录             │
│    管理端前端构建      │ ──────────────►   │ 5. 替换二进制文件             │
│ 3. pnpm build:h5    │    tar.gz          │ 6. systemctl restart        │
│    WAP端H5构建       │                    │ 7. 检查服务状态               │
└─────────────────────┘                    └─────────────────────────────┘
```

## 前置准备（只需做一次）

### 1. 安装必备工具

| 工具 | 用途 | 安装方式 |
|------|------|----------|
| Go 1.23+ | 编译后端 | https://go.dev/dl/ |
| pnpm | 构建前端/WAP | `npm install -g pnpm` |
| SSH | 上传和远程执行 | Win11 自带 OpenSSH |

### 2. 配置 SSH 免密登录

```powershell
# 生成密钥（如果没有）
ssh-keygen -t ed25519

# 上传公钥到服务器（输入一次密码）
type $env:USERPROFILE\.ssh\id_ed25519.pub | ssh root@pw.easytestdev.online "mkdir -p ~/.ssh && cat >> ~/.ssh/authorized_keys && chmod 700 ~/.ssh && chmod 600 ~/.ssh/authorized_keys"

# 验证免密登录
ssh root@pw.easytestdev.online "echo OK"
```

## 使用方法

在项目根目录打开 PowerShell：

```powershell
# 全量部署（后端 + 管理端 + WAP）
.\deploy.ps1

# 只部署后端（改了 Go 代码后最常用）
.\deploy.ps1 -Only backend

# 只部署管理端前端
.\deploy.ps1 -Only frontend

# 只部署 WAP 端
.\deploy.ps1 -Only wap

# 只打包不上传（本地验证构建是否成功）
.\deploy.ps1 -SkipUpload

# 跳过构建直接上传（已有 dist 目录时）
.\deploy.ps1 -SkipBuild
```

## 参数说明

| 参数 | 可选值 | 默认值 | 说明 |
|------|--------|--------|------|
| `-Only` | `all` / `backend` / `frontend` / `wap` | `all` | 选择部署哪个部分 |
| `-SkipBuild` | 开关 | 否 | 跳过编译构建步骤 |
| `-SkipUpload` | 开关 | 否 | 跳过上传部署步骤 |

## 服务器目录结构

```
/www/wwwroot/pw.easytestdev.online/
├── system/              # system 服务
│   ├── system           # 二进制文件
│   └── manifest/config/ # GoFrame 配置
├── play/                # play 服务
│   ├── play             # 二进制文件
│   └── manifest/config/ # GoFrame 配置
├── upload/              # upload 服务
│   ├── upload           # 二进制文件
│   └── manifest/config/ # GoFrame 配置
├── admin/               # 管理端前端静态文件
└── wap/                 # WAP 端 H5 静态文件
```

## 服务管理命令（SSH 到服务器后）

```bash
# 查看服务状态
systemctl status gba-system
systemctl status gba-play
systemctl status gba-upload

# 手动重启某个服务
systemctl restart gba-play

# 查看日志
journalctl -u gba-play -f
```

## 常见问题

### Q: PowerShell 提示"无法加载文件，因为在此系统上禁止运行脚本"
```powershell
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
```

### Q: SSH 连接超时
检查服务器防火墙是否开放 22 端口，或联系服务商确认 SSH 端口。

### Q: Go 编译报错
确保在项目根目录执行，且 `admin-go/go.mod` 存在。可先手动测试：
```powershell
cd admin-go
$env:CGO_ENABLED="0"; $env:GOOS="linux"; $env:GOARCH="amd64"
go build ./app/system/...
```

### Q: 服务器上配置文件被覆盖了
脚本只在首次部署时（config.yaml 不存在）才复制 manifest，不会覆盖已有配置。如需更新配置请手动 SSH 修改。
