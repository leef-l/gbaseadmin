# ============================================
# GBaseAdmin Windows 一键部署脚本
# 功能: 编译后端 + 构建前端 + 打包 + 上传 + 重启
# 用法: .\deploy.ps1 [-Only backend|frontend|wap|all] [-SkipBuild] [-SkipUpload]
# ============================================

param(
    [ValidateSet("all", "backend", "frontend", "wap")]
    [string]$Only = "all",
    [switch]$SkipBuild,
    [switch]$SkipUpload
)

# ---------- 配置区（按需修改） ----------
$SERVER     = "root@pw.easytestdev.online"
$DEPLOY_DIR = "/www/wwwroot/pw.easytestdev.online"
$APPS       = @("system", "play", "upload")
$PORTS      = @("8000", "8001", "8002")

# 项目路径（脚本所在目录）
$PROJECT_DIR  = Split-Path -Parent $MyInvocation.MyCommand.Path
$BACKEND_DIR  = Join-Path $PROJECT_DIR "admin-go"
$FRONTEND_DIR = Join-Path $PROJECT_DIR "vue-vben-admin"
$WAP_DIR      = Join-Path $PROJECT_DIR "wap-ui"
$DIST_DIR     = Join-Path $PROJECT_DIR "dist"

# ---------- 工具函数 ----------
function Info($msg)  { Write-Host "[INFO] $msg" -ForegroundColor Green }
function Warn($msg)  { Write-Host "[WARN] $msg" -ForegroundColor Yellow }
function Fail($msg)  { Write-Host "[FAIL] $msg" -ForegroundColor Red; exit 1 }

function Check-Command($cmd) {
    if (-not (Get-Command $cmd -ErrorAction SilentlyContinue)) {
        Fail "$cmd 未安装，请先安装"
    }
}

# ---------- 前置检查 ----------
Check-Command "go"
Check-Command "ssh"
Check-Command "scp"

# ---------- 清理/创建 dist 目录 ----------
if (Test-Path $DIST_DIR) { Remove-Item $DIST_DIR -Recurse -Force }
New-Item -ItemType Directory -Path $DIST_DIR -Force | Out-Null
New-Item -ItemType Directory -Path (Join-Path $DIST_DIR "backend") -Force | Out-Null

# ============================================
# 1. 编译后端 Go 服务
# ============================================
function Build-Backend {
    Info "===== 编译后端服务 ====="

    $env:CGO_ENABLED = "0"
    $env:GOOS = "linux"
    $env:GOARCH = "amd64"
    $env:GOPROXY = "https://goproxy.cn,direct"

    Push-Location $BACKEND_DIR
    try {
        foreach ($app in $APPS) {
            Info "编译 $app ..."
            $outDir = Join-Path $DIST_DIR "backend\$app"
            New-Item -ItemType Directory -Path $outDir -Force | Out-Null

            # 编译二进制
            go build -ldflags "-s -w" -o "$outDir\$app" ".\app\$app\main.go"
            if ($LASTEXITCODE -ne 0) { Fail "$app 编译失败" }

            # 复制 manifest 配置目录
            $manifestSrc = Join-Path $BACKEND_DIR "app\$app\manifest"
            if (Test-Path $manifestSrc) {
                Copy-Item $manifestSrc -Destination $outDir -Recurse -Force
            }

            Info "$app 编译完成"
        }
    } finally {
        Pop-Location
        # 恢复环境变量
        $env:GOOS = ""
        $env:GOARCH = ""
    }
}

# ============================================
# 2. 构建管理端前端
# ============================================
function Build-Frontend {
    Info "===== 构建管理端前端 ====="
    Check-Command "pnpm"

    Push-Location $FRONTEND_DIR
    try {
        pnpm build:antd
        if ($LASTEXITCODE -ne 0) { Fail "前端构建失败" }

        $distSrc = Join-Path $FRONTEND_DIR "apps\web-antd\dist"
        if (-not (Test-Path $distSrc)) { Fail "前端 dist 目录不存在: $distSrc" }

        Copy-Item $distSrc -Destination (Join-Path $DIST_DIR "frontend") -Recurse -Force
        Info "管理端前端构建完成"
    } finally {
        Pop-Location
    }
}

# ============================================
# 3. 构建 WAP 端 (H5)
# ============================================
function Build-Wap {
    Info "===== 构建 WAP 端 (H5) ====="
    Check-Command "pnpm"

    Push-Location $WAP_DIR
    try {
        pnpm build:h5
        if ($LASTEXITCODE -ne 0) { Fail "WAP 构建失败" }

        $distSrc = Join-Path $WAP_DIR "dist"
        if (-not (Test-Path $distSrc)) { Fail "WAP dist 目录不存在: $distSrc" }

        Copy-Item $distSrc -Destination (Join-Path $DIST_DIR "wap") -Recurse -Force
        Info "WAP 端构建完成"
    } finally {
        Pop-Location
    }
}

