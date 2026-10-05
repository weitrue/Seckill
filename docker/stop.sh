#!/usr/bin/env bash
# 关闭 docker-compose 起的所有依赖容器
#   bash docker/stop.sh         # 停容器,保留数据卷
#   PURGE=1 bash docker/stop.sh # 停容器 + 清空数据卷(谨慎!会清掉 redis/mysql/etcd 所有数据)

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

log() { printf '\033[1;33m[seckill]\033[0m %s\n' "$*"; }

if [ "${PURGE:-0}" = "1" ]; then
    log "停止依赖并清空数据卷"
    docker compose -f docker/docker-compose.yml down -v
else
    log "停止依赖,数据卷保留"
    docker compose -f docker/docker-compose.yml down
fi
