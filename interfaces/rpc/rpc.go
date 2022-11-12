/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午3:49
 * Description:
 **/

package rpc

import (
	"sync"

	"github.com/weitrue/Seckill/application/api"
	"github.com/weitrue/Seckill/application/api/rpc"
	"github.com/weitrue/Seckill/infrastructure/config"
	"github.com/weitrue/Seckill/infrastructure/config/cluster"
	utils2 "github.com/weitrue/Seckill/pkg/utils"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	grpcS *grpc.Server
	node  *cluster.Node // 集群中服务的节点
	once  = &sync.Once{}
)

func Run() error {
	bind := viper.GetString("api.rpc")
	logrus.Info("run RPC Server on ", bind)
	listen, err := utils2.Listen(config.GetTcpNet(), bind)
	if err != nil {
		return err
	}
	grpcS = grpc.NewServer()
	activityRPC := &api.ActivityRPCServer{}
	rpc.RegisterActivityRPCServer(grpcS, activityRPC)
	// 支持 gRPC reflection，方便调试
	reflection.Register(grpcS)
	// 初始化集群信息
	cluster.Init(config.GetServiceName())
	var addr string
	if addr, err = utils2.Extract(bind); err == nil {
		// 注册节点信息
		version := viper.GetString("api.version")
		if version == "" {
			version = "v0.1"
		}
		once.Do(func() {
			node = &cluster.Node{
				Addr:    addr,
				Version: version,
				Proto:   viper.GetString(config.GetGRPCNet()),
			}
			err = cluster.Register(node, viper.GetInt("api.ttl"))
		})
	}
	if err != nil {
		return err
	}

	return grpcS.Serve(listen)
}

func Exit() {
	grpcS.GracefulStop()
	logrus.Info("rpc server exit!")
}
