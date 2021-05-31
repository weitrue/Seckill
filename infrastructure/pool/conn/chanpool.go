/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 下午7:43
 * Description: 基于channel实现的连接池 高并发情况下，因为内部使用锁保证并发访问的可靠性，性能较低
 **/

package conn

import (
	"io"

	"github.com/weitrue/Seckill/infrastructure/pool/conn/pooli"
)

type chanPool struct {
	name    string                    // 命名
	size    int                       // 缓冲区大小
	ch      chan io.Closer            // channel
	newFunc func() (io.Closer, error) // 创建函数

}

func NewChanPool(name string, size int, newFunc func() (io.Closer, error)) pooli.Pool {
	return &chanPool{
		name:    name,
		size:    size,
		ch:      make(chan io.Closer, size),
		newFunc: newFunc,
	}
}

func (p *chanPool) Get() (io.Closer, error) {
	// 非阻塞模式
	select {
	case c := <-p.ch:
		return c, nil
	default:
		if p.newFunc != nil {
			return p.newFunc()
		}
	}
	return nil, pooli.Failed
}

func (p *chanPool) Put(c io.Closer) {
	if c == nil {
		return
	}
	select {
	case p.ch <- c:
		break
	default:
		_ = c.Close()
	}
}

func (p *chanPool) Close() error {
	close(p.ch)
	p.ch = nil
	return nil
}
