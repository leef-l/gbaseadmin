# ============================================
# GBaseAdmin Windows Deploy Script (rsync + sequential)
# Usage: .\deploy.ps1 [-Only backend|frontend|wap|all] [-SkipBuild] [-SkipUpload]
#
# Optimized for low-resource servers (2C4G):
#   - rsync incremental transfer (no tar/unpack on server)
#   - sequential service restart (one at a time, wait for stable)
# ============================================

param(
    [ValidateSet("all", "backend", "frontend", "wap")]
    [string]$Only = "all",
    [switch]$SkipBuild,
    [switch]$SkipUpload
)

# ---------- Config ----------
$SERVER     = "root@pw.easytestdev.online"
$SERVER_HOST = "pw.easytestdev.online"
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
Check-Command "rsync"

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
# 4. Deploy backend: rsync + sequential restart (one by one)
# ============================================
function Deploy-Backend {
    Info "===== Deploying backend services (sequential) ====="

    foreach ($app in $APPS) {
        $localAppDir = Join-Path $DIST_DIR "backend\$app"
        if (-not (Test-Path "$localAppDir\$app")) {
            Warn "$app binary not found, skipping"
            continue
        }

        Info "--- Deploying $app ---"

        # Step 1: rsync binary to staging area on server (avoid overwriting running binary)
        Info "[$app] Uploading via rsync ..."
        ssh $SERVER "mkdir -p /tmp/gba_stage/$app"
        rsync -az --compress-level=1 --progress -e ssh "$localAppDir/" "${SERVER}:/tmp/gba_stage/$app/"
        if ($LASTEXITCODE -ne 0) { Fail "[$app] rsync upload failed" }
        Info "[$app] Upload OK"

        # Step 2: stop service, wait, replace, start (all on server)
        Info "[$app] Stopping, replacing, starting ..."
        $remoteCmd = @"
set -e
echo '[$app] Stopping service...'
systemctl stop gba-$app 2>/dev/null || true
for i in `$(seq 1 10); do
    if ! pgrep -x "$app" >/dev/null 2>&1; then break; fi
    sleep 1
done

echo '[$app] Replacing binary...'
mkdir -p $DEPLOY_DIR/$app
cp /tmp/gba_stage/$app/$app $DEPLOY_DIR/$app/$app
chmod +x $DEPLOY_DIR/$app/$app

if [ ! -f $DEPLOY_DIR/$app/manifest/config/config.yaml ]; then
    echo '[$app] Copying manifest (first deploy)...'
    cp -r /tmp/gba_stage/$app/manifest $DEPLOY_DIR/$app/
fi

echo '[$app] Starting service...'
systemctl start gba-$app

# Wait for service to become active (up to 8s)
for i in `$(seq 1 8); do
    status=`$(systemctl is-active gba-$app 2>/dev/null || echo 'unknown')
    if [ "`$status" = "active" ]; then break; fi
    sleep 1
done

status=`$(systemctl is-active gba-$app 2>/dev/null || echo 'unknown')
echo "[$app] Service status: `$status"

rm -rf /tmp/gba_stage/$app
echo '[$app] Done'
"@
        $output = ssh $SERVER $remoteCmd 2>&1
        $output | ForEach-Object {
            Write-Host "  $_"
            Log "  $_"
        }
        if ($LASTEXITCODE -ne 0) { Fail "[$app] Remote deploy failed" }
        Info "[$app] Deploy completed"

        # Brief pause between services to let server stabilize
        if ($app -ne $APPS[-1]) {
            Info "Waiting 3s before next service ..."
            Start-Sleep -Seconds 3
        }
    }
}

# ============================================
# 5. Deploy frontend: rsync directly to target (no service restart needed)
# ============================================
function Deploy-Frontend {
    $localFrontendDir = Join-Path $DIST_DIR "frontend"
    if (-not (Test-Path $localFrontendDir)) {
        Warn "Frontend dist not found, skipping"
        return
    }

    Info "===== Deploying admin frontend (rsync) ====="
    ssh $SERVER "mkdir -p $DEPLOY_DIR/admin"
    rsync -az --delete --compress-level=1 --progress -e ssh "$localFrontendDir/" "${SERVER}:$DEPLOY_DIR/admin/"
    if ($LASTEXITCODE -ne 0) { Fail "Frontend rsync failed" }
    Info "Admin frontend deployed"
}

# ============================================
# 6. Deploy WAP: rsync directly to target (no service restart needed)
# ============================================
function Deploy-Wap {
    $localWapDir = Join-Path $DIST_DIR "wap"
    if (-not (Test-Path $localWapDir)) {
        Warn "WAP dist not found, skipping"
        return
    }

    Info "===== Deploying WAP (rsync) ====="
    ssh $SERVER "mkdir -p $DEPLOY_DIR/wap"
    rsync -az --delete --compress-level=1 --progress -e ssh "$localWapDir/" "${SERVER}:$DEPLOY_DIR/wap/"
    if ($LASTEXITCODE -ne 0) { Fail "WAP rsync failed" }
    Info "WAP deployed"
}

# ============================================
# Main flow
# ============================================
$startTime = Get-Date
Info "GBaseAdmin deploy started (rsync + sequential mode)"
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

# Step 2: Upload & Deploy (sequential, resource-friendly)
if (-not $SkipUpload) {
    switch ($Only) {
        "backend"  { Deploy-Backend }
        "frontend" { Deploy-Frontend }
        "wap"      { Deploy-Wap }
        "all"      {
            # Deploy backend first (sequential, one service at a time)
            Deploy-Backend
            # Then static files (low resource usage)
            Deploy-Frontend
            Deploy-Wap
        }
    }
} else {
    Warn "Skipping upload"
}

# Step 3: Cleanup local dist
Remove-Item $DIST_DIR -Recurse -Force -ErrorAction SilentlyContinue

# Step 4: Final status check
if (-not $SkipUpload -and ($Only -eq "all" -or $Only -eq "backend")) {
    Info "===== Final service status ====="
    $statusOutput = ssh $SERVER "for app in system play upload; do echo `"gba-`$app: `$(systemctl is-active gba-`$app 2>/dev/null || echo unknown)`"; done" 2>&1
    $statusOutput | ForEach-Object {
        Write-Host "  $_" -ForegroundColor Cyan
        Log "  $_"
    }
}

# Done
$elapsed = [math]::Round(((Get-Date) - $startTime).TotalSeconds)
Info "All done! Elapsed: ${elapsed}s"
Info "Deploy log saved: $LOG_FILE"

# Auto-cleanup logs older than 30 days
Get-ChildItem $LOG_DIR -Filter "deploy_*.log" |
    Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-30) } |
    Remove-Item -Force
