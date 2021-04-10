/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 上午10:00
 * Description:
 **/

package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetServiceName() string {
	// 获取服务名称
	return viper.GetString("global.service")
}

func GetConfigDir() string {
	// 获取etcd中服务配置根目录
	return viper.GetString("global.config")
}

func GetRedisConfig() (dbMap map[string]int) {
	// 获取redis不同db对应的
	tmpMap := viper.GetStringMap("redis.DB")
	dbMap = make(map[string]int)
	for k, val := range tmpMap {
		switch t := val.(type) {
		case int:
			dbMap[k] = int(t)
		case int8:
			dbMap[k] = int(t)
		case int16:
			dbMap[k] = int(t)
		case int32:
			dbMap[k] = int(t)
		case int64:
			dbMap[k] = int(t)
		//case bool:
		//	continue
		case float32:
			dbMap[k] = int(t)
		case float64:
			dbMap[k] = int(t)
		case uint8:
			dbMap[k] = int(t)
		case uint16:
			dbMap[k] = int(t)
		case uint32:
			dbMap[k] = int(t)
		case uint64:
			dbMap[k] = int(t)
		//case string:
		//	continue
		default:
			fmt.Print("%v = %T\n", t, t)
		}
	}
	return dbMap
}
