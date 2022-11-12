/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午6:35
 * Description:
 **/

package mq

import (
	"errors"
	"io"

	"github.com/weitrue/Seckill/infrastructure/mq/kafka"
	"github.com/weitrue/Seckill/infrastructure/mq/memory"
	"github.com/weitrue/Seckill/infrastructure/worker"
)

type Queue interface {
	LocalQueue
	io.Closer
}

type LocalQueue interface {
	Producer
	Consumer
}

type Producer interface { // 生产者
	Produce(task worker.Task) error
}

type Consumer interface { // 消费者
	Consume() (worker.Task, error)
}

// ThirdQueue 第三方mq
type ThirdQueue interface {
	ThirdProducer
	ThirdConsumer
}

type ThirdProducer interface {
	Publish(topic string, body []byte, reqID string) error
}

type ThirdConsumer interface {
	Subscribe(topic, queueName string, fn kafka.MsgCb) error
}

type messageQueue struct {
	queue Queue
}

func NewMQ(mqType string) (Queue, error) {
	var queue Queue
	var err error

	switch mqType {
	case "memory":
		queue, err = memory.NewMemoryMQ("")
	case "kafka":
		return nil, errors.New("unknown mq type")
		// queue, err = kafka.NewKafkaQueue("")
	default:
		return nil, errors.New("unknown mq type")
	}

	return &messageQueue{queue: queue}, err
}

func (m *messageQueue) Produce(task worker.Task) error {
	return m.queue.Produce(task)
}

func (m *messageQueue) Consume() (worker.Task, error) {
	return m.queue.Consume()
}

func (m *messageQueue) Close() error {
	return m.queue.Close()
}
