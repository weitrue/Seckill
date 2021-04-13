/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午7:35
 * Description: 工厂方法
 **/

package mqfatory

import "Seckill/infrastructure/mq"

type FactoryFunc func(name string) (mq.Queue, error)

func (f FactoryFunc) New(name string) (mq.Queue, error) {
	return f(name)
}

func (f FactoryFunc) NewProducer(name string) (mq.Producer, error) {
	return f.New(name)
}

func (f FactoryFunc) NewConsumer(name string) (mq.Consumer, error) {
	return f.New(name)
}
