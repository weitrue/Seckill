/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午3:53
 * Description:
 **/

package utils

import (
	"context"
	"net"
	"syscall"

	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

func Listen(network string, addr string) (net.Listener, error) {
	listenConf := &net.ListenConfig{
		Control: func(network, address string, c syscall.RawConn) error {
			var err error
			lisErr := c.Control(func(fd uintptr) {
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
	return listenConf.Listen(context.Background(), "tcp", addr)
}
