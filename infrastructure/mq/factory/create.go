/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午8:23
 * Description:
 **/

package factory

import (
	"github.com/weitrue/Seckill/infrastructure/mq"
)

var drivers = make(map[string]Factory)

func NewFactory(fm string) Factory {
	// 根据工厂类型参数创建对应的工厂
	return drivers[fm]
}

func Register(fm string, f Factory) {
	if f == nil {
		panic("Register factory is nil")
	}

	// 将不同的工厂类注册到统一的map中
	if _, ok := drivers[fm]; ok {
		panic("Duplicate queue factory " + fm)
	}

	drivers[fm] = f
}

func init() {
	// 注册采用Fan-In(漏桶限流)限流器 限制redis的消费
	Register("memory", Func(mq.NewMQ))
	Register("kafka", Func(mq.NewMQ))
}
