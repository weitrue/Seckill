/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午5:04
 * Description: HTTP API 服务启动入口
 *   职责:
 *     1. 初始化基础设施(redis / mysql / listener)
 *     2. 装配仓储 + 应用服务(依赖显式构造并注入)
 *     3. 预热活动数据到 Redis(不阻塞启动)
 *     4. 启动 gin,注册路由
 *     5. Exit 时优雅关闭
 **/

package api

import (
	"context"
	"net"
	"runtime"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/weitrue/Seckill/interfaces/http/api/handler"
	activityapp "github.com/weitrue/Seckill/internal/application/activity"
	orderapp "github.com/weitrue/Seckill/internal/application/order"
	shopapp "github.com/weitrue/Seckill/internal/application/shop"
	"github.com/weitrue/Seckill/internal/domain/stock"
	"github.com/weitrue/Seckill/internal/infrastructure/config"
	"github.com/weitrue/Seckill/internal/infrastructure/mq/factory"
	"github.com/weitrue/Seckill/internal/infrastructure/persistence/activityrepo"
	"github.com/weitrue/Seckill/internal/infrastructure/persistence/orderrepo"
	"github.com/weitrue/Seckill/internal/infrastructure/persistence/stockrepo/memstock"
	"github.com/weitrue/Seckill/internal/infrastructure/persistence/stockrepo/redisstock"
	"github.com/weitrue/Seckill/internal/infrastructure/stores/mysql"
	"github.com/weitrue/Seckill/internal/infrastructure/stores/redis"
	"github.com/weitrue/Seckill/internal/infrastructure/worker/grab"
	utils2 "github.com/weitrue/Seckill/pkg/utils"
)

const (
	// 配置 key
	workerNumKey  = "queue.shop.workerNum"
	workerSizeKey = "queue.shop.workerSize"
	cacheDBKey    = "redis.DB.cache"

	defaultWorkerSize = 1024
	defaultCacheDB    = 11
	prewarmTimeout    = 15 * time.Second
)

var (
	listener net.Listener
	err      error
	shopSvc  *shopapp.Service // 保存给 Exit 优雅关闭
)

/*Run
 *@Description: 启动 HTTP API 服务
 *@return error
 */
func Run() error {
	logType := "ApiRun"
	bind := viper.GetString("api.bind")
	logrus.Info("run api server on ", bind)
	listener, err = utils2.Listen(config.GetTcpNet(), bind)
	if err != nil {
		return err
	}

	engine := gin.New()
	utils2.UpdateProcess("api")
	config.WatchBlacklistConfig()

	engine.Use(gin.Logger())
	pprof.Register(engine)

	// 1. Redis
	logrus.Info("------------------ init redis ------------------")
	if err := redis.Init(); err != nil {
		panic(err)
	}

	// 2. MySQL
	logrus.Info("------------------ init mysql ------------------")
	if err := mysql.Init(); err != nil {
		panic(err)
	}

	// 3. 装配 shop 应用服务(内部装配 activity/order 依赖)
	shopSvc, err = buildShopService()
	if err != nil {
		panic(err)
	}
	if err := shopSvc.Start(); err != nil {
		panic(err)
	}

	// 4. 预热活动数据到 Redis(不阻塞启动,失败只打日志)
	go doPrewarm()

	// 5. 注册路由(依赖显式注入)
	handler.InitRouters(engine, handler.Dependencies{
		ShopService: shopSvc,
	})

	_ = logType
	return engine.RunListener(listener)
}

/*Exit
 *@Description: 优雅关闭
 */
func Exit() {
	_ = listener.Close()
	if shopSvc != nil {
		if err := shopSvc.Close(); err != nil {
			logrus.Errorf("logType:ApiExit, err:%s, step:close shop service", err.Error())
		}
	}
	if err := mysql.Close(); err != nil {
		logrus.Errorf("logType:ApiExit, err:%s, step:close mysql", err.Error())
	}
	time.Sleep(time.Second)
	logrus.Info("api server exit")
}

/*buildShopService
 *@Description: 构造 shop 应用服务,装配 queue + worker + stock + activity + order
 *@return *shopapp.Service
 *@return error
 */
func buildShopService() (*shopapp.Service, error) {
	// --- 消费端装配 ---
	q, err := factory.NewQueue(factory.DriverMemory, "shop")
	if err != nil {
		return nil, err
	}
	num := viper.GetInt(workerNumKey)
	if num <= 0 {
		num = runtime.NumCPU() * 4
	}
	size := viper.GetInt(workerSizeKey)
	if size <= 0 {
		size = defaultWorkerSize
	}
	workers := grab.NewWorker(num, size)
	logrus.Infof("logType:BuildShopService, msg:worker pool ready, workerNum:%d, workerSize:%d",
		num, size)

	// --- 库存工厂装配 ---
	redisFactory := stock.Factory(redisstock.NewRedisStock)
	memFactory := memstock.NewFactory(func(ctx context.Context, aid, gid string) (int64, error) {
		rs, err := redisFactory(aid, gid)
		if err != nil {
			return 0, err
		}
		return rs.Get(ctx)
	}).New

	// --- Activity / Order 仓储装配 ---
	cacheDB := viper.GetInt(cacheDBKey)
	if cacheDB <= 0 {
		cacheDB = defaultCacheDB
	}
	activityRepo := activityrepo.New(mysql.GetDB(), redis.GetRedisClient(cacheDB), 0)
	orderRepo := orderrepo.New(mysql.GetDB())

	// --- Service 装配 ---
	activitySvc := activityapp.NewService(activityRepo, redisFactory)
	orderSvc := orderapp.NewService(orderRepo)

	return shopapp.NewService(shopapp.Dependencies{
		NewMemStock:    memFactory,
		NewRedisStock:  redisFactory,
		Queue:          q,
		Workers:        workers,
		ActivityLookup: activitySvc,
		OrderCreator:   orderSvc,
	})
}

/*doPrewarm
 *@Description: 单独 goroutine 跑预热,失败不阻塞 gin 启动
 */
func doPrewarm() {
	logType := "ApiPrewarm"
	cacheDB := viper.GetInt(cacheDBKey)
	if cacheDB <= 0 {
		cacheDB = defaultCacheDB
	}
	activityRepo := activityrepo.New(mysql.GetDB(), redis.GetRedisClient(cacheDB), 0)
	activitySvc := activityapp.NewService(activityRepo, stock.Factory(redisstock.NewRedisStock))

	ctx, cancel := context.WithTimeout(context.Background(), prewarmTimeout)
	defer cancel()
	if err := activitySvc.Prewarm(ctx); err != nil {
		logrus.Errorf("logType:%s, err:%s, msg:prewarm failed, continue serving", logType, err.Error())
	}
}
