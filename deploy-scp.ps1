# ============================================
# GBaseAdmin Windows Deploy Script (scp + sequential)
# Usage: .\deploy-scp.ps1 [-Only backend|frontend|wap|all] [-SkipBuild] [-SkipUpload]
#
# Zero extra dependencies — only needs ssh/scp (Windows built-in)
# Optimized for low-resource servers (2C4G):
#   - sequential service restart (one at a time, wait for stable)
#   - no large tar packing/unpacking on server
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

# ---------- Encoding ----------
$script:utf8NoBom = New-Object System.Text.UTF8Encoding($false)

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

# Write shell script with LF endings, scp to server, execute, cleanup
function Run-RemoteScript {
    param([string]$Name, [string]$Script)
    $localFile = Join-Path $env:TEMP "gba_${Name}.sh"
    $lfContent = $Script -replace "`r`n", "`n" -replace "`r", "`n"
    [System.IO.File]::WriteAllText($localFile, $lfContent, $script:utf8NoBom)

    scp -q $localFile "${SERVER}:/tmp/gba_${Name}.sh"
    $scpExit = $LASTEXITCODE
    Remove-Item $localFile -Force -ErrorAction SilentlyContinue
    if ($scpExit -ne 0) { Fail "[$Name] Upload script failed" }

    $output = ssh $SERVER "bash /tmp/gba_${Name}.sh; ret=`$?; rm -f /tmp/gba_${Name}.sh; exit `$ret" 2>&1
    $output | ForEach-Object {
        Write-Host "  $_"
        Log "  $_"
    }
    if ($LASTEXITCODE -ne 0) { Fail "[$Name] Remote script failed" }
}

# ---------- Pre-checks ----------
Check-Command "go"
Check-Command "ssh"
Check-Command "scp"

# Verify SSH connectivity before spending time on builds
if (-not $SkipUpload) {
    Info "Checking server connectivity ..."
    ssh -o ConnectTimeout=5 -o BatchMode=yes $SERVER "echo ok" >$null 2>&1
    if ($LASTEXITCODE -ne 0) { Fail "Cannot connect to $SERVER via SSH. Check your SSH key and network." }
    Info "Server connectivity OK"
}

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
# 4. Deploy backend: scp + sequential restart (one by one)
# ============================================
function Deploy-Backend {
    Info "===== Deploying backend services (one by one) ====="

    foreach ($app in $APPS) {
        $localAppDir = Join-Path $DIST_DIR "backend\$app"
        if (-not (Test-Path "$localAppDir\$app")) {
            Warn "$app binary not found, skipping"
            continue
        }

        Info "--- [$app] Start ---"

        # Step 1: scp binary to staging area on server
        Info "[$app] Uploading binary ..."
        ssh $SERVER "mkdir -p /tmp/gba_stage/$app"
        scp -q "$localAppDir\$app" "${SERVER}:/tmp/gba_stage/$app/$app"
        if ($LASTEXITCODE -ne 0) { Fail "[$app] Upload binary failed" }

        # Upload manifest
        $manifestDir = Join-Path $localAppDir "manifest"
        if (Test-Path $manifestDir) {
            scp -q -r "$manifestDir" "${SERVER}:/tmp/gba_stage/$app/"
        }
        Info "[$app] Upload OK"

        # Step 2: stop -> replace -> start (via temp script on server)
        Info "[$app] Stopping, replacing, starting ..."
        $shellBody = @'
#!/bin/bash
set -e
APP="__APP__"
DEPLOY_DIR="__DEPLOY_DIR__"

echo "[$APP] Stopping service..."
systemctl stop gba-$APP 2>/dev/null || true
for i in $(seq 1 10); do
    if ! pgrep -x "$APP" >/dev/null 2>&1; then break; fi
    sleep 1
done

echo "[$APP] Replacing binary..."
mkdir -p $DEPLOY_DIR/$APP
cp /tmp/gba_stage/$APP/$APP $DEPLOY_DIR/$APP/$APP
chmod +x $DEPLOY_DIR/$APP/$APP

if [ ! -f $DEPLOY_DIR/$APP/manifest/config/config.yaml ]; then
    echo "[$APP] Copying manifest (first deploy)..."
    cp -r /tmp/gba_stage/$APP/manifest $DEPLOY_DIR/$APP/
fi

echo "[$APP] Starting service..."
systemctl start gba-$APP

# Wait for service to become active (up to 8s)
for i in $(seq 1 8); do
    status=$(systemctl is-active gba-$APP 2>/dev/null || echo 'unknown')
    if [ "$status" = "active" ]; then break; fi
    sleep 1
done

status=$(systemctl is-active gba-$APP 2>/dev/null || echo 'unknown')
echo "[$APP] Status: $status"

rm -rf /tmp/gba_stage/$APP
'@
        $shellBody = $shellBody.Replace("__APP__", $app).Replace("__DEPLOY_DIR__", $DEPLOY_DIR)
        Run-RemoteScript -Name "deploy_$app" -Script $shellBody

        Info "[$app] Done"

        # Pause between services to let server stabilize
        if ($app -ne $APPS[-1]) {
            Info "Waiting 3s before next service ..."
            Start-Sleep -Seconds 3
        }
    }
}

