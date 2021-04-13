/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午7:31
 * Description:
 **/

package mqfatory

import "Seckill/infrastructure/mq"

type Factory interface {
	New(name string) (mq.Queue, error)
	NewProducer(name string) (mq.Producer, error)
	NewConsumer(name string) (mq.Consumer, error)
}
