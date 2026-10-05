/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午8:32
 * Description: shop 场景的入口与消费装配
 *   架构:
 *     Produce(shop.Handle)
 *        │ RateLimiter 漏桶(rate = 下游 QPS 保护阀门)
 *        ▼
 *     mq.Queue (memory)
 *        │ 单 goroutine Consume (消费节奏由漏桶控制)
 *        ▼
 *     worker.Worker (grab 协程池, N 个 goroutine 并发执行 task)
 *        │
 *        ▼
 *     Redis Sub (Lua 原子扣减)
 **/

package shop

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/weitrue/Seckill/internal/domain/stock/redisstock"
	"github.com/weitrue/Seckill/internal/infrastructure/mq"
	"github.com/weitrue/Seckill/internal/infrastructure/mq/factory"
	"github.com/weitrue/Seckill/internal/infrastructure/worker"
	"github.com/weitrue/Seckill/internal/infrastructure/worker/grab"
	"github.com/weitrue/Seckill/pkg/utils"
)

const (
	// 配置 key: queue.shop.*
	queueName      = "shop"
	workerNumKey   = "queue.shop.workerNum"
	workerSizeKey  = "queue.shop.workerSize"
	defaultWorkerN = 0    // 0 表示自动计算 = NumCPU * 4
	defaultWorkerS = 1024 // worker 池内部 chan 缓冲默认值
)

var (
	queue   mq.Queue      // 入口漏桶队列
	workers worker.Worker // 并发消费协程池
	initMu  sync.Mutex    // 保护 Init / Close
	inited  bool          // 幂等标记
)

/*Init
 *@Description: 初始化 shop 的入口队列和消费协程池
 *   读取 viper:
 *     queue.shop.rate        漏桶速率(mq/memory 读)
 *     queue.shop.size        漏桶缓冲(mq/memory 读)
 *     queue.shop.workerNum   消费协程数(本函数读,<=0 时取 runtime.NumCPU()*4)
 *     queue.shop.workerSize  worker 池 chan 缓冲(本函数读,<=0 取 1024)
 */
func Init() {
	logType := "ShopInit"
	initMu.Lock()
	defer initMu.Unlock()
	if inited {
		logrus.Warnf("logType:%s, msg:already inited, skip", logType)
		return
	}

	// 1. 创建漏桶队列
	q, err := factory.NewQueue(factory.DriverMemory, queueName)
	if err != nil {
		logrus.Errorf("logType:%s, err:%s, driver:%s, name:%s",
			logType, err.Error(), factory.DriverMemory, queueName)
		panic(err)
	}
	queue = q

	// 2. 创建 worker 协程池
	num := viper.GetInt(workerNumKey)
	if num <= 0 {
		num = runtime.NumCPU() * 4
	}
	size := viper.GetInt(workerSizeKey)
	if size <= 0 {
		size = defaultWorkerS
	}
	workers = grab.NewWorker(num, size)
	logrus.Infof("logType:%s, msg:worker pool started, workerNum:%d, workerSize:%d",
		logType, num, size)

	// 3. 启动 1 个 consume goroutine 作搬运工
	go consumeLoop()

	inited = true
}

/*Close
 *@Description: 优雅关闭:先关队列(让 Consume 返回 err),再关 worker 池
 *@return error
 */
func Close() error {
	logType := "ShopClose"
	initMu.Lock()
	defer initMu.Unlock()
	if !inited {
		return nil
	}
	inited = false

	var firstErr error
	if queue != nil {
		if err := queue.Close(); err != nil {
			logrus.Errorf("logType:%s, err:%s, step:close queue", logType, err.Error())
			firstErr = err
		}
	}
	if workers != nil {
		if err := workers.Close(); err != nil {
			logrus.Errorf("logType:%s, err:%s, step:close workers", logType, err.Error())
			if firstErr == nil {
				firstErr = err
			}
		}
	}
	return firstErr
}

/*consumeLoop
 *@Description: 从漏桶队列拉任务,转交 worker 池并发执行
 *   - 单 goroutine 做搬运,协程数量可控(规范 21)
 *   - 外层 recover 兜底整个消费循环;每个 task 的 recover 已在 grab.worker.process 内处理
 *   - Consume 返回 err 表示队列已关闭,自然退出
 */
func consumeLoop() {
	logType := "ShopConsumeLoop"
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("logType:%s, panic:%v", logType, r)
		}
	}()
	for {
		task, err := queue.Consume()
		if err != nil {
			logrus.Infof("logType:%s, msg:queue closed, err:%s", logType, err.Error())
			return
		}
		if ok := workers.Push(task); !ok {
			logrus.Warnf("logType:%s, msg:worker pool closed, drop task and exit loop", logType)
			return
		}
	}
}

const (
	OK         = 0
	ErrNoStock = 1001
	ErrRedis   = 1002
	ErrTimeout = 1003

	requestTimeout = 60
)

type Context struct {
	Request *http.Request
	Conn    net.Conn
	Writer  *bufio.ReadWriter
	GoodsID string
	EventID string
	UID     string
}

/*Handle
 *@Description: 处理一次抢购请求: 请求 -> 熔断 -> 令牌桶限流 -> 业务逻辑 -> 漏桶限流 -> 减库存
 *   本函数只做"投递 task 到漏桶队列",消费侧由 consumeLoop + worker 池完成
 *@param ctx 请求上下文
 */
func Handle(ctx *Context) {
	start := time.Now().Unix()
	t := func() {
		data := &utils.Response{
			Code: OK,
			Data: nil,
			Msg:  "ok",
		}
		status := http.StatusOK
		now := time.Now().Unix()
		if now-start > requestTimeout {
			data.Msg = "request timeout"
			data.Code = ErrTimeout
		} else {
			// 异步真扣:走 Redis Lua(原子扣减 + 用户去重)
			st, err := redisstock.NewRedisStock(ctx.EventID, ctx.GoodsID)
			if err != nil {
				logrus.Errorf("logType:ShopHandleTask, err:%s, step:redisstock init, eventID:%s, goodsID:%s, uid:%s",
					err.Error(), ctx.EventID, ctx.GoodsID, ctx.UID)
				data.Msg = err.Error()
				data.Code = ErrRedis
			} else if s, subErr := st.Sub(ctx.UID); subErr != nil {
				logrus.Errorf("logType:ShopHandleTask, err:%s, step:redisstock sub, eventID:%s, goodsID:%s, uid:%s",
					subErr.Error(), ctx.EventID, ctx.GoodsID, ctx.UID)
				data.Msg = subErr.Error()
				data.Code = ErrRedis
			} else if s < 0 {
				data.Msg = "no stock"
				data.Code = ErrNoStock
			}
		}
		// 此处实现操作购物车的逻辑

		body, _ := json.Marshal(data)
		resp := &http.Response{
			Proto:         ctx.Request.Proto,
			ProtoMinor:    ctx.Request.ProtoMinor,
			ProtoMajor:    ctx.Request.ProtoMajor,
			Header:        make(http.Header),
			ContentLength: int64(len(body)),
			Body:          ioutil.NopCloser(bytes.NewReader(body)),
			StatusCode:    status,
			Close:         false,
		}
		resp.Header.Set("Content-Type", "application/json")
		resp.Write(ctx.Writer)
		ctx.Writer.Flush()
		ctx.Conn.Close()
	}

	if err := queue.Produce(worker.QueueTaskFunc(t)); err != nil {
		logrus.Errorf("logType:ShopHandle, err:%s, goodsID:%s, uid:%s",
			err.Error(), ctx.GoodsID, ctx.UID)
	}
}
