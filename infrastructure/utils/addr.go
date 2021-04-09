/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/9 下午5:23
 * Description: url格式化工具包
 **/

package utils

import (
	"fmt"
	"strings"

	"github.com/micro/go-micro/util/addr"
)

func Extract(bind string) (string, error) {
	var (
		ip   string
		port string
		err  error
	)
	parts := strings.Split(bind, ":")

	if len(parts) == 2 {
		ip = parts[0]
		port = parts[1]
	} else {
		ip = "0.0.0.0"
		port = parts[0]
	}
	ip, err = addr.Extract(ip)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s", ip, port), err
}
