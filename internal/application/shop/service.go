/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2024/10/06
 * Description: shop 应用服务: 协议无关的下单用例编排
 *   链路:
 *     handler.AddCart
 *        │ cmd(AddCartCommand) + ctx
 *        ▼
 *     Service.AddCart
 *        ├─ 1. 活动校验: ActivityLookup.Get + IsOngoing + FindGoods
 *        ├─ 2. memstock.Sub(ctx)                        [同步内存预扣,快速拦截]
 *        ├─ 3. queue.Produce(task)                      [投递到漏桶]
 *        │        task 内部:
 *        │          redisstock.Sub(ctx, uid)            [Redis Lua 真扣+去重]
 *        │          OrderCreator.Create(...)            [落库;唯一键做兜底幂等]
 *        └─ 4. select { <-resultCh / ctx.Done / timeout }
 **/

package shop

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/weitrue/Seckill/internal/application/order"
	activityDomain "github.com/weitrue/Seckill/internal/domain/activity"
	orderDomain "github.com/weitrue/Seckill/internal/domain/order"
	"github.com/weitrue/Seckill/internal/domain/stock"
	"github.com/weitrue/Seckill/internal/infrastructure/mq"
	"github.com/weitrue/Seckill/internal/infrastructure/worker"
)

// ActivityLookup 活动查询依赖(通常由 application/activity.Service 实现)
type ActivityLookup interface {
	Get(ctx context.Context, id int64) (*activityDomain.Activity, error)
}

// OrderCreator 订单创建依赖(通常由 application/order.Service 实现)
type OrderCreator interface {
	Create(ctx context.Context, cmd order.CreateCommand) (*orderDomain.Order, error)
}

// Dependencies Service 的构造依赖
type Dependencies struct {
	// NewMemStock 内存预扣库存工厂(通常 = memstock.NewFactory(loader).New)
	NewMemStock stock.Factory
	// NewRedisStock Redis 真扣库存工厂(通常 = redisstock.NewRedisStock)
	NewRedisStock stock.Factory
	// Queue 入口漏桶队列,外部构造好后注入(本阶段用 memory driver)
	Queue mq.Queue
	// Workers 消费端协程池
	Workers worker.Worker
	// ActivityLookup 活动查询(校验 IsOngoing + 查商品活动价)
	ActivityLookup ActivityLookup
	// OrderCreator 订单落库
	OrderCreator OrderCreator
	// TaskTimeout AddCart 等 task 的最大等待时间;<=0 使用默认 30s
	TaskTimeout time.Duration
}

// Service shop 应用服务
type Service struct {
	deps    Dependencies
	timeout time.Duration

	startMu sync.Mutex
	started bool
	stopCh  chan struct{}
}

const defaultTaskTimeout = 30 * time.Second

/*NewService
 *@Description: 构造 Service 实例,校验必填依赖
 *@param deps 依赖
 *@return *Service
 *@return error 必填依赖缺失时返回错误
 */
func NewService(deps Dependencies) (*Service, error) {
	if deps.NewMemStock == nil {
		return nil, errors.New("shop.Service: NewMemStock is required")
	}
	if deps.NewRedisStock == nil {
		return nil, errors.New("shop.Service: NewRedisStock is required")
	}
	if deps.Queue == nil {
		return nil, errors.New("shop.Service: Queue is required")
	}
	if deps.Workers == nil {
		return nil, errors.New("shop.Service: Workers is required")
	}
	if deps.ActivityLookup == nil {
		return nil, errors.New("shop.Service: ActivityLookup is required")
	}
	if deps.OrderCreator == nil {
		return nil, errors.New("shop.Service: OrderCreator is required")
	}
	timeout := deps.TaskTimeout
	if timeout <= 0 {
		timeout = defaultTaskTimeout
	}
	return &Service{
		deps:    deps,
		timeout: timeout,
		stopCh:  make(chan struct{}),
	}, nil
}

// activityCheck 活动校验结果,内部透传给 doDeduct
type activityCheck struct {
	activityID int64
	price      int64 // 活动价(分)
}

/*AddCart
 *@Description: 下单(加购)用例;协议无关,ctx 全链路携带
 *@receiver s
 *@param ctx 请求上下文(链路追踪 / 超时 / 取消)
 *@param cmd 下单命令
 *@return AddCartResult 业务结果(即使命中业务错误码也通过返回值表达)
 *@return error 系统级错误(ctx cancel 等)
 */
