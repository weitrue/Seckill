/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午2:57
 * Description: gRPC 服务端实现(协议细节)
 *   现阶段接口是空壳,仅返回 ok;等 application.activity.Service 落地后在此做协议转换 + 调 application
 **/

package rpc

import (
	"context"

	"github.com/sirupsen/logrus"

	pb "github.com/weitrue/Seckill/interfaces/grpc/pb"
)

// ActivityRPCServer 活动 / 专题的 gRPC server 实现
type ActivityRPCServer struct {
	// TODO: 阶段 2 之后注入 application.activity.Service / topic.Service
}

/*ActivityOnline
 *@Description: 活动上线
 */
func (s *ActivityRPCServer) ActivityOnline(ctx context.Context, in *pb.Activity) (*pb.Response, error) {
	logrus.Infof("logType:RpcActivityOnline, in:%+v", in)
	return &pb.Response{}, nil
}

/*ActivityOffline
 *@Description: 活动下线
 */
func (s *ActivityRPCServer) ActivityOffline(ctx context.Context, in *pb.Activity) (*pb.Response, error) {
	logrus.Infof("logType:RpcActivityOffline, in:%+v", in)
	return &pb.Response{}, nil
}

/*TopicOnline
 *@Description: 专题上线
 */
func (s *ActivityRPCServer) TopicOnline(ctx context.Context, in *pb.Topic) (*pb.Response, error) {
	logrus.Infof("logType:RpcTopicOnline, in:%+v", in)
	return &pb.Response{}, nil
}

/*TopicOffline
 *@Description: 专题下线
 */
func (s *ActivityRPCServer) TopicOffline(ctx context.Context, in *pb.Topic) (*pb.Response, error) {
	logrus.Infof("logType:RpcTopicOffline, in:%+v", in)
	return &pb.Response{}, nil
}