# ============================================
# 4. 打包成 tar.gz
# ============================================
function Pack-All {
    Info "===== 打包部署文件 ====="

    $timestamp = Get-Date -Format "yyyyMMdd_HHmm"
    $tarName = "gbaseadmin_$timestamp.tar.gz"
    $script:TarFile = Join-Path $PROJECT_DIR $tarName

    Push-Location $DIST_DIR
    try {
        tar -czf $script:TarFile *
        if ($LASTEXITCODE -ne 0) { Fail "打包失败" }

        $sizeMB = [math]::Round((Get-Item $script:TarFile).Length / 1MB, 2)
        Info "打包完成: $tarName ($sizeMB MB)"
    } finally {
        Pop-Location
    }
}

# ============================================
# 5. 上传到服务器并部署
# ============================================
function Deploy-ToServer {
    Info "===== 上传到服务器 ====="

    # 上传压缩包
    scp $script:TarFile "${SERVER}:/tmp/gbaseadmin_deploy.tar.gz"
    if ($LASTEXITCODE -ne 0) { Fail "上传失败" }
    Info "上传完成"

    # 远程解压并部署
    $remoteScript = @"
set -e

echo '[INFO] 解压部署文件...'
cd /tmp
rm -rf gbaseadmin_deploy && mkdir gbaseadmin_deploy
tar -xzf gbaseadmin_deploy.tar.gz -C gbaseadmin_deploy

# 部署后端
if [ -d /tmp/gbaseadmin_deploy/backend ]; then
    echo '[INFO] 部署后端服务...'
    for app in system play upload; do
        if [ -f /tmp/gbaseadmin_deploy/backend/\$app/\$app ]; then
            echo "[INFO] 停止 gba-\$app ..."
            systemctl stop gba-\$app 2>/dev/null || true

            cp /tmp/gbaseadmin_deploy/backend/\$app/\$app $DEPLOY_DIR/\$app/\$app
            chmod +x $DEPLOY_DIR/\$app/\$app

            # manifest 只在首次部署时复制（不覆盖已有配置）
            # 如需更新配置，手动操作或加 -force 参数
            if [ ! -f $DEPLOY_DIR/\$app/manifest/config/config.yaml ]; then
                cp -r /tmp/gbaseadmin_deploy/backend/\$app/manifest $DEPLOY_DIR/\$app/
            fi

            systemctl start gba-\$app
            echo "[INFO] gba-\$app 已重启"
        fi
    done
fi

# 部署管理端前端
if [ -d /tmp/gbaseadmin_deploy/frontend ]; then
    echo '[INFO] 部署管理端前端...'
    rm -rf $DEPLOY_DIR/admin/*
    mkdir -p $DEPLOY_DIR/admin
    cp -rf /tmp/gbaseadmin_deploy/frontend/* $DEPLOY_DIR/admin/
    echo '[INFO] 管理端前端部署完成'
fi

# 部署 WAP 端
if [ -d /tmp/gbaseadmin_deploy/wap ]; then
    echo '[INFO] 部署 WAP 端...'
    rm -rf $DEPLOY_DIR/wap/*
    mkdir -p $DEPLOY_DIR/wap
    cp -rf /tmp/gbaseadmin_deploy/wap/* $DEPLOY_DIR/wap/
    echo '[INFO] WAP 端部署完成'
fi

# 清理
rm -rf /tmp/gbaseadmin_deploy /tmp/gbaseadmin_deploy.tar.gz

echo '[INFO] ========================================='
echo '[INFO] 部署完成！'
echo '[INFO] ========================================='

# 检查服务状态
for app in system play upload; do
    status=\$(systemctl is-active gba-\$app 2>/dev/null || echo "unknown")
    echo "[INFO] gba-\$app: \$status"
done
"@

    ssh $SERVER $remoteScript
    if ($LASTEXITCODE -ne 0) { Fail "远程部署失败" }
}

# ============================================
# 主流程
# ============================================
$startTime = Get-Date
Info "GBaseAdmin 一键部署开始"
Info "部署目标: $Only"

if (-not $SkipBuild) {
    switch ($Only) {
        "backend"  { Build-Backend }
        "frontend" { Build-Frontend }
        "wap"      { Build-Wap }
        "all"      { Build-Backend; Build-Frontend; Build-Wap }
    }
} else {
    Warn "跳过构建步骤"
}

Pack-All

if (-not $SkipUpload) {
    Deploy-ToServer
} else {
    Warn "跳过上传，打包文件: $($script:TarFile)"
}

# 清理本地 dist 和 tar
Remove-Item $DIST_DIR -Recurse -Force -ErrorAction SilentlyContinue
Remove-Item $script:TarFile -Force -ErrorAction SilentlyContinue

$elapsed = [math]::Round(((Get-Date) - $startTime).TotalSeconds)
Info "全部完成！耗时 ${elapsed} 秒"
