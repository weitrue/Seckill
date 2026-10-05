# Seckill

基于 DDD 架构实现的秒杀服务。

## 目录结构

```
Seckill/
├── application/            # 应用层(gin handler / grpc server,负责参数绑定、鉴权、调下游)
│   └── api/                #   HTTP handler: shop / activity / user ...
├── domain/                 # 领域层(业务规则、聚合、领域服务)
│   ├── shop/               #   商品(入口漏桶 + worker 池 + 响应回写)
│   ├── stock/              #   库存
│   │   ├── memstock/       #     内存预扣(sync.Map + atomic,懒加载自 redisstock)
│   │   └── redisstock/     #     Redis Lua 原子扣减 + 用户去重
│   └── user/               #   用户
├── infrastructure/         # 基础设施层
│   ├── config/             #   配置(viper + cluster 配置监听)
│   ├── mq/                 #   消息队列抽象 + 工厂
│   │   ├── factory/        #     NewFactory / NewQueue / NewPubSub
│   │   ├── memory/         #     本地任务队列(Fan-In 漏桶,实现 mq.Queue)
│   │   ├── kafka/          #     Kafka 实现 mq.PubSub
│   │   └── rabbitmq/       #     RabbitMQ 实现 mq.PubSub(amqp091-go)
│   ├── stores/             #   数据存储(etcd / redis / mysql / cache)
│   ├── services/           #   本地服务(限流器 / 熔断器 / 中间件)
│   └── worker/             #   协程池(grab 抢占式 + schedule 优先级)
├── interfaces/             # 接口层(api / rpc / admin 启动入口)
├── pkg/                    # 通用工具包
├── config/                 # 配置文件
│   └── Seckill.toml
├── docker/                 # 本地开发环境
│   ├── docker-compose.yml  #   Redis / etcd / MySQL
│   ├── start.sh            #   bash 启动脚本
│   ├── stop.sh             #   bash 停止脚本
│   └── start.ps1           #   PowerShell 启动脚本
├── cmd/                    # cobra 子命令(api / admin)
└── main.go
```

## 架构(秒杀主链路)

```
HTTP 请求 (POST /shop/cart/add)
    │
    ├─ 熔断 + 鉴权 + 黑名单(中间件)
    │
    ▼
AddCart (同步)
    │
    ├─ memstock.Sub(内存原子预扣,首次从 Redis 懒加载)
    │     ├─ s < 0  → 直接返回 no stock(大部分失败请求在此被拦)
    │     └─ s ≥ 0 → 继续
    │
    └─ Hijack + queue.Produce(task) → 立刻返回(response 由异步 task 写)
             │
             ▼
         漏桶 (memory, rate=2000) → consumeLoop(单协程搬运) → worker 池 (grab, N=64)
             │
             ▼
         task.Do(): redisstock.Sub (Lua: 原子扣减 + history 去重)
             │
             └─ 回写 HTTP response
```

设计要点:
- **内存预扣**: 快速过滤大多数超额请求,减轻 Redis 压力;允许多实例独立偏扣,Redis Lua 兜底不超卖
- **漏桶 + 协程池**: memory 漏桶控制下游 QPS 阀门,worker 池并发执行慢任务(扣 Redis + 写 response),两者解耦可独立调参
- **Redis Lua**: `get history / get stock / decr stock / set history ex 86400` 一次性原子,保证去重和不超卖

## 前置条件

- Docker Desktop
- Go 1.25+
- 端口未被占用: 6379(redis) / 2379-2380(etcd) / 3306(mysql) / 8080(api) / 8082(rpc)

## 一键启动

### Linux / macOS / Git Bash / WSL

```bash
# 启动依赖 + 交叉编译 + 跑 Seckill
bash docker/start.sh

# 预埋样例库存(A1 活动的 G1 商品,库存 1000)
SEED_STOCK=1 bash docker/start.sh

# 只起依赖,不跑 Seckill
ONLY_DEPS=1 bash docker/start.sh

# 跳过 go build 复用已有二进制
SKIP_BUILD=1 bash docker/start.sh

# 停止依赖(保留数据卷)
bash docker/stop.sh

# 停止 + 清空数据卷(清掉所有数据)
PURGE=1 bash docker/stop.sh
```

### Windows PowerShell

```powershell
.\docker\start.ps1                # 启动依赖 + 交叉编译
.\docker\start.ps1 -SeedStock     # 顺带预埋库存
.\docker\start.ps1 -OnlyDeps      # 只起依赖
.\docker\start.ps1 -SkipBuild     # 跳过构建
```

