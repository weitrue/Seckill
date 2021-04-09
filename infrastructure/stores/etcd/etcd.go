/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/9 下午12:00
 * Description: etcd 存储秒杀服务集群信息
 **/

package etcd

import (
	"sync"
	"time"

	etcd "github.com/coreos/etcd/clientv3"
	"github.com/spf13/viper"
)

var (
	etcdClient *etcd.Client
	etcdOnce   = &sync.Once{}
)

func Init() error {
	// 初始化etcd客户端
	var err error
	etcdOnce.Do(func() {
		endpoints := viper.GetStringSlice("etcd.endpoints")
		username := viper.GetString("etcd.username")
		password := viper.GetString("etcd.password")
		config := etcd.Config{
			Endpoints:   endpoints,
			DialTimeout: 5 * time.Second,
			Username:    username,
			Password:    password,
		}
		etcdClient, err = etcd.New(config)
	})
	return err
}

func GetETCDClient() *etcd.Client {
	return etcdClient
}
