/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午6:35
 * Description: mq 顶层接口定义
 *   - MQ:       所有队列实现的共同能力(Close)
 *   - Queue:    本地任务队列,Produce/Consume worker.Task(如 memory)
 *   - PubSub:   第三方发布订阅队列,Publish/Subscribe []byte(如 kafka / rabbitmq)
 *   具体实现放在同级子包: memory / kafka / rabbitmq
 *   工厂(统一创建入口)放在 factory 子包
 **/

package mq

import (
	"io"

	"github.com/weitrue/Seckill/infrastructure/worker"
)

// MQ 所有队列实现的共同能力
type MQ interface {
	io.Closer
}

// Queue 本地任务队列: 进程内 Produce/Consume worker.Task
type Queue interface {
	MQ
	Produce(task worker.Task) error
	Consume() (worker.Task, error)
}

// PubSub 第三方发布订阅队列: Publish/Subscribe []byte
type PubSub interface {
	MQ
	// Publish 发布消息
	//   topic : 主题/路由键
	//   body  : 消息体
	//   reqID : 请求 ID,用于链路追踪
	Publish(topic string, body []byte, reqID string) error
	// Subscribe 订阅消息
	//   topic : 主题/路由键
	//   group : 消费组/队列名,同一 group 下的多实例分摊消费
	//   fn    : 消息回调,返回 nil 表示处理成功
	Subscribe(topic, group string, fn MsgCb) error
}

// MsgCb 消费回调
type MsgCb func(m ConsumerMsg) error

// ConsumerMsg 消费到的消息抽象
type ConsumerMsg interface {
	GetBody() []byte // 消息体
	GetID() string   // 消息 ID / 请求 ID
	Ack() error      // 手动 ack
}
