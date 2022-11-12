/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/12/23 3:24 下午
 * Description:
 **/

package int

import (
	"sync/atomic"

	"github.com/weitrue/Seckill/infrastructure/stores/cache"
)

type integerCache struct {
	data atomic.Value
}

func NewIntegerCache() cache.IntCache {
	ic := &integerCache{}
	ic.data.Store(make(map[string]*int64))
	return ic
}


func (i *integerCache) Del(key string) {
	panic("implement me")
}

func (i *integerCache) Keys() []string {
	panic("implement me")
}

func (i *integerCache) getPtr(key string) *int64 {
	valPtr, _ := i.data.Load().(map[string]*int64)[key]
	return valPtr
}

func (i *integerCache) Get(key string) (int64, bool) {
	if valPtr := i.getPtr(key); valPtr != nil {
		return *valPtr, true
	} else {
		return 0, false
	}
}

func (i *integerCache) Set(key string, val int64) {
	oldMap := i.data.Load().(map[string]*int64)
	newMap := make(map[string]*int64, len(oldMap)+1)
	for k, v := range oldMap {
		newMap[k] = v
	}
	newMap[key] = &val
	i.data.Store(newMap)
}

func (i *integerCache) Add(key string, delta int64) int64 {
	panic("implement me")
}
