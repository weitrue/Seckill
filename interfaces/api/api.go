/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午5:04
 * Description:
 **/

package api

import (
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
	listener, err = utils.Listen("tcp", bind)
	if err != nil {
		return err
	}
	engine := gin.New()
	// TODO 更新程序， 给老版本发信号

	// TODO redis初始化

	// TODO 黑名单监控

	// 初始化路由
	initRouters(engine)

	// TODO 初始化秒杀服务
	//shop.Init()
	pprof.Register(engine)

	// TODO 加入集群

	// 运行服务
	return engine.RunListener(listener)
}

func Exit() {
	_ = listener.Close()
	// TODO等待请求处理完
	time.Sleep(time.Second)
	logrus.Info("api server exit")
}