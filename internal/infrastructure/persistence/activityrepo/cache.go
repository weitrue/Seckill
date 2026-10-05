/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2024/10/06
 * Description: Activity 仓储的 Redis 缓存层
 *   - Hash 存元数据: seckill:activity:<id>  → {id, name, start_time, end_time, status}
 *   - Hash 存商品:   seckill:activity:<id>:goods → goods_id → json(Goods)
 *   - 过期时间 defaultTTL(7d),后台 Create/Update 时主动刷新;读取未命中时回源 MySQL 并写回
 **/

package activityrepo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/weitrue/Seckill/internal/domain/activity"
)

const (
	keyPrefix       = "seckill:activity:"
	goodsKeySuffix  = ":goods"
	defaultCacheTTL = 7 * 24 * time.Hour
)

// cache activity Redis Hash 缓存实现(私有,通过组合层对外)
type cache struct {
	client *redis.Client
	ttl    time.Duration
}

func newCache(client *redis.Client, ttl time.Duration) *cache {
	if ttl <= 0 {
		ttl = defaultCacheTTL
	}
	return &cache{client: client, ttl: ttl}
}

// activityKey 元数据 Hash key: seckill:activity:<id>
func activityKey(id int64) string {
	return keyPrefix + strconv.FormatInt(id, 10)
}

// goodsKey 商品清单 Hash key: seckill:activity:<id>:goods
func goodsKey(id int64) string {
	return keyPrefix + strconv.FormatInt(id, 10) + goodsKeySuffix
}

/*Set
 *@Description: 写入活动元数据 + 商品清单(整体替换)
 *@receiver c
 *@param ctx
 *@param a
 *@return error
 */
func (c *cache) Set(ctx context.Context, a *activity.Activity) error {
	if a == nil || a.ID <= 0 {
		return errors.New("activityrepo.cache: invalid activity")
	}
	ak := activityKey(a.ID)
	gk := goodsKey(a.ID)
	meta := map[string]interface{}{
		"id":          a.ID,
		"name":        a.Name,
		"start_time":  a.StartTime,
		"end_time":    a.EndTime,
		"status":      int8(a.Status),
		"create_time": a.CreateTime,
		"update_time": a.UpdateTime,
	}
	pipe := c.client.TxPipeline()
	// 元数据先删后写(应对字段删除场景)
	pipe.Del(ctx, ak)
	pipe.HSet(ctx, ak, meta)
	pipe.Expire(ctx, ak, c.ttl)
	// 商品清单
	pipe.Del(ctx, gk)
	if len(a.Goods) > 0 {
		fields := make(map[string]interface{}, len(a.Goods))
		for _, g := range a.Goods {
			b, err := json.Marshal(g)
			if err != nil {
				return err
			}
			fields[g.GoodsID] = b
		}
		pipe.HSet(ctx, gk, fields)
		pipe.Expire(ctx, gk, c.ttl)
	}
	_, err := pipe.Exec(ctx)
	return err
}

/*Get
 *@Description: 读取活动元数据 + 商品清单;未命中返回 (nil, nil)
 *@receiver c
 *@param ctx
 *@param id
 *@return *activity.Activity
 *@return error
 */
func (c *cache) Get(ctx context.Context, id int64) (*activity.Activity, error) {
	if id <= 0 {
		return nil, nil
	}
	ak := activityKey(id)
	gk := goodsKey(id)
	meta, err := c.client.HGetAll(ctx, ak).Result()
	if err != nil {
		return nil, err
	}
	if len(meta) == 0 {
		return nil, nil
	}
	a, err := parseMeta(meta)
	if err != nil {
		return nil, err
	}
	// 读商品清单(不存在也容忍,活动可能没配商品)
	goodsMap, err := c.client.HGetAll(ctx, gk).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}
	for _, raw := range goodsMap {
		var g activity.Goods
		if err = json.Unmarshal([]byte(raw), &g); err != nil {
			return nil, err
		}
		a.Goods = append(a.Goods, &g)
	}
	return a, nil
}

/*Del
 *@Description: 清除活动缓存
 *@receiver c
 *@param ctx
 *@param id
 *@return error
 */
func (c *cache) Del(ctx context.Context, id int64) error {
	if id <= 0 {
		return nil
	}
	return c.client.Del(ctx, activityKey(id), goodsKey(id)).Err()
}

// parseMeta 从 HGetAll 的 string map 解出 Activity 元数据
func parseMeta(meta map[string]string) (*activity.Activity, error) {
	a := &activity.Activity{}
	for k, v := range meta {
		switch k {
		case "id":
			if n, err := strconv.ParseInt(v, 10, 64); err == nil {
				a.ID = n
			}
		case "name":
			a.Name = v
		case "start_time":
			a.StartTime, _ = strconv.ParseInt(v, 10, 64)
		case "end_time":
			a.EndTime, _ = strconv.ParseInt(v, 10, 64)
		case "status":
			if n, err := strconv.ParseInt(v, 10, 8); err == nil {
				a.Status = activity.Status(n)
			}
		case "create_time":
			a.CreateTime, _ = strconv.ParseInt(v, 10, 64)
		case "update_time":
			a.UpdateTime, _ = strconv.ParseInt(v, 10, 64)
		}
	}
	if a.ID <= 0 {
		return nil, fmt.Errorf("activityrepo.cache: invalid meta, id missing")
	}
	return a, nil
}
