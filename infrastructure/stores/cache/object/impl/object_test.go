/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 下午6:01
 * Description:
 **/

package impl

import (
	"Seckill/infrastructure/stores/cache/object"
	"os"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObjCache(t *testing.T) {
	c := NewObjCache()
	a := assert.New(t)
	key := "test"
	c.Set(key, int64(1))

	v, ok := c.Get(key)
	a.Equal(int64(1), v)
	a.Equal(true, ok)

	c.Del(key)
	_, ok = c.Get(key)
	a.Equal(false, ok)
}

func initKeys(b *testing.B) []string {
	var keys = make([]string, 0)
	maxKeyStr := os.Getenv("maxKey")
	maxKey, _ := strconv.Atoi(maxKeyStr)
	if maxKey <= 0 {
		maxKey = b.N
	}
	for i := 0; i < maxKey; i++ {
		keys = append(keys, strconv.Itoa(i))
	}
	return keys
}

func initSyncMap(b *testing.B, c sync.Map, keys []string) {
	l := len(keys)
	for i := 0; i < b.N; i++ {
		c.Store(keys[i%l], int64(i))
	}
}

func initObjCache(b *testing.B, c object.ObjCache, keys []string) {
	l := len(keys)
	for i := 0; i < b.N; i++ {
		c.Set(keys[i%l], int64(i))
	}
}

func benchmarkCacheSet(b *testing.B, setter func(key string, val int64), keys []string) {
	b.ReportAllocs()
	b.StartTimer()
	l := len(keys)
	for i := 0; i < b.N; i++ {
		setter(keys[i%l], int64(i))
	}
	b.StopTimer()
}

func BenchmarkCache_Set(b *testing.B) {
	keys := make([]string, b.N, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = strconv.Itoa(i)
	}
	b.ResetTimer()

	b.Run("objCache", func(b *testing.B) {
		c := NewObjCache()
		setter := func(key string, val int64) {
			c.Set(key, val)
		}
		benchmarkCacheSet(b, setter, keys)
	})
	b.Run("syncMap", func(b *testing.B) {
		c := sync.Map{}
		setter := func(key string, val int64) {
			c.Store(key, val)
		}
		benchmarkCacheSet(b, setter, keys)
	})
}

func BenchmarkObjCache_Set(b *testing.B) {
	keys := initKeys(b)
	c := NewObjCache()

	b.ReportAllocs()
	b.StartTimer()
	initObjCache(b, c, keys)
	b.StopTimer()
}

func BenchmarkSyncMap_Set(b *testing.B) {
	keys := initKeys(b)
	c := sync.Map{}

	b.ReportAllocs()
	b.StartTimer()
	initSyncMap(b, c, keys)
	b.StopTimer()
}

func BenchmarkObjCache_Get(b *testing.B) {
	keys := initKeys(b)
	c := NewObjCache()
	initObjCache(b, c, keys)
	l := len(keys)

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Get(keys[i%l])
	}
	b.StopTimer()
}

func BenchmarkSyncMap_Get(b *testing.B) {
	keys := initKeys(b)
	c := sync.Map{}
	initSyncMap(b, c, keys)
	l := len(keys)

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Load(keys[i%l])
	}
	b.StopTimer()
}

func BenchmarkObjCache_Del(b *testing.B) {
	keys := initKeys(b)
	c := NewObjCache()
	initObjCache(b, c, keys)
	l := len(keys)

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Del(keys[i%l])
	}
	b.StopTimer()
}

func BenchmarkSyncMap_Del(b *testing.B) {
	keys := initKeys(b)
	c := sync.Map{}
	initSyncMap(b, c, keys)
	l := len(keys)

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Delete(keys[i%l])
	}
	b.StopTimer()
}
