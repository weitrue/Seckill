/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 下午8:40
 * Description: 环形缓冲区(环形队列)
 **/

package ringbufferpool

import (
	"sync/atomic"
	"unsafe"
)

type RingBuffer struct {
	count int32            // 当前缓冲区数据数量
	size  int32            // 缓冲区大小
	head  int32            // 控制读位置
	tail  int32            // 控制写位置
	buf   []unsafe.Pointer // 缓冲区
}

func NewRingBuffer(size int32) *RingBuffer {
	return &RingBuffer{
		count: 0,
		size:  size,
		head:  0,
		tail:  0,
		buf:   make([]unsafe.Pointer, size),
	}
}

func (r *RingBuffer) Get() interface{} {
	// 从buf中取出对象
	if atomic.LoadInt32(&r.count) <= 0 {
		// 在高并发开始的时候，队列容易空，直接判断空性能最优
		return nil
	}

	if atomic.AddInt32(&r.count, -1) >= 0 {
		// 当扣减数量后没有超，就从队列里取出对象
		idx := (atomic.AddInt32(&r.head, 1) - 1) % r.size
		if obj := atomic.LoadPointer(&r.buf[idx]); obj != unsafe.Pointer(nil) {
			o := *(*interface{})(obj)
			atomic.StorePointer(&r.buf[idx], nil)
			return o
		}
	}

	// 当减数量超了，再加回去
	atomic.AddInt32(&r.count, 1)
	return nil
}

func (r *RingBuffer) Put(obj interface{}) bool {
	// 将对象放回到buf中。如果buf满了，返回false
	if atomic.LoadInt32(&r.count) >= r.size {
		// 在高并发结束的时候，队列容易满，直接判满性能最优
		return false
	}

	if atomic.AddInt32(&r.count, 1) <= r.size {
		// 当增加数量后没有超，就将对象放到队列里
		idx := (atomic.AddInt32(&r.tail, 1) - 1) % r.size
		atomic.StorePointer(&r.buf[idx], unsafe.Pointer(&obj))
		return true
	}
	return false
}
