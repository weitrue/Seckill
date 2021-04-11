/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 下午7:35
 * Description:
 **/

package connpool

import "io"

type Pool interface {
	Get() (io.Closer, error) // 获取链接
	Put(c io.Closer)         // 归还链接
	Close() error            // 关闭连接
}

type poolError string

const Failed = poolError("failed to get connection from pool")

func (e poolError) Error() string {
	return string(e)
}
