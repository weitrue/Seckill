/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 下午3:55
 * Description: 缓存中value为int 黑白名单 库存信息用户抢购次数等
 **/

package impl

import (
	"Seckill/infrastructure/stores/cache/integer"
	"sync"
	"sync/atomic"
)

type intCache struct {
	sync.RWMutex
	data map[string]*int64 // 整数指针，方便原子操作
}

func NewIntCache() integer.IntCache {
	return &intCache{
		data: make(map[string]*int64),
	}
}

func (c *intCache) getPtr(key string) *int64 {
	c.RLock()
	defer c.RUnlock()
	valPtr, _ := c.data[key]
	return valPtr
}

func (c *intCache) Get(key string) (int64, bool) {

	if valPtr := c.getPtr(key); valPtr != nil {
		return atomic.LoadInt64(valPtr), true
	} else {
		return 0, false
	}

}

func (c *intCache) Set(key string, val int64) {

	if valPtr := c.getPtr(key); valPtr != nil {
		// 指针存在，使用原子操作
		atomic.StoreInt64(valPtr, val)
	} else {
		valPtr = new(int64)
		*valPtr = val
		c.Lock()
		defer c.Unlock()
		c.data[key] = valPtr
	}
}

func (c *intCache) Add(key string, delta int64) int64 {
	if valPtr := c.getPtr(key); valPtr != nil {
		return atomic.AddInt64(valPtr, delta)
	} else {
		var val int64
		var ok bool
		c.Lock()
		defer c.Unlock()
		if valPtr, ok = c.data[key]; ok {
			val = atomic.AddInt64(valPtr, delta)
		} else {
			val = delta
			valPtr = &val
			c.data[key] = valPtr
		}
		return val
	}
}

func (c *intCache) Del(key string) {
	if valPtr := c.getPtr(key); valPtr != nil {
		c.Lock()
		defer c.Unlock()
		delete(c.data, key)
	}
}

func (c *intCache) Keys() []string {
	keys := make([]string, 0)
	c.RLock()
	defer c.RUnlock()
	for k, _ := range c.data {
		keys = append(keys, k)
	}
	return keys
}
