/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午2:57
 * Description: rpc服务实现
 **/

package api

import (
	"Seckill/application/api/rpc"
	"context"

	"github.com/sirupsen/logrus"
)

type ActivityRPCServer struct {
	// rpc服务
}

func (s *ActivityRPCServer) ActivityOnline(ctx context.Context, in *rpc.Activity) (*rpc.Response, error) {
	logrus.Info("ActivityOnline", in)
	response := &rpc.Response{}
	return response, nil
}

func (s *ActivityRPCServer) ActivityOffline(ctx context.Context, in *rpc.Activity) (*rpc.Response, error) {
	logrus.Info("ActivityOffline", in)
	response := &rpc.Response{}
	return response, nil
}

func (s *ActivityRPCServer) TopicOnline(ctx context.Context, in *rpc.Topic) (*rpc.Response, error) {
	logrus.Info("TopicOnline", in)
	response := &rpc.Response{}
	return response, nil
}

func (s *ActivityRPCServer) TopicOffline(ctx context.Context, in *rpc.Topic) (*rpc.Response, error) {
	logrus.Info("TopicOffline", in)
	response := &rpc.Response{}
	return response, nil
}
