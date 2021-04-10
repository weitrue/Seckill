/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/9 下午11:59
 * Description:
 **/

package redis

import (
	"Seckill/infrastructure/utils"
	"errors"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var (
	client    *redis.Client
	clientMap map[string]*redis.Client
	Nil       = redis.Nil
)

func Init() error {
	addr := viper.GetString("redis.address")
	//auth := viper.GetString("redis.auth")
	if addr == "" {
		addr = "127.0.0.0:6379"
	}
	// TODO for循环加载不同的db
	dbMap := utils.GetRedisConfig()
	clientMap = make(map[string]*redis.Client)

	for key, val := range dbMap {
		opt := &redis.Options{
			Network: "tcp",
			Addr:    addr,
			DB:      val,
		}
		client = redis.NewClient(opt)

		if client == nil {
			return errors.New("init redis client failed")
		}
		clientMap[key] = client
	}
	fmt.Println(clientMap)
	return nil
}

func GetRedisClient() *redis.Client {
	return client
}

func GetRedisClientByDB(db string) *redis.Client {
	if client, ok := clientMap[db]; ok {
		return client
	}
	return nil
}
