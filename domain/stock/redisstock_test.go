/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 下午2:30
 * Description:
 **/

package stock

import (
	"reflect"
	"testing"

	"github.com/weitrue/Seckill/domain/stock/stocki"
	"github.com/weitrue/Seckill/infrastructure/stores/redis"

	"github.com/stretchr/testify/assert"
)

func TestStock(t *testing.T) {
	var (
		st  stocki.Stock
		err error
		val int64
	)
	a := assert.New(t)
	if err = redis.Init(); err != nil {
		t.Fatal(err)
	}

	st, err = NewRedisStock("101", "1001")
	a.Nil(err)

	defer func() {
		c := redis.GetRedisClient(11)
		c.Del("seckill:101:1001")
		c.Del("seckill:101:1001:111")
	}()
	err = st.Set(10, 100)
	a.Nil(err)

	val, err = st.Get()
	a.Nil(err)
	a.Equal(10, val)

	val, err = st.Sub("111")
	a.Nil(err)
	a.Equal(9, val)

	err = st.Del()
	a.Nil(err)

	val, err = st.Get()
	a.Nil(err)
	a.Equal(0, val)
}

func TestNewRedisStock(t *testing.T) {
	_ = redis.Init()
	client := redis.GetRedisClient(db)
	script := `
	return redis.call('get', 'seckill:101:1001')
	`
	res, err := client.Eval(script, []string{}).Result()
	t.Log(res, reflect.TypeOf(res), err)
}
