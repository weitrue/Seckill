/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 下午5:55
 * Description:
 **/

package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weitrue/Seckill/infrastructure/stores/cache/cachei"
)

func TestIntCache(t *testing.T) {
	c := NewIntCache()
	a := assert.New(t)
	key := "test"
	c.Set(key, 1)
	v, ok := c.Get(key)
	a.Equal(int64(1), v)
	a.Equal(true, ok)

	v = c.Add(key, 5)
	a.Equal(int64(6), v)

	c.Del(key)
	_, ok = c.Get(key)
	a.Equal(false, ok)
}

func TestIntCache_Add(t *testing.T) {
	cache := NewIntCache()
	cases := []struct {
		key    string
		delta  int64
		expect int64
	}{
		{"test1", 0, 0},
		{"test1", 1, 1},
		{"test1", -1, 0},
		{"test1", 0, 0},
		{"test2", 1, 1},
		{"test3", -1, -1},
	}
	for _, c := range cases {
		if cache.Add(c.key, c.delta) != c.expect {
			t.Fatal(c)
		}
	}
}

func initIntCache(b *testing.B, c cachei.IntCache, keys []string) {
	l := len(keys)
	for i := 0; i < b.N; i++ {
		c.Set(keys[i%l], int64(i))
	}
}

func BenchmarkIntCache_Add(b *testing.B) {
	keys := initKeys(b)
	c := NewIntCache()
	initIntCache(b, c, keys)
	l := len(keys)

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Add(keys[i%l], 1)
	}
	b.StopTimer()
}

// Set
func BenchmarkIntCache_Set(b *testing.B) {
	keys := initKeys(b)
	c := NewIntCache()

	b.ReportAllocs()
	b.StartTimer()
	initIntCache(b, c, keys)
	b.StopTimer()
}

// Get
func BenchmarkIntCache_Get(b *testing.B) {
	keys := initKeys(b)
	c := NewIntCache()
	initIntCache(b, c, keys)
	l := len(keys)

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Get(keys[i%l])
	}
	b.StopTimer()
}

// 删除
func BenchmarkIntCache_Del(b *testing.B) {
	keys := initKeys(b)
	c := NewIntCache()
	initIntCache(b, c, keys)
	l := len(keys)

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Del(keys[i%l])
	}
	b.StopTimer()
}