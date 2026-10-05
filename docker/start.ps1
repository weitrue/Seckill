#Requires -Version 5.1
<#
.SYNOPSIS
    Seckill 本地启动脚本(Windows PowerShell 版)

.DESCRIPTION
    1. 起 docker-compose 的依赖(redis + etcd + mysql)
    2. 等依赖就绪(healthcheck)
    3. 创建必要的运行时文件(blacklist.txt)
    4. 可选:预埋一条库存样例
    5. 启动 Seckill api 服务(交叉编译 Linux 二进制后用 WSL 运行,或直接在 Linux 容器里跑)

.PARAMETER OnlyDeps
    只起依赖容器,不启动 Seckill 服务

.PARAMETER SkipBuild
    跳过 go build,复用已有 bin/Seckill_linux

.PARAMETER SeedStock
    预埋样例库存到 Redis

.EXAMPLE
    .\docker\start.ps1 -OnlyDeps
    .\docker\start.ps1 -SeedStock
    .\docker\start.ps1 -SeedStock -SkipBuild
#>

[CmdletBinding()]
param(
    [switch]$OnlyDeps,
    [switch]$SkipBuild,
    [switch]$SeedStock,
    [string]$SeedActivityID = "A1",
    [string]$SeedGoodsID = "G1",
    [int]$SeedStockValue = 1000,
    [int]$SeedRedisDB = 11
)

$ErrorActionPreference = "Stop"
$RootDir = (Resolve-Path "$PSScriptRoot\..").Path
Set-Location $RootDir

$BinPath       = "bin\Seckill_linux"
$ConfigPath    = "config/Seckill.toml"
$BlacklistFile = "blacklist.txt"

function Write-Log($msg)   { Write-Host "[seckill] $msg" -ForegroundColor Green }
function Write-Fatal($msg) { Write-Host "[seckill][fatal] $msg" -ForegroundColor Red; exit 1 }

# 1. 起依赖
Write-Log "启动 Redis / etcd / MySQL..."
docker compose -f docker/docker-compose.yml up -d
if ($LASTEXITCODE -ne 0) { Write-Fatal "docker compose up 失败" }

# 2. 等就绪
function Wait-Healthy($containerName) {
    $maxWait = 60
    for ($i = 1; $i -le $maxWait; $i++) {
        try {
            $status = docker inspect -f '{{.State.Health.Status}}' $containerName 2>$null
        } catch { $status = "none" }

        if ($status -eq "healthy") {
            Write-Log "$containerName 就绪"
            return
        }
        if ($i % 5 -eq 0) {
            Write-Log "等待 $containerName 就绪... ($i/$maxWait, 当前:$status)"
        }
        Start-Sleep -Seconds 1
    }
    Write-Fatal "$containerName 启动超时"
}
Wait-Healthy "seckill_redis"
Wait-Healthy "seckill_etcd"
Wait-Healthy "seckill_mysql"

# 3. 运行时文件
if (-not (Test-Path $BlacklistFile)) {
    Write-Log "创建空的 $BlacklistFile"
    New-Item -Path $BlacklistFile -ItemType File | Out-Null
}

# 4. 预埋库存
if ($SeedStock) {
    $key = "seckill:${SeedActivityID}:${SeedGoodsID}"
    Write-Log "预埋库存: db=$SeedRedisDB, key=$key, value=$SeedStockValue"
    docker exec seckill_redis redis-cli -n $SeedRedisDB SET $key $SeedStockValue | Out-Null
}

# 5. 只起依赖
if ($OnlyDeps) {
    Write-Log "依赖已启动,OnlyDeps 跳过 Seckill 启动"
    exit 0
}

# 6. 构建
if (-not $SkipBuild) {
    Write-Log "交叉编译 Seckill (GOOS=linux GOARCH=amd64)..."
    $env:GOOS = "linux"
    $env:GOARCH = "amd64"
    try {
        go build -o $BinPath main.go
        if ($LASTEXITCODE -ne 0) { Write-Fatal "go build 失败" }
    } finally {
        $env:GOOS = ""
        $env:GOARCH = ""
    }
}

if (-not (Test-Path $BinPath)) {
    Write-Fatal "找不到可执行文件 $BinPath,去掉 -SkipBuild 或手动构建"
}

# 7. 启动 Seckill(Windows 下 Linux 二进制跑不动,提示用户去 WSL)
Write-Log "二进制已生成: $BinPath"
Write-Log "Windows 下无法直接运行 Linux 二进制,请在 WSL / 容器 / Linux 环境执行:"
Write-Host ""
Write-Host "    ./$($BinPath -replace '\\','/') api -c $ConfigPath" -ForegroundColor Cyan
Write-Host ""
Write-Log "或者直接在容器里跑: docker run --rm --network host -v `${PWD}:/app -w /app golang:1.25 ./$($BinPath -replace '\\','/') api -c $ConfigPath"
