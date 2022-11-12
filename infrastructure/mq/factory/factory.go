/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午7:31
 * Description: 抽象多种类型的队列: 队列
 **/

package factory

import (
	"github.com/weitrue/Seckill/infrastructure/mq"
)

type Factory interface {
	New(name string) (mq.Queue, error)
}

type Func func(name string) (mq.Queue, error)

func (f Func) New(name string) (mq.Queue, error) {
	return f(name)
}
