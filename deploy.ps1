# ============================================
# GBaseAdmin Windows Deploy Script
# Usage: .\deploy.ps1 [-Only backend|frontend|wap|all] [-SkipBuild] [-SkipUpload]
# ============================================

param(
    [ValidateSet("all", "backend", "frontend", "wap")]
    [string]$Only = "all",
    [switch]$SkipBuild,
    [switch]$SkipUpload
)

# ---------- Config ----------
$SERVER     = "root@pw.easytestdev.online"
$DEPLOY_DIR = "/www/wwwroot/pw.easytestdev.online"
$APPS       = @("system", "play", "upload")
$PORTS      = @("8000", "8001", "8002")

# Project paths
$PROJECT_DIR  = Split-Path -Parent $MyInvocation.MyCommand.Path
$BACKEND_DIR  = Join-Path $PROJECT_DIR "admin-go"
$FRONTEND_DIR = Join-Path $PROJECT_DIR "vue-vben-admin"
$WAP_DIR      = Join-Path $PROJECT_DIR "wap-ui"
$DIST_DIR     = Join-Path $PROJECT_DIR "dist"

# ---------- Log config ----------
$LOG_DIR  = Join-Path $PROJECT_DIR "deploy-logs"
if (-not (Test-Path $LOG_DIR)) { New-Item -ItemType Directory -Path $LOG_DIR -Force | Out-Null }
$logTimestamp = Get-Date -Format "yyyyMMdd_HHmmss"
$LOG_FILE = Join-Path $LOG_DIR "deploy_${logTimestamp}.log"

# ---------- Helper functions ----------
function Log($msg) {
    $ts = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
    "$ts $msg" | Out-File -Append -FilePath $LOG_FILE -Encoding utf8
}
function Info($msg)  { Write-Host "[INFO] $msg" -ForegroundColor Green;  Log "[INFO] $msg" }
function Warn($msg)  { Write-Host "[WARN] $msg" -ForegroundColor Yellow; Log "[WARN] $msg" }
function Fail($msg)  { Write-Host "[FAIL] $msg" -ForegroundColor Red;    Log "[FAIL] $msg"; exit 1 }

function Check-Command($cmd) {
    if (-not (Get-Command $cmd -ErrorAction SilentlyContinue)) {
        Fail "$cmd not found, please install it first"
    }
}

# ---------- Pre-checks ----------
Check-Command "go"
Check-Command "ssh"
Check-Command "scp"

# ---------- Clean/create dist dir ----------
if (Test-Path $DIST_DIR) { Remove-Item $DIST_DIR -Recurse -Force }
New-Item -ItemType Directory -Path $DIST_DIR -Force | Out-Null
New-Item -ItemType Directory -Path (Join-Path $DIST_DIR "backend") -Force | Out-Null

# ============================================
# 1. Build backend Go services (cross-compile for Linux)
# ============================================
function Build-Backend {
    Info "===== Building backend services ====="

    $env:CGO_ENABLED = "0"
    $env:GOOS = "linux"
    $env:GOARCH = "amd64"
    $env:GOPROXY = "https://goproxy.cn,direct"

    Push-Location $BACKEND_DIR
    try {
        foreach ($app in $APPS) {
            Info "Compiling $app ..."
            $outDir = Join-Path $DIST_DIR "backend\$app"
            New-Item -ItemType Directory -Path $outDir -Force | Out-Null

            $buildOutput = go build -ldflags "-s -w" -o "$outDir\$app" ".\app\$app\main.go" 2>&1
            if ($buildOutput) { Log $buildOutput }
            if ($LASTEXITCODE -ne 0) { Fail "$app build failed: $buildOutput" }

            $manifestSrc = Join-Path $BACKEND_DIR "app\$app\manifest"
            if (Test-Path $manifestSrc) {
                Copy-Item $manifestSrc -Destination $outDir -Recurse -Force
            }

            Info "$app build OK"
        }
    } finally {
        Pop-Location
        $env:GOOS = ""
        $env:GOARCH = ""
    }
}

