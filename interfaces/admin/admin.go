/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午5:37
 * Description:
 **/

package admin

import (
	"Seckill/infrastructure/utils"
	"net"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	listener net.Listener
	err      error
)

func Run() error {
	bind := viper.GetString("admin.bind")
	logrus.Info("run admin server on ", bind)
	listener, err = utils.Listen("tcp", bind)
	if err != nil {
		return err
	}

	engine := gin.New()
	// TODO 更新程序， 给老版本发信号

	//初始化路由
	InitRouters(engine)

	// TODO 加入集群

	return engine.RunListener(listener)
}

func Exit()  {
	_ = listener.Close()
	// TODO等待请求处理完
	time.Sleep(time.Second)
	logrus.Info("admin server exit")
}

