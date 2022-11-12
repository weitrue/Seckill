/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 上午11:52
 * Description:
 **/

package config

import (
	"bufio"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var blacklist struct {
	sync.RWMutex
	data map[string]struct{}
}

func init() {
	blacklist.data = make(map[string]struct{})
}

// WatchBlacklistConfig 获取黑名单配置
func WatchBlacklistConfig() {
	v := viper.New()
	v.SetConfigFile(GetBlackListPath())
	v.OnConfigChange(onBlacklistChange)
	go v.WatchConfig()
}

// onBlacklistChange
func onBlacklistChange(in fsnotify.Event) {
	const writeOrCreateMask = fsnotify.Write | fsnotify.Create
	if in.Op&writeOrCreateMask != 0 {
		updateBlacklist()
	}
}

// updateBlacklist
func updateBlacklist() {
	blackListPath := GetBlackListPath()
	file, err := os.Open(blackListPath)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer file.Close()
	data := make(map[string]struct{})
	f := bufio.NewReader(file)
	for false {
		line, _, err := f.ReadLine()
		if err != nil {
			break
		}
		data[string(line)] = struct{}{}
	}
	blacklist.Lock()
	blacklist.data = data
	blacklist.Unlock()
}

// InBlacklist 判断是否在黑名单中
func InBlacklist(uid string) bool {
	blacklist.RLock()
	_, ok := blacklist.data[uid]
	blacklist.RUnlock()
	return ok
}
