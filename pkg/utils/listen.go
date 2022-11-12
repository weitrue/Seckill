/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午3:53
 * Description:
 **/

package utils

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"syscall"

	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"

	"github.com/weitrue/Seckill/infrastructure/config"
)

// Listen 服务监听
func Listen(network, addr string) (net.Listener, error) {
	// 设置服务监听为可重复监听
	listenConf := &net.ListenConfig{
		Control: func(network, address string, c syscall.RawConn) error {
			var err error
			lisErr := c.Control(func(fd uintptr) {
				// 参数 net.ipv4.tcp_tw_reuse=1
				err = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, unix.SO_REUSEPORT, 1)
				if err != nil {
					logrus.Error("set socket option failed", err)
				}
			})
			if lisErr != nil {
				logrus.Error("control listener failed", err)
				err = lisErr
			}
			return err
		},
	}
	return listenConf.Listen(context.Background(), network, addr)
}

// UpdateProcess 更新程序
func UpdateProcess(service string) {
	// 更新程序 基于信号机制给老版本发信号

	fileName := config.GetPid()
	fileName = fmt.Sprintf("%s.%s", fileName, service)

	// 获取老版本进程 pid
	if pidFile, err := os.Open(fileName); err == nil {
		pidBytes, err := ioutil.ReadAll(pidFile)
		if err != nil {
			return
		}
		pid, err := strconv.Atoi(string(pidBytes))
		if err != nil {
			return
		}

		if pid > 0 {
			// 为了避免因某些原因老版本程序无法退出，尝试发送多个信号，最后一次 SIGKILL 将强制结束老版程序
			signals := []syscall.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL}
			if process, err := os.FindProcess(pid); err == nil {
				for _, sig := range signals {
					if err = process.Signal(sig); err == nil {
						break
					}
					var stat *os.ProcessState
					//  等待老版程序退出
					stat, err = process.Wait()
					if err != nil || stat.Exited() {
						break
					}
				}
			}
		}
		_ = pidFile.Close()
	}
	// 将文件中的pid更新为新版本的pid
	if pidFile, err := os.Create(fileName); err == nil {
		pid := os.Getpid()
		_, _ = pidFile.Write([]byte(strconv.Itoa(pid)))
		_ = pidFile.Close()
	}
}