# ============================================
# 2. Build admin frontend (Vben Admin + Ant Design Vue)
# ============================================
function Build-Frontend {
    Info "===== Building admin frontend ====="
    Check-Command "pnpm"

    Push-Location $FRONTEND_DIR
    try {
        pnpm build:antd
        if ($LASTEXITCODE -ne 0) { Fail "Frontend build failed" }

        $distSrc = Join-Path $FRONTEND_DIR "apps\web-antd\dist"
        if (-not (Test-Path $distSrc)) { Fail "Frontend dist not found: $distSrc" }

        Copy-Item $distSrc -Destination (Join-Path $DIST_DIR "frontend") -Recurse -Force
        Info "Admin frontend build OK"
    } finally {
        Pop-Location
    }
}

# ============================================
# 3. Build WAP (Taro H5)
# ============================================
function Build-Wap {
    Info "===== Building WAP (H5) ====="
    Check-Command "pnpm"

    Push-Location $WAP_DIR
    try {
        pnpm build:h5
        if ($LASTEXITCODE -ne 0) { Fail "WAP build failed" }

        $distSrc = Join-Path $WAP_DIR "dist"
        if (-not (Test-Path $distSrc)) { Fail "WAP dist not found: $distSrc" }

        Copy-Item $distSrc -Destination (Join-Path $DIST_DIR "wap") -Recurse -Force
        Info "WAP build OK"
    } finally {
        Pop-Location
    }
}

# ============================================
# 4. Pack all into tar.gz
# ============================================
function Pack-All {
    Info "===== Packing deploy files ====="

    $packTimestamp = Get-Date -Format "yyyyMMdd_HHmm"
    $tarName = "gbaseadmin_$packTimestamp.tar.gz"
    $script:TarFile = Join-Path $PROJECT_DIR $tarName

    Push-Location $DIST_DIR
    try {
        tar -czf $script:TarFile *
        if ($LASTEXITCODE -ne 0) { Fail "Pack failed" }

        $sizeMB = [math]::Round((Get-Item $script:TarFile).Length / 1MB, 2)
        Info "Packed: $tarName ($sizeMB MB)"
    } finally {
        Pop-Location
    }
}

