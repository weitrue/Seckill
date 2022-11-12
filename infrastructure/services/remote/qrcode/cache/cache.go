package cache

import (
	"sync"
)

type Cache interface {
	Del(key string)
	Get(key string) (interface{}, bool)
	Set(key string, val interface{})
}

type objCache struct {
	sync.RWMutex
	data map[string]interface{}
}

func NewObjCache() Cache {
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
