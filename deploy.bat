@echo off
chcp 65001 >nul
title GBaseAdmin 一键部署

echo ==========================================
echo   GBaseAdmin 一键部署
echo ==========================================
echo.
echo   [1] 全量部署（后端 + 管理端 + WAP）
echo   [2] 只部署后端
echo   [3] 只部署管理端前端
echo   [4] 只部署 WAP 端
echo   [5] 只打包不上传
echo   [0] 退出
echo.
set /p choice=请选择:

if "%choice%"=="1" powershell -ExecutionPolicy Bypass -File "%~dp0deploy.ps1" -Only all
if "%choice%"=="2" powershell -ExecutionPolicy Bypass -File "%~dp0deploy.ps1" -Only backend
if "%choice%"=="3" powershell -ExecutionPolicy Bypass -File "%~dp0deploy.ps1" -Only frontend
if "%choice%"=="4" powershell -ExecutionPolicy Bypass -File "%~dp0deploy.ps1" -Only wap
if "%choice%"=="5" powershell -ExecutionPolicy Bypass -File "%~dp0deploy.ps1" -Only all -SkipUpload
if "%choice%"=="0" exit

echo.
pause
