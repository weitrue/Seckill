/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/12/23 3:30 下午
 * Description:
 **/

package int

import (
	"os"
	"strconv"
	"testing"

	"github.com/weitrue/Seckill/infrastructure/stores/cache"
)


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

func initIntCache(b *testing.B, c cache.IntCache, keys []string) {
	l := len(keys)
	for i := 0; i < b.N; i++ {
		c.Set(keys[i%l], int64(i))
	}
}

// Set
func BenchmarkIntCache_Set(b *testing.B) {
	keys := initKeys(b)
	c := NewIntegerCache()
	b.ReportAllocs()
	b.StartTimer()
	initIntCache(b, c, keys)
	b.StopTimer()
}

// Get
func BenchmarkIntCache_Get(b *testing.B) {
	keys := initKeys(b)
	c := NewIntegerCache()
	initIntCache(b, c, keys)
	l := len(keys)

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Get(keys[i%l])
	}
	b.StopTimer()
}
