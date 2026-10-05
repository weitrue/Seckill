/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 上午9:47
 * Description: 基于 redis 实现库存数据的缓存
 *   库存在 redis 中的数据为 string,key 为 ServiceName:ActivityID:GoodsID,value 为库存值
 *   使用 hash 的话,key 为 ServiceName:ActivityID, field 为 GoodsID, value 为库存值
 *   在 redis 集群中,如果用 hash 的话,一场活动的 key 是一样的,一场活动的库存只会落到一个节点,容易产生热 key
 *   可以利用 Set 数据结构关联活动与商品信息,这样便可以批量获取一场活动的所有商品库存
 *
 *   基于 go-redis v9: 所有命令第一参数吃 context.Context,
 *   Nil 判断用 errors.Is(err, redis.Nil)
 **/

package redisstock

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/weitrue/Seckill/internal/domain/stock"
	"github.com/weitrue/Seckill/internal/infrastructure/config"
	"github.com/weitrue/Seckill/internal/infrastructure/config/cluster"
	seckillredis "github.com/weitrue/Seckill/internal/infrastructure/stores/redis"
)

// 默认库存 db(cache 维度),运行时若 cluster 配置有 RedisDB.Cache 会覆盖
var db = 11

type redisStack struct {
	activityID string // 活动ID
	goodsID    string // 商品ID
	key        string // 库存key
}

// 编译期类型断言: 确保 *redisStack 实现了 stock.Stock
var _ stock.Stock = (*redisStack)(nil)

/*NewRedisStock
 *@Description: 构造 redis 库存实例;签名直接满足 stock.Factory
 *@param activityID
 *@param goodsID
 *@return stock.Stock
 *@return error
 */
func NewRedisStock(activityID, goodsID string) (stock.Stock, error) {
	if activityID == "" || goodsID == "" {
		return nil, errors.New("invalid event id or goods id")
	}
	oneStack := &redisStack{
		activityID: activityID,
		goodsID:    goodsID,
		key:        fmt.Sprintf("%s:%s:%s", config.GetServiceName(), activityID, goodsID),
	}
	return oneStack, nil
}

// client 取当前 cache db 的客户端;cluster 热更新库存 db 时在这里拾取
func (rs *redisStack) client() *redis.Client {
	if conf := cluster.GetClusterConfig(); &conf.RedisDB != nil && conf.RedisDB.Cache > 0 {
		db = conf.RedisDB.Cache
	}
	return seckillredis.GetRedisClient(db)
}

/*Set
 *@Description: 设置库存并设定过期时间
 *@receiver rs
 *@param ctx
 *@param val 库存值
 *@param expire 过期秒数
 *@return error
 */
func (rs *redisStack) Set(ctx context.Context, val, expire int64) error {
	return rs.client().Set(ctx, rs.key, val, time.Duration(expire)*time.Second).Err()
}

/*Get
 *@Description: 读取剩余库存;key 不存在返回 (0, nil)
 *@receiver rs
 *@param ctx
 *@return int64
 *@return error
 */
func (rs *redisStack) Get(ctx context.Context) (int64, error) {
	val, err := rs.client().Get(ctx, rs.key).Int64()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, nil
		}
		return 0, err
	}
	return val, nil
}

/*Sub
 *@Description: 基于 Lua 脚本原子扣减,同时做用户维度幂等(history 字段)
 *@receiver rs
 *@param ctx
 *@param uid 用户 ID
 *@return int64 扣减后剩余库存;超卖/重复/异常 返回 -1
 *@return error
 */
func (rs *redisStack) Sub(ctx context.Context, uid string) (int64, error) {
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

	r, err := rs.client().Eval(ctx, script, []string{fmt.Sprintf("%s:%s", rs.key, uid), rs.key}).Result()
	if err != nil {
		return -1, err
	}
	res, ok := r.(int64)
	if !ok || res == -1 {
		return -1, errors.New("redis error")
	}
	return res, nil
}

/*Del
 *@Description: 删除库存 key
 *@receiver rs
 *@param ctx
 *@return error
 */
func (rs *redisStack) Del(ctx context.Context) error {
	return rs.client().Del(ctx, rs.key).Err()
}

/*GetActivityID
 *@Description: 返回活动 ID
 *@receiver rs
 *@return string
 */
func (rs *redisStack) GetActivityID() string {
	return rs.activityID
}

/*GetGoodsID
 *@Description: 返回商品 ID
 *@receiver rs
 *@return string
 */
func (rs *redisStack) GetGoodsID() string {
	return rs.goodsID
}