Windows 下 Linux 二进制无法直接运行,`start.ps1` 编译完会打印进 WSL / 容器运行的提示。

## 验证链路

```bash
# 1. 预埋库存
docker exec seckill_redis redis-cli -n 11 SET "seckill:A1:G1" 1000

# 2. 调用下单接口(鉴权 token 自行准备)
curl -X POST http://localhost:8080/shop/cart/add \
  -H "Content-Type: application/json" \
  -H "Authorization: <your-token>" \
  -d '{"event_id":"A1","goods_id":"G1"}'

# 3. 查看剩余库存
docker exec seckill_redis redis-cli -n 11 GET "seckill:A1:G1"

# 4. 查看某用户是否已下单(Lua history 字段)
docker exec seckill_redis redis-cli -n 11 GET "seckill:A1:G1:<uid>"
```

## 配置

`config/Seckill.toml` 关键段:

| 段 | 说明 |
|---|---|
| `[redis]`   | Redis 地址 + 分 db(lock/cron/rank/cache) |
| `[etcd]`    | etcd endpoints(集群配置监听 + 服务注册) |
| `[mysql]`   | MySQL 地址 + 账号(seckill 库) |
| `[api]`     | HTTP 监听地址 + gRPC 监听地址 + 节点 ttl |
| `[queue.shop]` | shop 用的漏桶 rate/size + worker 池 workerNum/workerSize |

`docker-compose.yml` 的端口 / 账号与 toml 对齐,改任一侧时两边一起改:

| 中间件 | toml | docker-compose |
|---|---|---|
| Redis | `redis.address = "localhost:6379"` | `6379:6379` |
| etcd  | `etcd.endpoints = ["localhost:2379"]` | `2379:2379` |
| MySQL | `mysql.*` | `MYSQL_ROOT_PASSWORD / MYSQL_DATABASE` |

### 调优参数

秒杀吞吐取决于 `queue.shop` 的 `rate` 和 `workerNum`:

- `rate`: 下游 Redis 保护阀门(每秒最多消费这么多任务)
- `size`: 入口漏桶缓冲(瞬时能接住多少请求,再多会 Produce 失败)
- `workerNum`: 并发消费协程数(通常 ≥ rate × 单任务耗时(秒))
- `workerSize`: worker 池 chan 缓冲

典型配置:

| 场景 | rate | workerNum | 预期 QPS 上限 |
|---|---|---|---|
| 开发/联调 | 500 | 16 | ≈500 |
| 预发(Redis 同机房) | 2000 | 64 | ≈2000 |
| 生产(Redis 低延迟) | 5000 | 128 | ≈5000 |

实际数字必须压测,`rate` 调太高(> 10000)会受 `ratelimiter.FanIn` 的毫秒级 sleep 精度拖累。

## MQ 工厂

`infrastructure/mq/factory` 统一入口创建队列:

```go
// 本地任务队列(memory)
q, err := factory.NewQueue(factory.DriverMemory, "shop")
q.Produce(task)

// Kafka / RabbitMQ 发布订阅
ps, err := factory.NewPubSub(factory.DriverRabbitMQ, "order")
ps.Publish("order.created", body, reqID)
ps.Subscribe("order.created", "order-group", func(m mq.ConsumerMsg) error {
    // ...
    return nil
})
```

支持的 driver: `memory` / `kafka` / `rabbitmq`,新增 driver 通过 `factory.Register(name, factory.Func(...))` 注入。

## 构建

```bash
# Linux 交叉编译(宿主机运行二进制)
GOOS=linux GOARCH=amd64 go build -o bin/Seckill_linux main.go

# 直接 go run(非 Windows,因为 pkg/utils/listen.go 用了 Linux-only syscall)
go run main.go api -c config/Seckill.toml
```

## 常见问题

**`docker compose` 命令找不到**
用的是 Docker Compose v2(与 Docker Desktop 一起装)。老版 `docker-compose`(带连字符)也行,把脚本里 `docker compose` 整体替换掉即可。

**etcd 容器健康检查一直 unhealthy**
Bitnami 镜像的 etcdctl 位置可能变化,如果 `docker logs seckill_etcd` 显示 ready 但 health check 持续 fail,可以把 compose 里对应 healthcheck 临时注释掉。

**MySQL 第一次启动很慢**
初始化数据目录和 `seckill` 库需要几十秒,脚本最多等 60 秒。

**Windows 下跑不了 Linux 二进制**
是的,Windows 下 Linux ELF 二进制无法直接执行,进 WSL(推荐)或启动一个 Linux 容器挂载项目目录来跑。
