/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午5:04
 * Description:
 **/

package api

import (
	"net"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/weitrue/Seckill/domain/shop"
	"github.com/weitrue/Seckill/infrastructure/config"
	"github.com/weitrue/Seckill/infrastructure/stores/redis"
	"github.com/weitrue/Seckill/interfaces/api/handler"
	utils2 "github.com/weitrue/Seckill/pkg/utils"
)

var (
	listener net.Listener
	err      error
)

func Run() error {
	bind := viper.GetString("api.bind")
	logrus.Info("run api server on ", bind)
	listener, err = utils2.Listen(config.GetTcpNet(), bind)
	if err != nil {
		return err
	}
	engine := gin.New()
	// TODO 更新程序， 给老版本发信号 Done
	utils2.UpdateProcess("api")

	// TODO 黑名单监控 Done
	config.WatchBlacklistConfig()
	// 初始化路由
	handler.InitRouters(engine)
	engine.Use(gin.Logger())

	logrus.Info("------------------ init redis ------------------")
	// TODO redis初始化 v1.0
	if err := redis.Init(); err != nil {
		panic(err)
	}
	pprof.Register(engine)
	// TODO 初始化秒杀服务
	shop.Init()
	// 运行服务
	return engine.RunListener(listener)
}

func Exit() {
	_ = listener.Close()
	// TODO等待请求处理完
	time.Sleep(time.Second)
	logrus.Info("api server exit")
}
