package cache

import (
	"os"
	"strconv"
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
	keys := make([]string, 0)
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

func initCache(b *testing.B, c Cache, keys []string) {
	l := len(keys)
	for i := 0; i < b.N; i++ {
		c.Set(keys[i%l], int64(i))
	}
}

func BenchmarkCache_Set(b *testing.B) {
	keys := initKeys(b)
	c := NewObjCache()

	b.ReportAllocs()
	b.StartTimer()
	initCache(b, c, keys)
	b.StopTimer()
}

func BenchmarkObjCache_Get(b *testing.B) {
	keys := initKeys(b)
	c := NewObjCache()
	initCache(b, c, keys)
	l := len(keys)

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Get(keys[i%l])
	}
	b.StopTimer()
}

func BenchmarkCache_Del(b *testing.B) {
	keys := initKeys(b)
	c := NewObjCache()
	initCache(b, c, keys)
	l := len(keys)

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Del(keys[i%l])
	}
	b.StopTimer()
}
