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

type IntCache interface {
	Cache
	// Get 缓存中value为int 黑白名单 库存信息用户抢购次数等
	Get(key string) (int64, bool)
	Set(key string, val int64)
	Add(key string, delta int64) int64
}

type ObjCache interface {
	Cache
	// Get 缓存中value为结构体 活动信息
	Get(key string) (interface{}, bool)
	Set(key string, val interface{})
}