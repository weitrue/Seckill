/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午7:31
 * Description: 队列工厂接口定义
 *   统一返回 mq.MQ,具体类型(mq.Queue / mq.PubSub)由调用方按需断言,
 *   或通过 NewQueue / NewPubSub 辅助函数直接拿到对应接口
 **/

package factory

import (
	"github.com/weitrue/Seckill/infrastructure/mq"
)

// Factory 队列工厂抽象
type Factory interface {
	// New 按 name 创建一个队列实例;name 作为配置 key(queue.<name>.*)
	New(name string) (mq.MQ, error)
}

// Func 函数适配器,方便直接用普通函数注册为 Factory
type Func func(name string) (mq.MQ, error)

/*New
 *@Description: Func 适配 Factory 接口
 *@receiver f
 *@param name 配置 key
 *@return mq.MQ
 *@return error
 */
func (f Func) New(name string) (mq.MQ, error) {
	return f(name)
}
