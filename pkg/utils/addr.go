/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/9 下午5:23
 * Description: url格式化工具包
 **/

package utils

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

// 占位地址:等价于"任意网卡",需要被 extractIP 替换为具体 IP
var placeholderIPs = map[string]struct{}{
	"":        {},
	"0.0.0.0": {},
	"::":      {},
	"[::]":    {},
}

/*Extract
 *@Description: 提取 ip 和端口,把 0.0.0.0 这类占位地址替换成本机真实非 loopback IPv4
 *@param bind 绑定地址,格式: "host:port" 或 ":port"
 *@return string 格式化后的 "ip:port"
 *@return error
 */
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
	ip, err = extractIP(ip)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s", ip, port), err
}

/*extractIP
 *@Description: 占位地址返回本机一张非 loopback 的 IPv4;具体 IP 原样返回
 *@param ip 入参 ip
 *@return string
 *@return error
 */
func extractIP(ip string) (string, error) {
	if _, isPlaceholder := placeholderIPs[ip]; !isPlaceholder {
		return ip, nil
	}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, a := range addrs {
		ipNet, ok := a.(*net.IPNet)
		if !ok || ipNet.IP.IsLoopback() {
			continue
		}
		v4 := ipNet.IP.To4()
		if v4 == nil {
			continue
		}
		return v4.String(), nil
	}
	return "", errors.New("no suitable non-loopback ipv4 address found")
}
