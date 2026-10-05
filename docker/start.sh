#!/usr/bin/env bash
# Seckill 本地启动脚本
#   1. 起 docker-compose 的依赖(redis + etcd + mysql)
#   2. 等依赖就绪
#   3. 创建必要的运行时文件(blacklist.txt)
#   4. 可选: 预埋一条库存样例(通过 SEED_STOCK=1 开启)
#   5. 启动 Seckill api 服务
#
# 用法:
#   bash docker/start.sh                       # 启动依赖并跑服务
#   SEED_STOCK=1 bash docker/start.sh          # 启动前预埋一条样例库存
#   SKIP_BUILD=1 bash docker/start.sh          # 跳过交叉编译(复用 bin/Seckill_linux)
#   ONLY_DEPS=1 bash docker/start.sh           # 只起依赖,不跑 Seckill

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

BIN_PATH="bin/Seckill_linux"
CONFIG_PATH="config/Seckill.toml"
BLACKLIST_FILE="blacklist.txt"

# 样例库存配置(SEED_STOCK=1 时生效)
SEED_ACTIVITY_ID="${SEED_ACTIVITY_ID:-A1}"
SEED_GOODS_ID="${SEED_GOODS_ID:-G1}"
SEED_STOCK_VALUE="${SEED_STOCK_VALUE:-1000}"
SEED_REDIS_DB="${SEED_REDIS_DB:-11}" # 对齐 toml 的 redis.DB.cache

log() { printf '\033[1;32m[seckill]\033[0m %s\n' "$*"; }
die() { printf '\033[1;31m[seckill][fatal]\033[0m %s\n' "$*" >&2; exit 1; }

# 1. 起依赖
log "启动 Redis / etcd / MySQL..."
docker compose -f docker/docker-compose.yml up -d

# 2. 等依赖就绪(循环检查 healthcheck)
wait_healthy() {
    local name="$1"
    local max=60
    for i in $(seq 1 $max); do
        status=$(docker inspect -f '{{.State.Health.Status}}' "$name" 2>/dev/null || echo "none")
        if [ "$status" = "healthy" ]; then
            log "$name 就绪"
            return 0
        fi
        if [ $((i % 5)) -eq 0 ]; then
            log "等待 $name 就绪... ($i/$max,当前状态:$status)"
        fi
        sleep 1
    done
    die "$name 启动超时"
}
wait_healthy seckill_redis
wait_healthy seckill_etcd
wait_healthy seckill_mysql

# 3. 运行时文件
if [ ! -f "$BLACKLIST_FILE" ]; then
    log "创建空的 $BLACKLIST_FILE"
    touch "$BLACKLIST_FILE"
fi

# 4. 可选: 预埋库存
if [ "${SEED_STOCK:-0}" = "1" ]; then
    key="seckill:${SEED_ACTIVITY_ID}:${SEED_GOODS_ID}"
    log "预埋库存: db=$SEED_REDIS_DB, key=$key, value=$SEED_STOCK_VALUE"
    docker exec seckill_redis redis-cli -n "$SEED_REDIS_DB" SET "$key" "$SEED_STOCK_VALUE" >/dev/null
fi

# 5. 只起依赖模式
if [ "${ONLY_DEPS:-0}" = "1" ]; then
    log "依赖已启动,ONLY_DEPS=1 跳过 Seckill 启动"
    exit 0
fi

# 6. 构建二进制(默认跑,SKIP_BUILD=1 跳过)
if [ "${SKIP_BUILD:-0}" != "1" ]; then
    log "交叉编译 Seckill (GOOS=linux GOARCH=amd64)..."
    GOOS=linux GOARCH=amd64 go build -o "$BIN_PATH" main.go
fi

if [ ! -x "$BIN_PATH" ]; then
    die "找不到可执行文件 $BIN_PATH,去掉 SKIP_BUILD 或手动构建"
fi

# 7. 启动 Seckill
log "启动 Seckill api 服务 (配置: $CONFIG_PATH)"
exec "./$BIN_PATH" api -c "$CONFIG_PATH"
