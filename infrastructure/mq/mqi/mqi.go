/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午6:35
 * Description:
 **/

package mqi

import (
	"io"

	"github.com/weitrue/Seckill/infrastructure/pool/worker/taski"
)

type Queue interface {
	Producer
	Consumer
	io.Closer
}
type Producer interface { // 生产者
	Produce(task taski.Task) error
}
type Consumer interface { // 消费者
	Consume() (taski.Task, error)
}