func (s *Service) AddCart(ctx context.Context, cmd AddCartCommand) (AddCartResult, error) {
	logType := "ShopServiceAddCart"
	if cmd.UID == "" || cmd.ActivityID == "" || cmd.GoodsID == "" {
		return AddCartResult{Code: CodeBadRequest, Msg: "bad command"},
			fmt.Errorf("empty cmd field: uid=%q activityID=%q goodsID=%q",
				cmd.UID, cmd.ActivityID, cmd.GoodsID)
	}

	// 1. 活动校验(走进程内缓存 / Redis / MySQL 多层)
	check, result := s.checkActivity(ctx, cmd)
	if result != nil {
		return *result, nil
	}

	// 2. 内存预扣
	memStock, err := s.deps.NewMemStock(cmd.ActivityID, cmd.GoodsID)
	if err != nil {
		logrus.Errorf("logType:%s, err:%s, step:newMemStock, cmd:%+v",
			logType, err.Error(), cmd)
		return AddCartResult{Code: CodeRedisErr, Msg: "internal error"}, err
	}
	remain, err := memStock.Sub(ctx, cmd.UID)
	if err != nil {
		logrus.Errorf("logType:%s, err:%s, step:memstock.Sub, cmd:%+v",
			logType, err.Error(), cmd)
		return AddCartResult{Code: CodeRedisErr, Msg: "internal error"}, err
	}
	if remain < 0 {
		return AddCartResult{Code: CodeNoStock, Msg: "no stock"}, nil
	}

	// 3. 投递异步真扣+落单任务(闭包捕获 ctx + check + resultCh)
	resultCh := make(chan AddCartResult, 1)
	task := worker.QueueTaskFunc(func() {
		resultCh <- s.doDeduct(ctx, cmd, check)
	})
	if err := s.deps.Queue.Produce(task); err != nil {
		logrus.Errorf("logType:%s, err:%s, step:queue.Produce, cmd:%+v",
			logType, err.Error(), cmd)
		return AddCartResult{Code: CodeRedisErr, Msg: "server busy"}, err
	}

	// 4. 等 task result / ctx 取消 / 超时
	select {
	case r := <-resultCh:
		return r, nil
	case <-ctx.Done():
		logrus.Warnf("logType:%s, msg:ctx canceled, err:%s, cmd:%+v",
			logType, ctx.Err().Error(), cmd)
		return AddCartResult{Code: CodeTimeout, Msg: "canceled"}, ctx.Err()
	case <-time.After(s.timeout):
		logrus.Warnf("logType:%s, msg:task timeout, cmd:%+v, timeout:%s",
			logType, cmd, s.timeout.String())
		return AddCartResult{Code: CodeTimeout, Msg: "timeout"}, nil
	}
}

/*checkActivity
 *@Description: 活动校验(存在 / 进行中 / 商品属于活动)
 *@receiver s
 *@param ctx
 *@param cmd
 *@return *activityCheck 校验通过时返回,否则 nil
 *@return *AddCartResult 失败时返回要给前端的结果(调用方直接 return)
 */
func (s *Service) checkActivity(ctx context.Context, cmd AddCartCommand) (*activityCheck, *AddCartResult) {
	logType := "ShopActivityCheck"
	aid, err := strconv.ParseInt(cmd.ActivityID, 10, 64)
	if err != nil || aid <= 0 {
		return nil, &AddCartResult{Code: CodeBadRequest, Msg: "invalid activity id"}
	}
	a, err := s.deps.ActivityLookup.Get(ctx, aid)
	if err != nil {
		if errors.Is(err, activityDomain.ErrNotFound) {
			return nil, &AddCartResult{Code: CodeNoActivity, Msg: "no activity"}
		}
		logrus.Errorf("logType:%s, err:%s, cmd:%+v", logType, err.Error(), cmd)
		return nil, &AddCartResult{Code: CodeRedisErr, Msg: "internal error"}
	}
	if a == nil {
		return nil, &AddCartResult{Code: CodeNoActivity, Msg: "no activity"}
	}
	if !a.IsOngoing(time.Now()) {
		return nil, &AddCartResult{Code: CodeActivityNotOngoing, Msg: "activity not ongoing"}
	}
	g := a.FindGoods(cmd.GoodsID)
	if g == nil {
		return nil, &AddCartResult{Code: CodeNoGoods, Msg: "no goods in activity"}
	}
	return &activityCheck{activityID: aid, price: g.ActivityPrice}, nil
}

/*doDeduct
 *@Description: worker 协程池里真正执行的"扣 Redis + 落单"逻辑
 *@receiver s
 *@param ctx 请求上下文(通过闭包捕获进来)
 *@param cmd
 *@param check 活动校验结果(含活动 int64 ID + 活动价)
 *@return AddCartResult
 */
func (s *Service) doDeduct(ctx context.Context, cmd AddCartCommand, check *activityCheck) AddCartResult {
	logType := "ShopServiceDoDeduct"
	rs, err := s.deps.NewRedisStock(cmd.ActivityID, cmd.GoodsID)
	if err != nil {
		logrus.Errorf("logType:%s, err:%s, step:newRedisStock, cmd:%+v",
			logType, err.Error(), cmd)
		return AddCartResult{Code: CodeRedisErr, Msg: err.Error()}
	}
	remain, err := rs.Sub(ctx, cmd.UID)
	if err != nil {
		logrus.Errorf("logType:%s, err:%s, step:redisstock.Sub, cmd:%+v",
			logType, err.Error(), cmd)
		return AddCartResult{Code: CodeRedisErr, Msg: err.Error()}
	}
	if remain < 0 {
		return AddCartResult{Code: CodeNoStock, Msg: "no stock"}
	}

	// Redis 扣减成功 → 同步创建订单
	o, err := s.deps.OrderCreator.Create(ctx, order.CreateCommand{
		UID:        cmd.UID,
		ActivityID: check.activityID,
		GoodsID:    cmd.GoodsID,
		Price:      check.price,
	})
	if err != nil {
		if errors.Is(err, orderDomain.ErrDuplicate) {
			// 唯一键命中幂等:Redis Lua 的 history 应该已经先拦截,这里是兜底
			logrus.Warnf("logType:%s, msg:order duplicate, cmd:%+v", logType, cmd)
			return AddCartResult{Code: CodeOrderDuplicate, Msg: "already ordered"}
		}
		logrus.Errorf("logType:%s, err:%s, step:order.Create, cmd:%+v",
			logType, err.Error(), cmd)
		return AddCartResult{Code: CodeOrderErr, Msg: err.Error()}
	}
	return AddCartResult{Code: CodeOK, Msg: "ok", OrderID: o.ID}
}
