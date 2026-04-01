# ============================================
# GBaseAdmin Windows 一键部署脚本
# 功能: 编译后端 + 构建前端 + 打包 + 上传 + 重启
# 用法: .\deploy.ps1 [-Only backend|frontend|wap|all] [-SkipBuild] [-SkipUpload]
# ============================================

# 定义脚本参数：-Only 选择部署内容，-SkipBuild 跳过编译，-SkipUpload 跳过上传
param(
    [ValidateSet("all", "backend", "frontend", "wap")]  # 限定 -Only 参数只能是这四个值
    [string]$Only = "all",                               # 默认全量部署
    [switch]$SkipBuild,                                  # 开关参数：是否跳过构建
    [switch]$SkipUpload                                  # 开关参数：是否跳过上传
)

# ---------- 配置区（按需修改） ----------
$SERVER     = "root@pw.easytestdev.online"               # SSH 连接地址（用户名@服务器域名）
$DEPLOY_DIR = "/www/wwwroot/pw.easytestdev.online"       # 服务器上的部署根目录
$APPS       = @("system", "play", "upload")              # 三个后端服务名称
$PORTS      = @("8000", "8001", "8002")                  # 对应的端口号（备用）

# 项目路径（根据脚本所在位置自动计算）
$PROJECT_DIR  = Split-Path -Parent $MyInvocation.MyCommand.Path  # 脚本所在目录 = 项目根目录
$BACKEND_DIR  = Join-Path $PROJECT_DIR "admin-go"                # 后端 Go 代码目录
$FRONTEND_DIR = Join-Path $PROJECT_DIR "vue-vben-admin"          # 管理端前端目录
$WAP_DIR      = Join-Path $PROJECT_DIR "wap-ui"                  # WAP 端目录
$DIST_DIR     = Join-Path $PROJECT_DIR "dist"                    # 本地临时构建产物目录

# ---------- 工具函数 ----------
function Info($msg)  { Write-Host "[INFO] $msg" -ForegroundColor Green }   # 绿色信息输出
function Warn($msg)  { Write-Host "[WARN] $msg" -ForegroundColor Yellow }  # 黄色警告输出
function Fail($msg)  { Write-Host "[FAIL] $msg" -ForegroundColor Red; exit 1 }  # 红色错误输出并退出

# 检查命令是否存在，不存在则报错退出
function Check-Command($cmd) {
    if (-not (Get-Command $cmd -ErrorAction SilentlyContinue)) {
        Fail "$cmd 未安装，请先安装"
    }
}

# ---------- 前置检查：确保必备工具已安装 ----------
Check-Command "go"    # Go 编译器
Check-Command "ssh"   # SSH 客户端（Win11 自带）
Check-Command "scp"   # SCP 文件传输（Win11 自带）

# ---------- 清理/创建本地临时目录 ----------
if (Test-Path $DIST_DIR) { Remove-Item $DIST_DIR -Recurse -Force }         # 删除旧的 dist 目录
New-Item -ItemType Directory -Path $DIST_DIR -Force | Out-Null             # 创建新的 dist 目录
New-Item -ItemType Directory -Path (Join-Path $DIST_DIR "backend") -Force | Out-Null  # 创建 dist/backend 子目录

# ============================================
# 1. 编译后端 Go 服务（交叉编译为 Linux 二进制）
# ============================================
function Build-Backend {
    Info "===== 编译后端服务 ====="

    # 设置 Go 交叉编译环境变量
    $env:CGO_ENABLED = "0"                      # 禁用 CGO（纯 Go 编译，无需 C 编译器）
    $env:GOOS = "linux"                         # 目标操作系统：Linux
    $env:GOARCH = "amd64"                       # 目标架构：x86_64
    $env:GOPROXY = "https://goproxy.cn,direct"  # Go 模块代理（加速国内下载）

    Push-Location $BACKEND_DIR                  # 切换到后端目录
    try {
        foreach ($app in $APPS) {               # 遍历 system、play、upload 三个服务
            Info "编译 $app ..."
            $outDir = Join-Path $DIST_DIR "backend\$app"                   # 输出目录：dist/backend/{app}
            New-Item -ItemType Directory -Path $outDir -Force | Out-Null   # 创建输出目录

            # 编译二进制：-ldflags "-s -w" 去除调试信息和符号表，减小体积
            go build -ldflags "-s -w" -o "$outDir\$app" ".\app\$app\main.go"
            if ($LASTEXITCODE -ne 0) { Fail "$app 编译失败" }  # 编译失败则退出

            # 复制 manifest 配置目录（包含 config.yaml 模板）
            $manifestSrc = Join-Path $BACKEND_DIR "app\$app\manifest"
            if (Test-Path $manifestSrc) {
                Copy-Item $manifestSrc -Destination $outDir -Recurse -Force
            }

            Info "$app 编译完成"
        }
    } finally {
        Pop-Location            # 恢复原工作目录
        $env:GOOS = ""          # 清除交叉编译环境变量，避免影响本地 Go 开发
        $env:GOARCH = ""
    }
}

