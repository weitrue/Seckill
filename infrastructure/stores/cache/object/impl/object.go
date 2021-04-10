/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 下午3:56
 * Description: 缓存中value为结构体 活动信息
 **/

package impl

import (
	"Seckill/infrastructure/stores/cache/object"
	"sync"
)

type objCache struct {
	sync.RWMutex
	data map[string]interface{}
}

func NewObjCache() object.ObjCache {
	return &objCache{
		data: make(map[string]interface{}),
	}
}

func (c *objCache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	val, ok := c.data[key]
	return val, ok
}

func (c *objCache) Set(key string, val interface{}) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = val
}

func (c *objCache) Del(key string) {
	if _, ok := c.Get(key); ok {
		c.Lock()
		defer c.Unlock()
		delete(c.data, key)
	}
}

func (c *objCache) Keys() []string {
	keys := make([]string, 0)

	for k, _ := range c.data {
		keys = append(keys, k)
	}
	return keys
}
