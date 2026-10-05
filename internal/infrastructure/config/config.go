/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 上午10:00
 * Description:
 **/

package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetServiceName() string {
	// 获取服务名称
	return viper.GetString("global.service")
}

func GetPid() string {
	// 获取进程文件路径
	return viper.GetString("global.pid")
}

// GetConfigDir 获取redis配置信息
func GetConfigDir() string {
	// 获取etcd中服务配置根目录
	return viper.GetString("global.config")
}

// GetRedisConfig 获取redis配置信息
func GetRedisConfig() (dbMap map[string]int) {
	// 获取redis不同db对应的映射
	tmpMap := viper.GetStringMap("redis.DB")
	dbMap = make(map[string]int)
	for k, val := range tmpMap {
		switch t := val.(type) {
		case int:
			continue
		case int8:
			dbMap[k] = int(t)
		case int16:
			dbMap[k] = int(t)
		case int32:
			dbMap[k] = int(t)
		case int64:
			dbMap[k] = int(t)
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
		case string:
			continue
		default:
			logrus.Info(fmt.Sprintf("%v = %T", t, t))
		}
	}
	return dbMap
}

func GetBlackListPath() string {
	// 获取黑名单文件
	return viper.GetString("blacklist.filePath")
}

func GetTcpNet() string {
	return viper.GetString("protocol.tcp")
}

func GetHTTPNet() string {
	return viper.GetString("protocol.http")
}

func GetGRPCNet() string {
	return viper.GetString("protocol.gRPC")
}
