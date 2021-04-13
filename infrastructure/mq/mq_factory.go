/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午8:23
 * Description:
 **/

package mq

import "Seckill/infrastructure/mq/mqfatory"

var queueFactories = make(map[string]mqfatory.Factory)

func NewFactory(tp string) mqfatory.Factory {
	// 根据工厂类型参数创建对应的工厂
	return queueFactories[tp]
}

func Register(tp string, f mqfatory.Factory) {
	// 将不同的工厂类注册
	if _, ok := queueFactories[tp]; ok {
		panic("duplicate queue factory " + tp)
	}
	queueFactories[tp] = f
}
