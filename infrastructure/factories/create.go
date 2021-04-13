/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午8:23
 * Description:
 **/

package factories

import "Seckill/infrastructure/mq/memory"

var factoryMap = make(map[string]Factory)

func NewFactory(fm string) Factory {
	// 根据工厂类型参数创建对应的工厂
	return factoryMap[fm]
}

func Register(fm string, f Factory) {
	// 将不同的工厂类注册到统一的map中
	if _, ok := factoryMap[fm]; ok {
		panic("duplicate queue factories " + fm)
	}
	factoryMap[fm] = f
}

func init() {
	Register("memory", FactoryFunc(memory.MQFactory))
}
