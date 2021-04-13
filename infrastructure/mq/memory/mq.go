/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午6:34
 * Description: 请求排队的队列有多个生产者，但只有一个消费者，且以固定速度消费，因此基于 Fan-In 模式的 RateLimiter 来实现内存队列
 **/

package memory

import (
	"Seckill/infrastructure/factories"
	"Seckill/infrastructure/mq"
	"Seckill/infrastructure/pool/coroutine"
	"Seckill/infrastructure/services/local/ratelimiter"
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type memoryQueue struct {
	queue ratelimiter.RateLimiter
}

func memoryQueueFactory(name string) (mq.Queue, error) {
	rate := viper.GetInt64(fmt.Sprintf("queue.%s.rate", name))
	size := viper.GetInt64(fmt.Sprintf("queue.%s.size", name))
	q, _ := ratelimiter.NewRateLimiter(size, rate, ratelimiter.FanIn)
	return &memoryQueue{queue: q}, nil
}

func (mq *memoryQueue) Produce(task coroutine.Task) error {
	if ok := mq.queue.Push(task); !ok {
		return errors.New("queue producer error")
	}
	return nil
}

func (mq *memoryQueue) Consume() (coroutine.Task, error) {
	task, ok := mq.queue.Pop()
	if !ok {
		return nil, errors.New("queue consumer error")
	}
	return task, nil
}

func (mq *memoryQueue) Close() error {
	return mq.queue.Close()
}

func init() {
	factories.Register("memory", factories.FactoryFunc(memoryQueueFactory))
}
