/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/9 下午6:50
 * Description: 服务配置元数据信息
 **/

package cluster

import (
	"Seckill/infrastructure/stores/etcd"
	"Seckill/infrastructure/utils"
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/sirupsen/logrus"
)

type Config struct {
	LogLevel  string `json:"logLevel"` // 日志等级
	RateLimit struct {
		Middle int `json:"middle"` // 中间层限流
		Low    int `json:"low"`    // 底层限流
	} `json:"rateLimit"` // 限流器配置
	CircuitBreaker struct { // 熔断器配置
		Cpu     int `json:"cpu"`     // cpu使用率配置
		Latency int `json:"latency"` // 请求延时，单位ms
	} `json:"circuitBreaker"` // 熔断器配置
	RedisDB struct {
		Lock  int `json:"lock"`  // 锁使用db
		Cron  int `json:"cron"`  // 定时任务使用db
		Rank  int `json:"rank"`  // 排行榜使用db
		Cache int `json:"cache"` // 缓存使用db
	} `json:"redisDb"` // redis db
}

var (
	configLock = &sync.RWMutex{}
	config     = &Config{}
	configKey  = fmt.Sprintf("/%s/%s", utils.GetServiceName(), utils.GetConfigDir())
)

func InitClusterConfig() error {
	return nil
}

func GetClusterConfig() Config {
	// 获取当前配置信息 加锁为了避免获取的同时做了同步
	configLock.RLock()
	defer configLock.RUnlock()
	return *config
}

func WatchClusterConfig() error {
	// 监听集群配置DELETE和PUT事件，并即时同步到内存中
	client := etcd.GetETCDClient()
	resp, err := client.Get(context.Background(), configKey)
	if err != nil {
		return err
	}
	update := func(kv *mvccpb.KeyValue) (bool, error) {
		if string(kv.Key) == configKey {
			var tmpConfig *Config
			err = json.Unmarshal(kv.Value, &tmpConfig)
			if err != nil {
				logrus.Error("update cluster config failed, error:", err)
				return false, err
			}
			configLock.Lock()
			*config = *tmpConfig
			logrus.Info("update cluster config ", *config)
			configLock.Unlock()
			return true, nil
		}
		return false, nil
	}
	for _, kv := range resp.Kvs {
		if ok, err := update(kv); ok {
			break
		} else {
			if err != nil {
				return nil
			}
		}
	}

	go func() {
		watchChan := client.Watch(context.Background(), configKey)
		for resp := range watchChan {
			for _, event := range resp.Events {
				if event.Type == clientv3.EventTypePut {
					if ok, err := update(event.Kv); ok {
						break
					} else if err != nil {
						break
					}
				}
			}
		}
	}()
	return nil
}
