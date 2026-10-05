/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午6:34
 * Description: 请求排队的本地队列: 多生产者 / 单消费者 / 固定速度消费,基于 Fan-In 模式的 RateLimiter 实现
 *   实现 mq.Queue 接口(Produce/Consume worker.Task + Close)
 **/

package memory

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"

	"github.com/weitrue/Seckill/internal/infrastructure/mq"
	"github.com/weitrue/Seckill/internal/infrastructure/services/local/ratelimiter"
	"github.com/weitrue/Seckill/internal/infrastructure/worker"
)

// memoryQueue 本地任务队列
type memoryQueue struct {
	queue ratelimiter.RateLimiter
}

// 编译期类型断言: 确保 memoryQueue 实现了 mq.Queue
var _ mq.Queue = (*memoryQueue)(nil)

/*Produce
 *@Description: 投递任务到本地队列
 *@receiver m
 *@param task 任务
 *@return error 失败原因
 */
func (m *memoryQueue) Produce(task worker.Task) error {
	if ok := m.queue.Push(task); !ok {
		return errors.New("memory queue produce failed")
	}

	return nil
}

/*Consume
 *@Description: 从本地队列取出一个任务
 *@receiver m
 *@return worker.Task 任务
 *@return error 失败原因
 */
func (m *memoryQueue) Consume() (worker.Task, error) {
	t, ok := m.queue.Pop()
	if !ok {
		return nil, errors.New("memory queue consume failed")
	}

	return t, nil
}

/*Close
 *@Description: 关闭本地队列
 *@receiver m
 *@return error
 */
func (m *memoryQueue) Close() error {
	return m.queue.Close()
}

/*NewMemoryMQ
 *@Description: 创建本地队列实例
 *   从 viper 读取 queue.<name>.rate / queue.<name>.size 作为限流器参数
 *@param name 队列名(对应配置中的 key)
 *@return mq.Queue 本地队列
 *@return error
 */
func NewMemoryMQ(name string) (mq.Queue, error) {
	rate := viper.GetInt64(fmt.Sprintf("queue.%s.rate", name))
	size := viper.GetInt64(fmt.Sprintf("queue.%s.size", name))
	q, err := ratelimiter.NewRateLimiter(size, rate, ratelimiter.FanIn)
	if err != nil {
		return nil, err
	}

	return &memoryQueue{queue: q}, nil
}
