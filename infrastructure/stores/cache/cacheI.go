/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 下午3:24
 * Description:内存缓存分为两部分数据：访问量大，变更少的数据（黑名单、活动信息等），实现高性能中间层流量拦截（抢购次数数据、库存信息等）
 **/

package cache

type Cache interface {
	Del(key string)
	Keys() []string
}
