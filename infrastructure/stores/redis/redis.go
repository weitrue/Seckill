/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/9 下午11:59
 * Description:
 **/

package redis

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"github.com/weitrue/Seckill/infrastructure/utils"
)

var (
	client    *redis.Client
	clientMap map[int]*redis.Client
	Nil       = redis.Nil
)

func Init() error {
	addr := viper.GetString("redis.address")
	//auth := viper.GetString("redis.auth")
	if addr == "" {
		addr = "127.0.0.0:6379"
	}
	// TODO 初始化时 获取 etcd 上配置
	var dbMap map[string]int
	//if conf := cluster.GetClusterConfig(); &conf.RedisDB != nil {
	//	data, _ := json.Marshal(&conf.RedisDB)
	//	dbMap = make(map[string]int)
	//	_ = json.Unmarshal(data, &dbMap)
	//} else {
	//
	//}
	dbMap = utils.GetRedisConfig()
	clientMap = make(map[int]*redis.Client)

	for _, val := range dbMap {
		opt := &redis.Options{
			Network: "tcp",
			Addr:    addr,
			DB:      val,
		}
		client = redis.NewClient(opt)

		if client == nil {
			return errors.New("init redis client failed")
		}
		clientMap[val] = client
	}
	fmt.Println(clientMap)
	return nil
}

func GetRedisClient(db int) *redis.Client {
	if client, ok := clientMap[db]; ok {
		return client
	}
	return nil
}
