/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午5:04
 * Description:
 **/

package api

import (
	"Seckill/domain/shop"
	"Seckill/infrastructure/config"
	"Seckill/infrastructure/stores/redis"
	"Seckill/infrastructure/utils"
	"net"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	listener net.Listener
	err      error
)

func Run() error {
	bind := viper.GetString("api.bind")
	logrus.Info("run api server on ", bind)
	listener, err = utils.Listen(utils.GetTcpNet(), bind)
	if err != nil {
		return err
	}
	engine := gin.New()
	// TODO 更新程序， 给老版本发信号 Done
	utils.UpdateProcess("api")

	// TODO 黑名单监控 Done
	config.WatchBlacklistConfig()
	// 初始化路由
	initRouters(engine)

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
