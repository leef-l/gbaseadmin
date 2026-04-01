# ============================================
# GBaseAdmin Fast Deploy
# Usage: .\deploy-fast.ps1 [-Only backend|frontend|wap|all] [-SkipBuild] [-Force]
#
# Optimizations:
#   - Parallel Go compilation (all services at once)
#   - SHA256 hash check: skip upload if binary unchanged
#   - scp -C compressed transfer (great for cross-ocean)
#   - Sequential restart on server (2C4G safe)
# ============================================

param(
    [ValidateSet("all", "backend", "frontend", "wap")]
    [string]$Only = "all",
    [switch]$SkipBuild,
    [switch]$Force
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
$HASH_DIR     = Join-Path $PROJECT_DIR ".deploy-cache"

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

# Write LF-encoded temp file
function Write-LfTempFile {
    param([string]$Name, [string]$Content)
    $path = Join-Path $env:TEMP "gba_${Name}.sh"
    $lfContent = $Content -replace "`r`n", "`n" -replace "`r", "`n"
    [System.IO.File]::WriteAllText($path, $lfContent, $script:utf8NoBom)
    return $path
}

# Run script on remote server
function Run-RemoteScript {
    param([string]$Name, [string]$Script)
    $localFile = Write-LfTempFile -Name $Name -Content $Script

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

# Get SHA256 hash of a file
function Get-FileHash256($path) {
    if (-not (Test-Path $path)) { return "" }
    return (Get-FileHash -Path $path -Algorithm SHA256).Hash
}

# ---------- Pre-checks ----------
Check-Command "go"
Check-Command "ssh"
Check-Command "scp"

# Verify SSH connectivity
Info "Checking server connectivity ..."
ssh -o ConnectTimeout=10 -o BatchMode=yes $SERVER "echo ok" >$null 2>&1
if ($LASTEXITCODE -ne 0) { Fail "Cannot connect to $SERVER. Check SSH key and network." }
Info "Server connectivity OK"

# Create dirs
if (-not (Test-Path $HASH_DIR)) { New-Item -ItemType Directory -Path $HASH_DIR -Force | Out-Null }
if (Test-Path $DIST_DIR) { Remove-Item $DIST_DIR -Recurse -Force }
New-Item -ItemType Directory -Path $DIST_DIR -Force | Out-Null
New-Item -ItemType Directory -Path (Join-Path $DIST_DIR "backend") -Force | Out-Null

# ============================================
# 1. Build backend: PARALLEL compilation
# ============================================
function Build-Backend {
    Info "===== Building backend services (parallel) ====="

    $env:CGO_ENABLED = "0"
    $env:GOOS = "linux"
    $env:GOARCH = "amd64"
    $env:GOPROXY = "https://goproxy.cn,direct"

    $jobs = @()
    foreach ($app in $APPS) {
        $outDir = Join-Path $DIST_DIR "backend\$app"
        New-Item -ItemType Directory -Path $outDir -Force | Out-Null

        $appName = $app
        $appOutDir = $outDir
        $backendDir = $BACKEND_DIR

        Info "Starting parallel build: $appName ..."
        $job = Start-Job -ScriptBlock {
            param($dir, $name, $out)
            $env:CGO_ENABLED = "0"
            $env:GOOS = "linux"
            $env:GOARCH = "amd64"
            $env:GOPROXY = "https://goproxy.cn,direct"
            Set-Location $dir
            $result = go build -ldflags "-s -w" -o "$out\$name" ".\app\$name\main.go" 2>&1
            if ($LASTEXITCODE -ne 0) {
                throw "$name build failed: $result"
            }
            return "$name OK"
        } -ArgumentList $backendDir, $appName, $appOutDir

        $jobs += @{ Name = $appName; Job = $job; OutDir = $appOutDir }
    }

    # Wait for all builds
    $failed = $false
    foreach ($j in $jobs) {
        $result = Receive-Job -Job $j.Job -Wait -ErrorAction SilentlyContinue
        if ($j.Job.State -eq 'Failed') {
            Warn "$($j.Name) build FAILED"
            $failed = $true
        } else {
            # Copy manifest
            $manifestSrc = Join-Path $BACKEND_DIR "app\$($j.Name)\manifest"
            if (Test-Path $manifestSrc) {
                Copy-Item $manifestSrc -Destination $j.OutDir -Recurse -Force
            }
            Info "$($j.Name) build OK"
        }
        Remove-Job -Job $j.Job -Force
    }

    # Restore env
    $env:GOOS = ""
    $env:GOARCH = ""

    if ($failed) { Fail "One or more backend builds failed" }
}

# ============================================
# 2. Build frontend
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
# 3. Build WAP
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
# 4. Deploy backend: hash check + scp -C + sequential restart
# ============================================
function Deploy-Backend {
    Info "===== Deploying backend services ====="

    foreach ($app in $APPS) {
        $localBinary = Join-Path $DIST_DIR "backend\$app\$app"
        if (-not (Test-Path $localBinary)) {
            Warn "$app binary not found, skipping"
            continue
        }

        # Hash check: skip if unchanged
        $currentHash = Get-FileHash256 $localBinary
        $hashFile = Join-Path $HASH_DIR "$app.hash"
        $lastHash = ""
        if (Test-Path $hashFile) { $lastHash = (Get-Content $hashFile -Raw).Trim() }

        if (-not $Force -and $currentHash -eq $lastHash) {
            Info "[$app] Binary unchanged, skipping"
            continue
        }

        Info "--- [$app] Start ---"

        # Upload with compression
        $sizeMB = [math]::Round((Get-Item $localBinary).Length / 1MB, 2)
        Info "[$app] Uploading binary ($sizeMB MB, compressed) ..."

        ssh $SERVER "mkdir -p /tmp/gba_stage/$app"
        scp -C -q "$localBinary" "${SERVER}:/tmp/gba_stage/$app/$app"
        if ($LASTEXITCODE -ne 0) { Fail "[$app] Upload binary failed" }

        # Upload manifest
        $manifestDir = Join-Path $DIST_DIR "backend\$app\manifest"
        if (Test-Path $manifestDir) {
            scp -C -q -r "$manifestDir" "${SERVER}:/tmp/gba_stage/$app/"
        }
        Info "[$app] Upload OK"

        # Stop -> replace -> start
        Info "[$app] Restarting service ..."
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

for i in $(seq 1 8); do
    status=$(systemctl is-active gba-$APP 2>/dev/null || echo 'unknown')
    if [ "$status" = "active" ]; then break; fi
    sleep 1
done

status=$(systemctl is-active gba-$APP 2>/dev/null || echo 'unknown')
echo "[$APP] Status: $status"
if [ "$status" != "active" ]; then exit 1; fi

rm -rf /tmp/gba_stage/$APP
'@
        $shellBody = $shellBody.Replace("__APP__", $app).Replace("__DEPLOY_DIR__", $DEPLOY_DIR)
        Run-RemoteScript -Name "deploy_$app" -Script $shellBody

        # Save hash on success
        $currentHash | Out-File -FilePath $hashFile -NoNewline -Encoding ascii
        Info "[$app] Done"

        # Cooldown
        if ($app -ne $APPS[-1]) {
            Info "Waiting 3s ..."
            Start-Sleep -Seconds 3
        }
    }
}

# ============================================
# 5. Deploy frontend: tar + scp -C
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

    scp -C -q $tarFile "${SERVER}:/tmp/gba_frontend.tar.gz"
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
# 6. Deploy WAP: tar + scp -C
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

    scp -C -q $tarFile "${SERVER}:/tmp/gba_wap.tar.gz"
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
Info "GBaseAdmin FAST deploy started"
Info "Target: $Only | Force: $Force"
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

# Step 2: Deploy
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

# Step 3: Final status
if ($Only -eq "all" -or $Only -eq "backend") {
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

# Cleanup dist
Remove-Item $DIST_DIR -Recurse -Force -ErrorAction SilentlyContinue

# Done
$elapsed = [math]::Round(((Get-Date) - $startTime).TotalSeconds)
Info "All done! Elapsed: ${elapsed}s"
Info "Deploy log saved: $LOG_FILE"

# Auto-cleanup logs older than 30 days
Get-ChildItem $LOG_DIR -Filter "deploy_*.log" |
    Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-30) } |
    Remove-Item -Force
