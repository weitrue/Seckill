/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 下午2:30
 * Description: redisstock 集成测试(需要本地 Redis 实例)
 **/

package redisstock

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/weitrue/Seckill/internal/domain/stock"
	"github.com/weitrue/Seckill/internal/infrastructure/stores/redis"
)

func TestStock(t *testing.T) {
	var (
		st  stock.Stock
		err error
		val int64
	)
	a := assert.New(t)
	if err = redis.Init(); err != nil {
		t.Fatal(err)
	}

	st, err = NewRedisStock("101", "1001")
	a.Nil(err)

	ctx := context.Background()
	defer func() {
		c := redis.GetRedisClient(11)
		c.Del(ctx, "seckill:101:1001")
		c.Del(ctx, "seckill:101:1001:111")
	}()
	err = st.Set(ctx, 10, 100)
	a.Nil(err)

	val, err = st.Get(ctx)
	a.Nil(err)
	a.Equal(int64(10), val)

	val, err = st.Sub(ctx, "111")
	a.Nil(err)
	a.Equal(int64(9), val)

	err = st.Del(ctx)
	a.Nil(err)

	val, err = st.Get(ctx)
	a.Nil(err)
	a.Equal(int64(0), val)
}

func TestNewRedisStock(t *testing.T) {
	_ = redis.Init()
	client := redis.GetRedisClient(db)
	ctx := context.Background()
	script := `
	return redis.call('get', 'seckill:101:1001')
	`
	res, err := client.Eval(ctx, script, []string{}).Result()
	t.Log(res, reflect.TypeOf(res), err)
}