# ============================================
# 2. 构建管理端前端（Vben Admin + Ant Design Vue）
# ============================================
function Build-Frontend {
    Info "===== 构建管理端前端 ====="
    Check-Command "pnpm"                        # 检查 pnpm 是否安装

    Push-Location $FRONTEND_DIR                 # 切换到前端目录
    try {
        pnpm build:antd                         # 执行 Vben Admin 的 Ant Design 版本构建
        if ($LASTEXITCODE -ne 0) { Fail "前端构建失败" }

        # 构建产物在 apps/web-antd/dist 目录下
        $distSrc = Join-Path $FRONTEND_DIR "apps\web-antd\dist"
        if (-not (Test-Path $distSrc)) { Fail "前端 dist 目录不存在: $distSrc" }

        # 复制到统一的 dist/frontend 目录
        Copy-Item $distSrc -Destination (Join-Path $DIST_DIR "frontend") -Recurse -Force
        Info "管理端前端构建完成"
    } finally {
        Pop-Location
    }
}

# ============================================
# 3. 构建 WAP 端（Taro H5 模式）
# ============================================
function Build-Wap {
    Info "===== 构建 WAP 端 (H5) ====="
    Check-Command "pnpm"                        # 检查 pnpm 是否安装

    Push-Location $WAP_DIR                      # 切换到 WAP 端目录
    try {
        pnpm build:h5                           # 执行 Taro H5 构建
        if ($LASTEXITCODE -ne 0) { Fail "WAP 构建失败" }

        $distSrc = Join-Path $WAP_DIR "dist"    # Taro 构建产物目录
        if (-not (Test-Path $distSrc)) { Fail "WAP dist 目录不存在: $distSrc" }

        # 复制到统一的 dist/wap 目录
        Copy-Item $distSrc -Destination (Join-Path $DIST_DIR "wap") -Recurse -Force
        Info "WAP 端构建完成"
    } finally {
        Pop-Location
    }
}

# ============================================
# 4. 将所有构建产物打包成 tar.gz 压缩包
# ============================================
function Pack-All {
    Info "===== 打包部署文件 ====="

    $timestamp = Get-Date -Format "yyyyMMdd_HHmm"          # 时间戳，如 20260401_2130
    $tarName = "gbaseadmin_$timestamp.tar.gz"               # 压缩包文件名
    $script:TarFile = Join-Path $PROJECT_DIR $tarName       # $script: 作用域使变量在函数外可访问

    Push-Location $DIST_DIR
    try {
        tar -czf $script:TarFile *                          # 用 tar 压缩 dist 目录下所有内容
        if ($LASTEXITCODE -ne 0) { Fail "打包失败" }

        # 显示压缩包大小
        $sizeMB = [math]::Round((Get-Item $script:TarFile).Length / 1MB, 2)
        Info "打包完成: $tarName ($sizeMB MB)"
    } finally {
        Pop-Location
    }
}

