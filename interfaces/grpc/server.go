/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午3:49
 * Description: gRPC 服务启动入口
 *   - 启动 gRPC server 并注册所有 service(当前只有 ActivityRPCServer)
 *   - 注册节点到 etcd 做服务发现
 **/

package rpc

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/weitrue/Seckill/interfaces/grpc/pb"
	"github.com/weitrue/Seckill/internal/infrastructure/config"
	"github.com/weitrue/Seckill/internal/infrastructure/config/cluster"
	utils2 "github.com/weitrue/Seckill/pkg/utils"
)

var (
	grpcS *grpc.Server
	node  *cluster.Node // 集群中服务的节点
	once  = &sync.Once{}
)

/*Run
 *@Description: 启动 gRPC server
 *@return error
 */
func Run() error {
	bind := viper.GetString("api.rpc")
	logrus.Info("run RPC Server on ", bind)
	listen, err := utils2.Listen(config.GetTcpNet(), bind)
	if err != nil {
		return err
	}
	grpcS = grpc.NewServer()
	pb.RegisterActivityRPCServer(grpcS, &ActivityRPCServer{})
	// 支持 gRPC reflection,方便调试
	reflection.Register(grpcS)

	// 初始化集群信息并注册节点
	cluster.Init(config.GetServiceName())
	var addr string
	if addr, err = utils2.Extract(bind); err == nil {
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

/*Exit
 *@Description: 优雅关闭
 */
func Exit() {
	if grpcS != nil {
		grpcS.GracefulStop()
	}
	logrus.Info("rpc server exit!")
}
