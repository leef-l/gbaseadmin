# ============================================
# GBaseAdmin Windows Deploy Script (WSL rsync + sequential)
# Usage: .\deploy.ps1 [-Only backend|frontend|wap|all] [-SkipBuild] [-SkipUpload]
#
# Optimized for low-resource servers (2C4G):
#   - rsync via WSL (incremental transfer, no tar/unpack on server)
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

# ---------- Encoding for shell scripts ----------
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

# Convert Windows path to WSL path: C:\foo\bar -> /mnt/c/foo/bar
function To-WslPath($winPath) {
    $full = [System.IO.Path]::GetFullPath($winPath)
    $drive = $full.Substring(0, 1).ToLower()
    $rest = $full.Substring(2).Replace('\', '/')
    return "/mnt/$drive$rest"
}

# Write a LF-encoded temp file, return its path
function Write-LfTempFile {
    param([string]$Name, [string]$Content)
    $path = Join-Path $env:TEMP "gba_${Name}.sh"
    $lfContent = $Content -replace "`r`n", "`n" -replace "`r", "`n"
    [System.IO.File]::WriteAllText($path, $lfContent, $script:utf8NoBom)
    return $path
}

# Write a shell script with LF endings, scp to server, execute, cleanup
function Run-RemoteScript {
    param([string]$Name, [string]$Script)
    $localFile = Write-LfTempFile -Name $Name -Content $Script

    scp -q $localFile "${SERVER}:/tmp/gba_${Name}.sh"
    $scpExit = $LASTEXITCODE
    Remove-Item $localFile -Force -ErrorAction SilentlyContinue
    if ($scpExit -ne 0) { Fail "[$Name] Upload script failed" }

    # Use simple ssh command — the script file itself is already LF-clean
    $output = ssh $SERVER "bash /tmp/gba_${Name}.sh; ret=`$?; rm -f /tmp/gba_${Name}.sh; exit `$ret" 2>&1
    $output | ForEach-Object {
        Write-Host "  $_"
        Log "  $_"
    }
    if ($LASTEXITCODE -ne 0) { Fail "[$Name] Remote script failed" }
}

# rsync via WSL: write temp script to avoid PowerShell arg mangling
function Wsl-Rsync {
    param(
        [string]$Src,
        [string]$Dest,
        [switch]$Delete
    )
    $deleteFlag = ""
    if ($Delete) { $deleteFlag = " --delete" }

    $scriptBody = "#!/bin/bash" + "`n" + "rsync -az${deleteFlag} --compress-level=1 --progress -e ssh '${Src}' '${Dest}'"
    # Use unique temp file name to avoid conflicts
    $uniqueName = "rsync_" + [System.IO.Path]::GetRandomFileName().Replace(".", "")
    $localFile = Write-LfTempFile -Name $uniqueName -Content $scriptBody
    $wslScriptPath = To-WslPath $localFile

    $output = wsl bash $wslScriptPath 2>&1
    $rsyncExit = $LASTEXITCODE
    Remove-Item $localFile -Force -ErrorAction SilentlyContinue

    $output | ForEach-Object {
        Write-Host "  $_"
        Log "  $_"
    }
    return $rsyncExit
}

# ---------- Pre-checks ----------
Check-Command "go"
Check-Command "ssh"
Check-Command "scp"
Check-Command "wsl"

# Verify WSL has rsync
$wslRsyncCheck = wsl which rsync 2>&1
if ($LASTEXITCODE -ne 0) {
    Fail "rsync not found in WSL, run: wsl sudo apt install rsync"
}

# Verify SSH connectivity before spending time on builds
if (-not $SkipUpload) {
    Info "Checking server connectivity ..."
    # Test Windows SSH (used for scp and remote scripts)
    ssh -o ConnectTimeout=5 -o BatchMode=yes $SERVER "echo ok" >$null 2>&1
    if ($LASTEXITCODE -ne 0) { Fail "Cannot connect to $SERVER via SSH. Check your SSH key and network." }
    # Test WSL SSH (used for rsync)
    $wslSshTest = Write-LfTempFile -Name "ssh_test" -Content "#!/bin/bash`nssh -o ConnectTimeout=5 -o BatchMode=yes ${SERVER} 'echo ok'"
    $wslSshTestPath = To-WslPath $wslSshTest
    wsl bash $wslSshTestPath >$null 2>&1
    $wslSshExit = $LASTEXITCODE
    Remove-Item $wslSshTest -Force -ErrorAction SilentlyContinue
    if ($wslSshExit -ne 0) { Fail "Cannot connect to $SERVER via WSL SSH. Run: wsl ssh $SERVER 'echo ok' to diagnose." }
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
# 4. Deploy backend: WSL rsync + sequential restart (one by one)
#    Key: deploy one service at a time to avoid OOM on 2C4G server
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

        # Step 1: rsync via WSL to staging area on server
        Info "[$app] Uploading via rsync ..."
        $wslSrc = "$(To-WslPath $localAppDir)/"
        ssh $SERVER "mkdir -p /tmp/gba_stage/$app"
        $exitCode = Wsl-Rsync -Src $wslSrc -Dest "${SERVER}:/tmp/gba_stage/$app/"
        if ($exitCode -ne 0) { Fail "[$app] rsync upload failed" }
        Info "[$app] Upload OK"

        # Step 2: stop -> replace -> start (on server via temp script)
        # Use single-quoted here-string to prevent PowerShell variable expansion
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
# 5. Deploy frontend: rsync via WSL directly to target
# ============================================
function Deploy-Frontend {
    $localFrontendDir = Join-Path $DIST_DIR "frontend"
    if (-not (Test-Path $localFrontendDir)) {
        Warn "Frontend dist not found, skipping"
        return
    }

    Info "===== Deploying admin frontend (rsync) ====="
    $wslSrc = "$(To-WslPath $localFrontendDir)/"
    ssh $SERVER "mkdir -p $DEPLOY_DIR/admin"
    $exitCode = Wsl-Rsync -Src $wslSrc -Dest "${SERVER}:$DEPLOY_DIR/admin/" -Delete
    if ($exitCode -ne 0) { Fail "Frontend rsync failed" }
    Info "Admin frontend deployed"
}

# ============================================
# 6. Deploy WAP: rsync via WSL directly to target
# ============================================
function Deploy-Wap {
    $localWapDir = Join-Path $DIST_DIR "wap"
    if (-not (Test-Path $localWapDir)) {
        Warn "WAP dist not found, skipping"
        return
    }

    Info "===== Deploying WAP (rsync) ====="
    $wslSrc = "$(To-WslPath $localWapDir)/"
    ssh $SERVER "mkdir -p $DEPLOY_DIR/wap"
    $exitCode = Wsl-Rsync -Src $wslSrc -Dest "${SERVER}:$DEPLOY_DIR/wap/" -Delete
    if ($exitCode -ne 0) { Fail "WAP rsync failed" }
    Info "WAP deployed"
}

# ============================================
# Main flow
# ============================================
$startTime = Get-Date
Info "GBaseAdmin deploy started (WSL rsync + sequential mode)"
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