# ============================================
# 5. Upload to server and deploy via SSH
# ============================================
function Deploy-ToServer {
    Info "===== Uploading to server ====="

    scp $script:TarFile "${SERVER}:/tmp/gbaseadmin_deploy.tar.gz"
    if ($LASTEXITCODE -ne 0) { Fail "Upload failed" }
    Info "Upload OK"

    $remoteShell = Join-Path $env:TEMP "gba_remote_deploy.sh"
    $shellContent = @'
#!/bin/bash
set -e

DEPLOY_DIR="__DEPLOY_DIR__"

echo '[INFO] Extracting deploy files...'
cd /tmp
rm -rf gbaseadmin_deploy && mkdir gbaseadmin_deploy
tar -xzf gbaseadmin_deploy.tar.gz -C gbaseadmin_deploy

# ---- Deploy backend ----
if [ -d /tmp/gbaseadmin_deploy/backend ]; then
    echo '[INFO] Deploying backend services...'
    for app in system play upload; do
        if [ -f /tmp/gbaseadmin_deploy/backend/$app/$app ]; then
            echo "[INFO] Stopping gba-$app ..."
            systemctl stop gba-$app 2>/dev/null || true
            sleep 2

            cp /tmp/gbaseadmin_deploy/backend/$app/$app $DEPLOY_DIR/$app/$app
            chmod +x $DEPLOY_DIR/$app/$app

            if [ ! -f $DEPLOY_DIR/$app/manifest/config/config.yaml ]; then
                cp -r /tmp/gbaseadmin_deploy/backend/$app/manifest $DEPLOY_DIR/$app/
            fi

            systemctl start gba-$app
            echo "[INFO] gba-$app restarted"
        fi
    done
fi

# ---- Deploy admin frontend ----
if [ -d /tmp/gbaseadmin_deploy/frontend ]; then
    echo '[INFO] Deploying admin frontend...'
    rm -rf $DEPLOY_DIR/admin/*
    mkdir -p $DEPLOY_DIR/admin
    cp -rf /tmp/gbaseadmin_deploy/frontend/* $DEPLOY_DIR/admin/
    echo '[INFO] Admin frontend deployed'
fi

# ---- Deploy WAP ----
if [ -d /tmp/gbaseadmin_deploy/wap ]; then
    echo '[INFO] Deploying WAP...'
    rm -rf $DEPLOY_DIR/wap/*
    mkdir -p $DEPLOY_DIR/wap
    cp -rf /tmp/gbaseadmin_deploy/wap/* $DEPLOY_DIR/wap/
    echo '[INFO] WAP deployed'
fi

# ---- Cleanup ----
rm -rf /tmp/gbaseadmin_deploy /tmp/gbaseadmin_deploy.tar.gz

echo '[INFO] ========================================='
echo '[INFO] Deploy completed!'
echo '[INFO] ========================================='

# ---- Service status ----
for app in system play upload; do
    status=$(systemctl is-active gba-$app 2>/dev/null || echo "unknown")
    echo "[INFO] gba-$app: $status"
done
'@
    $shellContent = $shellContent.Replace("__DEPLOY_DIR__", $DEPLOY_DIR)
    # Write with LF line endings and NO BOM (Linux bash chokes on UTF-8 BOM)
    $utf8NoBom = New-Object System.Text.UTF8Encoding($false)
    $shellContent = $shellContent.Replace("`r`n", "`n")
    [System.IO.File]::WriteAllText($remoteShell, $shellContent, $utf8NoBom)

    scp $remoteShell "${SERVER}:/tmp/gba_deploy.sh"
    if ($LASTEXITCODE -ne 0) { Fail "Upload deploy script failed" }

    $remoteOutput = ssh $SERVER "bash /tmp/gba_deploy.sh && rm -f /tmp/gba_deploy.sh" 2>&1
    $remoteOutput | ForEach-Object {
        Write-Host $_
        Log $_
    }
    Remove-Item $remoteShell -Force -ErrorAction SilentlyContinue
    if ($LASTEXITCODE -ne 0) { Fail "Remote deploy failed" }
}

# ============================================
# Main flow
# ============================================
$startTime = Get-Date
Info "GBaseAdmin deploy started"
Info "Target: $Only"
Info "Log file: $LOG_FILE"

# Step 1: Build
if (-not $SkipBuild) {
    switch ($Only) {
        "backend"  { Build-Backend }
        "frontend" { Build-Frontend }
        "wap"      { Build-Wap }
        "all"      { Build-Backend; Build-Frontend; Build-Wap }
    }
} else {
    Warn "Skipping build"
}

# Step 2: Pack
Pack-All

# Step 3: Upload & Deploy
if (-not $SkipUpload) {
    Deploy-ToServer
} else {
    Warn "Skipping upload, packed file: $($script:TarFile)"
}

# Step 4: Cleanup
Remove-Item $DIST_DIR -Recurse -Force -ErrorAction SilentlyContinue
Remove-Item $script:TarFile -Force -ErrorAction SilentlyContinue

# Done
$elapsed = [math]::Round(((Get-Date) - $startTime).TotalSeconds)
Info "All done! Elapsed: ${elapsed}s"
Info "Deploy log saved: $LOG_FILE"

# Auto-cleanup logs older than 30 days
Get-ChildItem $LOG_DIR -Filter "deploy_*.log" |
    Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-30) } |
    Remove-Item -Force
