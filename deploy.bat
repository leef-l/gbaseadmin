@echo off
chcp 65001 >nul
title GBaseAdmin Deploy

echo.
echo   ==========================================
echo     GBaseAdmin Deploy
echo     Optimized for 2C4G low-resource servers
echo   ==========================================
echo.
echo     --- Fast Mode (parallel build + hash skip + SSH reuse) ---
echo     [1] Fast: Deploy ALL
echo     [2] Fast: Backend only
echo     [3] Fast: Frontend only
echo     [4] Fast: WAP only
echo     [5] Fast: Deploy ALL (force, ignore cache)
echo.
echo     --- Standard Mode ---
echo     [6] Standard: Deploy ALL
echo     [7] Build only (no upload)
echo     [0] Exit
echo.
set /p choice=  Select:

if "%choice%"=="0" exit
if "%choice%"=="1" powershell -NoProfile -ExecutionPolicy Bypass -File "%~dp0deploy-fast.ps1" -Only all
if "%choice%"=="2" powershell -NoProfile -ExecutionPolicy Bypass -File "%~dp0deploy-fast.ps1" -Only backend
if "%choice%"=="3" powershell -NoProfile -ExecutionPolicy Bypass -File "%~dp0deploy-fast.ps1" -Only frontend
if "%choice%"=="4" powershell -NoProfile -ExecutionPolicy Bypass -File "%~dp0deploy-fast.ps1" -Only wap
if "%choice%"=="5" powershell -NoProfile -ExecutionPolicy Bypass -File "%~dp0deploy-fast.ps1" -Only all -Force
if "%choice%"=="6" powershell -NoProfile -ExecutionPolicy Bypass -File "%~dp0deploy-scp.ps1" -Only all
if "%choice%"=="7" powershell -NoProfile -ExecutionPolicy Bypass -File "%~dp0deploy-fast.ps1" -Only all -SkipBuild

echo.
echo   Press any key to close...
pause >nul
