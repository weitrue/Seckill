/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午5:37
 * Description:
 **/

package admin

import (
	"encoding/json"
	"net"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/weitrue/Seckill/infrastructure/config"
	"github.com/weitrue/Seckill/infrastructure/config/cluster"
	"github.com/weitrue/Seckill/interfaces/admin/handler"
	utils2 "github.com/weitrue/Seckill/pkg/utils"
)

var (
	listener net.Listener
	err      error
)

func Run() error {
	bind := viper.GetString("admin.bind")
	logrus.Info("run admin server on ", bind)
	listener, err = utils2.Listen(config.GetTcpNet(), bind)
	if err != nil {
		return err
	}

	engine := gin.New()
	// TODO 更新程序， 给老版本发信号 v1.0
	utils2.UpdateProcess("admin")

	// 初始化路由
	handler.InitRouters(engine)

	// TODO 加入集群
	cluster.Init(config.GetServiceName())
	if nodes, err := cluster.Discover(); err == nil {
		n, _ := json.Marshal(nodes)
		logrus.Info("discover nodes ", string(n))
	} else {
		logrus.Error("discover nodes error: ", err)
	}

	return engine.RunListener(listener)
}

func Exit() {
	_ = listener.Close()
	// TODO等待请求处理完
	time.Sleep(time.Second)
	logrus.Info("admin server exit")
}
