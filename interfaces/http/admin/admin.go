/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午5:37
 * Description: admin HTTP 服务启动入口
 *   - 初始化 Redis + MySQL
 *   - 装配 activity.AdminService(含 Repository + 库存工厂)
 *   - 注册路由,传入 Dependencies
 **/

package admin

import (
	"encoding/json"
	"net"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/weitrue/Seckill/interfaces/http/admin/handler"
	activityapp "github.com/weitrue/Seckill/internal/application/activity"
	"github.com/weitrue/Seckill/internal/domain/stock"
	"github.com/weitrue/Seckill/internal/infrastructure/config"
	"github.com/weitrue/Seckill/internal/infrastructure/config/cluster"
	"github.com/weitrue/Seckill/internal/infrastructure/persistence/activityrepo"
	"github.com/weitrue/Seckill/internal/infrastructure/persistence/stockrepo/redisstock"
	"github.com/weitrue/Seckill/internal/infrastructure/stores/mysql"
	"github.com/weitrue/Seckill/internal/infrastructure/stores/redis"
	utils2 "github.com/weitrue/Seckill/pkg/utils"
)

const (
	cacheDBKey     = "redis.DB.cache"
	defaultCacheDB = 11
)

var (
	listener net.Listener
	err      error
)

/*Run
 *@Description: 启动 admin HTTP 服务
 *@return error
 */
func Run() error {
	bind := viper.GetString("admin.bind")
	logrus.Info("run admin server on ", bind)
	listener, err = utils2.Listen(config.GetTcpNet(), bind)
	if err != nil {
		return err
	}

	engine := gin.New()
	utils2.UpdateProcess("admin")

	// 1. 基础设施
	if err := redis.Init(); err != nil {
		panic(err)
	}
	if err := mysql.Init(); err != nil {
		panic(err)
	}

	// 2. 装配 admin 的 Activity 应用服务
	cacheDB := viper.GetInt(cacheDBKey)
	if cacheDB <= 0 {
		cacheDB = defaultCacheDB
	}
	activityRepo := activityrepo.New(mysql.GetDB(), redis.GetRedisClient(cacheDB), 0)
	activityAdminSvc := activityapp.NewAdminService(activityRepo, stock.Factory(redisstock.NewRedisStock))

	// 3. 注册路由
	handler.InitRouters(engine, handler.Dependencies{
		ActivityAdminService: activityAdminSvc,
	})

	// 4. 加入集群
	cluster.Init(config.GetServiceName())
	if nodes, err := cluster.Discover(); err == nil {
		n, _ := json.Marshal(nodes)
		logrus.Info("discover nodes ", string(n))
	} else {
		logrus.Error("discover nodes error: ", err)
	}

	return engine.RunListener(listener)
}

/*Exit
 *@Description: 优雅关闭
 */
func Exit() {
	_ = listener.Close()
	if err := mysql.Close(); err != nil {
		logrus.Errorf("logType:AdminExit, err:%s, step:close mysql", err.Error())
	}
	time.Sleep(time.Second)
	logrus.Info("admin server exit")
}
