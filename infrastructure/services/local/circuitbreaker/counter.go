/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午4:18
 * Description: 熔断器 - 计数器
				基本思路是：
				 划分一个时间窗口，如 100 毫秒，并设置判断条件，如失败率超过 5%、请求数超过 1000 等；
				 在请求进来的时候，判断条件是否满足熔断条件，如果满足就拒绝请求，如果不满足就继续处理请求；
				 请求处理完后，统计时间窗口内请求失败率、延迟不达标率、请求数等指标，以便作为后续请求的判断条件。
 **/

package circuitbreaker

import "sync/atomic"

type Counter int64

func (c *Counter) Add() int64 {
	// 自增 原子加1操作
	return atomic.AddInt64((*int64)(c), 1)
}

func (c *Counter) Load() int64 {
	// 获取当前计数
	return atomic.LoadInt64((*int64)(c))
}

func (c *Counter) Reset() {
	// 重置
	atomic.StoreInt64((*int64)(c), 0)
}
