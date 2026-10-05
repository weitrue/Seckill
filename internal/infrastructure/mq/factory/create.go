/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午8:23
 * Description: 队列工厂注册中心 + 内置驱动注册
 *   支持的 driver:
 *     - memory   : 本地任务队列(mq.Queue)
 *     - kafka    : kafka 发布订阅(mq.PubSub)
 *     - rabbitmq : rabbitmq 发布订阅(mq.PubSub)
 **/

package factory

import (
	"fmt"
	"sync"

	"github.com/weitrue/Seckill/internal/infrastructure/mq"
	"github.com/weitrue/Seckill/internal/infrastructure/mq/kafka"
	"github.com/weitrue/Seckill/internal/infrastructure/mq/memory"
	"github.com/weitrue/Seckill/internal/infrastructure/mq/rabbitmq"
)

// 驱动类型常量
const (
	DriverMemory   = "memory"
	DriverKafka    = "kafka"
	DriverRabbitMQ = "rabbitmq"
)

var (
	driversMu sync.RWMutex
	drivers   = make(map[string]Factory)
)

/*NewFactory
 *@Description: 根据驱动名获取工厂,未注册返回 nil
 *@param driver 驱动名(memory/kafka/rabbitmq ...)
 *@return Factory
 */
func NewFactory(driver string) Factory {
	driversMu.RLock()
	defer driversMu.RUnlock()
	return drivers[driver]
}

/*Register
 *@Description: 注册一个驱动工厂,重复注册或传 nil 会 panic(初始化期行为)
 *@param driver 驱动名
 *@param f 工厂实现
 */
func Register(driver string, f Factory) {
	if f == nil {
		panic("mq/factory: Register factory is nil, driver=" + driver)
	}
	driversMu.Lock()
	defer driversMu.Unlock()
	if _, ok := drivers[driver]; ok {
		panic("mq/factory: Duplicate queue factory, driver=" + driver)
	}
	drivers[driver] = f
}

/*NewQueue
 *@Description: 按驱动创建本地任务队列(mq.Queue);不是本地任务队列类型则返回错误
 *@param driver 驱动名
 *@param name 配置 key
 *@return mq.Queue
 *@return error
 */
func NewQueue(driver, name string) (mq.Queue, error) {
	f := NewFactory(driver)
	if f == nil {
		return nil, fmt.Errorf("mq/factory: driver not registered, driver=%s", driver)
	}
	instance, err := f.New(name)
	if err != nil {
		return nil, err
	}
	q, ok := instance.(mq.Queue)
	if !ok {
		return nil, fmt.Errorf("mq/factory: driver %q does not implement mq.Queue", driver)
	}
	return q, nil
}

/*NewPubSub
 *@Description: 按驱动创建发布订阅队列(mq.PubSub);不是发布订阅类型则返回错误
 *@param driver 驱动名
 *@param name 配置 key
 *@return mq.PubSub
 *@return error
 */
func NewPubSub(driver, name string) (mq.PubSub, error) {
	f := NewFactory(driver)
	if f == nil {
		return nil, fmt.Errorf("mq/factory: driver not registered, driver=%s", driver)
	}
	instance, err := f.New(name)
	if err != nil {
		return nil, err
	}
	ps, ok := instance.(mq.PubSub)
	if !ok {
		return nil, fmt.Errorf("mq/factory: driver %q does not implement mq.PubSub", driver)
	}
	return ps, nil
}

func init() {
	// 本地任务队列: Fan-In 漏桶限流,供 memory 任务排队使用
	Register(DriverMemory, Func(func(name string) (mq.MQ, error) {
		return memory.NewMemoryMQ(name)
	}))
	// kafka pub/sub
	Register(DriverKafka, Func(func(name string) (mq.MQ, error) {
		return kafka.NewKafkaQueue(name)
	}))
	// rabbitmq pub/sub
	Register(DriverRabbitMQ, Func(func(name string) (mq.MQ, error) {
		return rabbitmq.NewRabbitMQ(name)
	}))
}
