@echo off
chcp 65001 >nul
title GBaseAdmin Deploy

echo.
echo   ==========================================
echo     GBaseAdmin Deploy
echo   ==========================================
echo.
echo     [1] Deploy ALL (backend + frontend + wap)
echo     [2] Deploy backend only
echo     [3] Deploy frontend only
echo     [4] Deploy wap only
echo     [5] Pack only (no upload)
echo     [0] Exit
echo.
set /p choice=  Select:

if "%choice%"=="0" exit
if "%choice%"=="1" powershell -NoProfile -ExecutionPolicy Bypass -File "%~dp0deploy.ps1" -Only all
if "%choice%"=="2" powershell -NoProfile -ExecutionPolicy Bypass -File "%~dp0deploy.ps1" -Only backend
if "%choice%"=="3" powershell -NoProfile -ExecutionPolicy Bypass -File "%~dp0deploy.ps1" -Only frontend
if "%choice%"=="4" powershell -NoProfile -ExecutionPolicy Bypass -File "%~dp0deploy.ps1" -Only wap
if "%choice%"=="5" powershell -NoProfile -ExecutionPolicy Bypass -File "%~dp0deploy.ps1" -Only all -SkipUpload

echo.
echo   Press any key to close...
pause >nul