# ============================================
# 5. Deploy frontend: small tar per component, scp, unpack on server
# ============================================
function Deploy-Frontend {
    $localFrontendDir = Join-Path $DIST_DIR "frontend"
    if (-not (Test-Path $localFrontendDir)) {
        Warn "Frontend dist not found, skipping"
        return
    }

    Info "===== Deploying admin frontend ====="
    $tarFile = Join-Path $DIST_DIR "frontend.tar.gz"
    Push-Location $localFrontendDir
    try { tar -czf $tarFile * } finally { Pop-Location }

    $sizeMB = [math]::Round((Get-Item $tarFile).Length / 1MB, 2)
    Info "Frontend packed: $sizeMB MB"

    scp -q $tarFile "${SERVER}:/tmp/gba_frontend.tar.gz"
    if ($LASTEXITCODE -ne 0) { Fail "Frontend upload failed" }

    $script = @'
#!/bin/bash
set -e
DEPLOY_DIR="__DEPLOY_DIR__"
mkdir -p $DEPLOY_DIR/admin
rm -rf $DEPLOY_DIR/admin/*
tar -xzf /tmp/gba_frontend.tar.gz -C $DEPLOY_DIR/admin/
rm -f /tmp/gba_frontend.tar.gz
echo "Frontend deployed"
'@
    $script = $script.Replace("__DEPLOY_DIR__", $DEPLOY_DIR)
    Run-RemoteScript -Name "deploy_frontend" -Script $script
    Info "Admin frontend deployed"
}

# ============================================
# 6. Deploy WAP: small tar, scp, unpack on server
# ============================================
function Deploy-Wap {
    $localWapDir = Join-Path $DIST_DIR "wap"
    if (-not (Test-Path $localWapDir)) {
        Warn "WAP dist not found, skipping"
        return
    }

    Info "===== Deploying WAP ====="
    $tarFile = Join-Path $DIST_DIR "wap.tar.gz"
    Push-Location $localWapDir
    try { tar -czf $tarFile * } finally { Pop-Location }

    $sizeMB = [math]::Round((Get-Item $tarFile).Length / 1MB, 2)
    Info "WAP packed: $sizeMB MB"

    scp -q $tarFile "${SERVER}:/tmp/gba_wap.tar.gz"
    if ($LASTEXITCODE -ne 0) { Fail "WAP upload failed" }

    $script = @'
#!/bin/bash
set -e
DEPLOY_DIR="__DEPLOY_DIR__"
mkdir -p $DEPLOY_DIR/wap
rm -rf $DEPLOY_DIR/wap/*
tar -xzf /tmp/gba_wap.tar.gz -C $DEPLOY_DIR/wap/
rm -f /tmp/gba_wap.tar.gz
echo "WAP deployed"
'@
    $script = $script.Replace("__DEPLOY_DIR__", $DEPLOY_DIR)
    Run-RemoteScript -Name "deploy_wap" -Script $script
    Info "WAP deployed"
}

# ============================================
# Main flow
# ============================================
$startTime = Get-Date
Info "GBaseAdmin deploy started (scp + sequential mode)"
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
            Deploy-Backend
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
    $statusScript = @'
#!/bin/bash
for app in system play upload; do
    status=$(systemctl is-active gba-$app 2>/dev/null || echo "unknown")
    echo "gba-$app: $status"
done
'@
    Run-RemoteScript -Name "status_check" -Script $statusScript
}

# Done
$elapsed = [math]::Round(((Get-Date) - $startTime).TotalSeconds)
Info "All done! Elapsed: ${elapsed}s"
Info "Deploy log saved: $LOG_FILE"

# Auto-cleanup logs older than 30 days
Get-ChildItem $LOG_DIR -Filter "deploy_*.log" |
    Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-30) } |
    Remove-Item -Force
