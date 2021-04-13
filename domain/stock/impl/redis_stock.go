/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 上午9:47
 * Description: 基于redis实现库存数据的缓存
				库存在redis中的数据为string,key为ServiceName:ActivityID:GoodsID,value为库存值
				使用hash的话，key为ServiceName:ActivityID, field为GoodsID, value为库存值
				在redis集群中,如果用hash的话，一场活动的key是一样的，所以一场活动的库存只会在一个节点上存储，这样会会容易产生热key问题
				可以利用Set数据结构关联活动与商品信息这样 便可以批量获取一场活动的所有商品库存。
 **/

package impl

import (
	"Seckill/domain/stock"
	"Seckill/infrastructure/config/cluster"
	"Seckill/infrastructure/stores/redis"
	"Seckill/infrastructure/utils"
	"errors"
	"fmt"
	"time"
)

var (
	db = 11
)

type redisStack struct {
	activityID string // 活动ID
	goodsID    string // 商品ID
	key        string //库存key
}

func NewRedisStock(activityID, goodsID string) (stock.Stock, error) {
	if activityID == "" || goodsID == "" {
		return nil, errors.New("invalid event id or goods id")
	}
	oneStack := &redisStack{
		activityID: activityID,
		goodsID:    goodsID,
		key:        fmt.Sprintf("%s:%s:%s", utils.GetServiceName(), activityID, goodsID),
	}
	return oneStack, nil
}

func (rs *redisStack) Set(val int64, expire int64) error {
	// 设置库存
	if conf := cluster.GetClusterConfig(); &conf.RedisDB != nil {
		db = conf.RedisDB.Cache
	}
	client := redis.GetRedisClient(db)
	return client.Set(rs.key, val, time.Duration(expire)*time.Second).Err()
}

func (rs *redisStack) Get() (int64, error) {
	// 获取库存
	if conf := cluster.GetClusterConfig(); &conf.RedisDB != nil {
		db = conf.RedisDB.Cache
	}
	client := redis.GetRedisClient(db)
	if val, err := client.Get(rs.key).Int64(); err != nil && err != redis.Nil {
		return 0, err
	} else {
		return val, nil
	}
}

func (rs *redisStack) Sub(uid string) (int64, error) {
	// 基于lua脚本实现减库存的原子性操作
	if conf := cluster.GetClusterConfig(); &conf.RedisDB != nil {
		db = conf.RedisDB.Cache
	}
	client := redis.GetRedisClient(db)
	script := `
	local history = redis.call('get', KEYS[1])
	local stock = redis.call('get', KEYS[2])
	if (history and history >= 1) or stock == false or stock <= '0' then
		return -1
	else
		stock = redis.call('decr', KEYS[2])
		if stock >= 0 and redis.call('set', KEYS[1], '1', 'ex', 86400) then
			return stock
		else
			return -1
		end
	end`
	// 执行脚本

	if r, err := client.Eval(script, []string{fmt.Sprintf("%s:%s", rs.key, uid), rs.key}).Result(); err != nil {
		return -1, err
	} else if res, ok := r.(int64); ok && res != -1 {
		return res, nil
	} else {
		return -1, errors.New("redis error")
	}
}

func (rs *redisStack) Del() error {
	// 删除库存
	if conf := cluster.GetClusterConfig(); &conf.RedisDB != nil {
		db = conf.RedisDB.Cache
	}
	client := redis.GetRedisClient(db)
	return client.Del(rs.key).Err()
}

func (rs *redisStack) GetActivityID() string {
	// 获取活动id
	return rs.activityID
}

func (rs *redisStack) GetGoodsID() string {
	// 获取商品id
	return rs.goodsID
}
