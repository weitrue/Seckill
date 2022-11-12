/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 下午7:44
 * Description:
 **/

package ringbuffpool

import (
	"io"
	"sync/atomic"

	"github.com/weitrue/Seckill/infrastructure/pool"
	"github.com/weitrue/Seckill/infrastructure/pool/ringbuffpool/ringbuffer"
)

type ringBufferPool struct {
	closed  int32
	name    string
	rb      *ringbuffer.RingBuffer
	newFunc func() (io.Closer, error)
}

func NewRingBufferPool(name string, size int, newFunc func() (io.Closer, error)) pool.Pool {
	return &ringBufferPool{
		name:    name,
		rb:      ringbuffer.NewRingBuffer(int32(size)),
		newFunc: newFunc,
	}
}

func (p ringBufferPool) Get() (io.Closer, error) {
	// 获取连接
	var (
		err error
		c   io.Closer
	)
	if atomic.LoadInt32(&p.closed) != 0 {
		// 连接池已关闭
		return nil, pool.Failed
	}
	obj := p.rb.Get()
	if c, _ = obj.(io.Closer); c != io.Closer(nil) {
		// 连接池对象合法
		return c, err
	} else if p.newFunc != nil {
		// 可以新建连接对象
		return p.newFunc()
	}
	return nil, pool.Failed
}

func (p ringBufferPool) Put(c io.Closer) {
	// 添加连接
	if c == io.Closer(nil) {
		return
	}
	if atomic.LoadInt32(&p.closed) != 0 || !p.rb.Put(c) {
		// 连接池已关闭 || 无法加入缓存队列
		_ = c.Close()
	}
}

func (p ringBufferPool) Close() error {
	if !atomic.CompareAndSwapInt32(&p.closed, 0, 1) {
		// closed != 0, 即已关闭
		return nil
	}
	// TODO 记录关闭连接池错误日志
	// errs := make(map[interface{}]error, 0)
	for obj := p.rb.Get(); obj != nil; obj = p.rb.Get() {
		// 循环取出所有连接
		if c, ok := obj.(io.Closer); ok {
			_ = c.Close()
			//if err != nil {
			//	errs[obj] = err
			//	logrus.Error(fmt.Sprintf("close %v pool err:", obj), err)
			//}
		}
	}
	return nil
}
