/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午7:31
 * Description: 抽象多种类型的队列: 队列
 **/

package factory

import (
	"github.com/weitrue/Seckill/infrastructure/mq/mqi"
)

type Factory interface {
	New(name string) (mqi.Queue, error)
	NewProducer(name string) (mqi.Producer, error)
	NewConsumer(name string) (mqi.Consumer, error)
}

type FactoryFunc func(name string) (mqi.Queue, error)

func (f FactoryFunc) New(name string) (mqi.Queue, error) {
	return f(name)
}

func (f FactoryFunc) NewProducer(name string) (mqi.Producer, error) {
	return f.New(name)
}

func (f FactoryFunc) NewConsumer(name string) (mqi.Consumer, error) {
	return f.New(name)
}
