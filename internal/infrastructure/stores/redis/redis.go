/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/9 下午11:59
 * Description: redis 客户端初始化
 *   - 按 toml 配置的 redis.DB.* 分 db 构造多实例 client
 *   - Init 时 Ping 一次,确保 Redis 可达,否则直接报错而不是等到第一次业务操作才炸
 **/

package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/weitrue/Seckill/internal/infrastructure/config"
)

// 重导出 go-redis 的类型 / 常量,便于上层包不直接依赖 go-redis
type (
	Client  = redis.Client
	Options = redis.Options
)

// Nil 查无数据的哨兵错误(重导出,供上层 errors.Is 判断)
var Nil = redis.Nil

// pingTimeout Init 时 Ping 的超时
const pingTimeout = 3 * time.Second

var (
	clientMap map[int]*redis.Client
)

/*Init
 *@Description: 初始化 Redis 客户端;遍历 redis.DB 配置,每个 db 建一个 client 并 Ping
 *@return error
 */
func Init() error {
	logType := "RedisInit"
	addr := viper.GetString("redis.address")
	if addr == "" {
		addr = "127.0.0.1:6379"
	}
	password := viper.GetString("redis.auth")

	dbMap := config.GetRedisConfig()
	if len(dbMap) == 0 {
		return errors.New("redis.DB config is empty")
	}

	clientMap = make(map[int]*redis.Client, len(dbMap))
	ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()

	for name, dbIdx := range dbMap {
		c := redis.NewClient(&redis.Options{
			Network:  "tcp",
			Addr:     addr,
			Password: password,
			DB:       dbIdx,
		})
		if c == nil {
			return fmt.Errorf("redis NewClient returned nil, name:%s db:%d", name, dbIdx)
		}
		if pong, err := c.Ping(ctx).Result(); err != nil {
			_ = c.Close()
			logrus.Errorf("logType:%s, err:%s, step:ping, addr:%s, name:%s, db:%d",
				logType, err.Error(), addr, name, dbIdx)
			return fmt.Errorf("redis ping failed, name:%s db:%d: %w", name, dbIdx, err)
		} else {
			logrus.Infof("logType:%s, msg:connected, name:%s, db:%d, pong:%s",
				logType, name, dbIdx, pong)
		}
		clientMap[dbIdx] = c
	}
	return nil
}

/*GetRedisClient
 *@Description: 按 db 索引取出 client,找不到返回 nil(调用方需判空或优先 Init)
 *@param db
 *@return *Client
 */
func GetRedisClient(db int) *redis.Client {
	if c, ok := clientMap[db]; ok {
		return c
	}
	return nil
}

/*Close
 *@Description: 关闭所有 db 的 client,幂等
 *@return error 首个错误(其余错误日志打印)
 */
func Close() error {
	logType := "RedisClose"
	var firstErr error
	for db, c := range clientMap {
		if err := c.Close(); err != nil {
			logrus.Errorf("logType:%s, err:%s, db:%d", logType, err.Error(), db)
			if firstErr == nil {
				firstErr = err
			}
		}
	}
	clientMap = nil
	return firstErr
}