# ============================================
# 5. 通过 SSH 上传到服务器并执行远程部署
# ============================================
function Deploy-ToServer {
    Info "===== 上传到服务器 ====="

    # SCP 上传压缩包到服务器 /tmp 目录
    scp $script:TarFile "${SERVER}:/tmp/gbaseadmin_deploy.tar.gz"
    if ($LASTEXITCODE -ne 0) { Fail "上传失败" }
    Info "上传完成"

    # 生成远程部署脚本（写入临时文件避免 PowerShell 变量转义问题）
    $remoteShell = Join-Path $env:TEMP "gba_remote_deploy.sh"       # 本地临时 shell 脚本路径
    # 用单引号 here-string，PowerShell 不会解析其中的 $ 符号
    $shellContent = @'
#!/bin/bash
set -e

DEPLOY_DIR="__DEPLOY_DIR__"

echo '[INFO] 解压部署文件...'
cd /tmp
rm -rf gbaseadmin_deploy && mkdir gbaseadmin_deploy
tar -xzf gbaseadmin_deploy.tar.gz -C gbaseadmin_deploy

# ---- 部署后端 ----
if [ -d /tmp/gbaseadmin_deploy/backend ]; then
    echo '[INFO] 部署后端服务...'
    for app in system play upload; do
        if [ -f /tmp/gbaseadmin_deploy/backend/$app/$app ]; then
            echo "[INFO] 停止 gba-$app ..."
            systemctl stop gba-$app 2>/dev/null || true

            cp /tmp/gbaseadmin_deploy/backend/$app/$app $DEPLOY_DIR/$app/$app
            chmod +x $DEPLOY_DIR/$app/$app

            if [ ! -f $DEPLOY_DIR/$app/manifest/config/config.yaml ]; then
                cp -r /tmp/gbaseadmin_deploy/backend/$app/manifest $DEPLOY_DIR/$app/
            fi

            systemctl start gba-$app
            echo "[INFO] gba-$app 已重启"
        fi
    done
fi

# ---- 部署管理端前端 ----
if [ -d /tmp/gbaseadmin_deploy/frontend ]; then
    echo '[INFO] 部署管理端前端...'
    rm -rf $DEPLOY_DIR/admin/*
    mkdir -p $DEPLOY_DIR/admin
    cp -rf /tmp/gbaseadmin_deploy/frontend/* $DEPLOY_DIR/admin/
    echo '[INFO] 管理端前端部署完成'
fi

# ---- 部署 WAP 端 ----
if [ -d /tmp/gbaseadmin_deploy/wap ]; then
    echo '[INFO] 部署 WAP 端...'
    rm -rf $DEPLOY_DIR/wap/*
    mkdir -p $DEPLOY_DIR/wap
    cp -rf /tmp/gbaseadmin_deploy/wap/* $DEPLOY_DIR/wap/
    echo '[INFO] WAP 端部署完成'
fi

# ---- 清理临时文件 ----
rm -rf /tmp/gbaseadmin_deploy /tmp/gbaseadmin_deploy.tar.gz

echo '[INFO] ========================================='
echo '[INFO] 部署完成！'
echo '[INFO] ========================================='

for app in system play upload; do
    status=$(systemctl is-active gba-$app 2>/dev/null || echo "unknown")
    echo "[INFO] gba-$app: $status"
done
'@
    # 替换占位符为实际部署目录
    $shellContent = $shellContent.Replace("__DEPLOY_DIR__", $DEPLOY_DIR)
    # 写入临时文件（确保 Linux 换行符 LF）
    $shellContent | Set-Content -Path $remoteShell -Encoding utf8 -NoNewline
    (Get-Content $remoteShell -Raw).Replace("`r`n", "`n") | Set-Content -Path $remoteShell -Encoding utf8 -NoNewline

    # 上传脚本到服务器并执行
    scp $remoteShell "${SERVER}:/tmp/gba_deploy.sh"                 # 上传部署脚本
    if ($LASTEXITCODE -ne 0) { Fail "上传部署脚本失败" }
    ssh $SERVER "bash /tmp/gba_deploy.sh && rm -f /tmp/gba_deploy.sh"  # 执行后删除脚本
    Remove-Item $remoteShell -Force -ErrorAction SilentlyContinue   # 删除本地临时脚本
    if ($LASTEXITCODE -ne 0) { Fail "远程部署失败" }
}

# ============================================
# 主流程：按参数依次执行 构建 → 打包 → 上传部署
# ============================================
$startTime = Get-Date                           # 记录开始时间
Info "GBaseAdmin 一键部署开始"
Info "部署目标: $Only"

# 步骤一：构建（除非 -SkipBuild）
if (-not $SkipBuild) {
    switch ($Only) {
        "backend"  { Build-Backend }                                    # 只编译后端
        "frontend" { Build-Frontend }                                   # 只构建管理端前端
        "wap"      { Build-Wap }                                        # 只构建 WAP 端
        "all"      { Build-Backend; Build-Frontend; Build-Wap }         # 全量构建
    }
} else {
    Warn "跳过构建步骤"
}

# 步骤二：打包
Pack-All

# 步骤三：上传部署（除非 -SkipUpload）
if (-not $SkipUpload) {
    Deploy-ToServer
} else {
    Warn "跳过上传，打包文件: $($script:TarFile)"
}

# 步骤四：清理本地临时文件
Remove-Item $DIST_DIR -Recurse -Force -ErrorAction SilentlyContinue     # 删除 dist 目录
Remove-Item $script:TarFile -Force -ErrorAction SilentlyContinue         # 删除 tar.gz 文件

# 输出总耗时
$elapsed = [math]::Round(((Get-Date) - $startTime).TotalSeconds)
Info "全部完成！耗时 ${elapsed} 秒"
