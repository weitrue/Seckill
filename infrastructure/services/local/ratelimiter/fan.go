/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午6:09
 * Description:
 **/

package ratelimiter

import (
	"Seckill/infrastructure/pool/coroutine"
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

const (
	minRate = 1
	minSize = 10
	FanIn   = 1 << 0
	FanOut  = 1 << 1
)

type fanInOut struct {
	sync.RWMutex
	queueIn   chan coroutine.Task // 用于 Fan-In 模式的队列 queueIn
	queueOut  chan coroutine.Task // 用于 Fan-Out 模式的队列 queueOut
	startTime int64               // 开始时间
	rate      int64               // 流量速度 控制每次操作数据的时间间隔
	duration  time.Duration
	closed    int64 // 状态：0表示开启，1表示关闭
	mode      int
}

func NewRateLimiter(size, rate int64, mode int) (RateLimiter, error) {
	// 接收 size、rate、mode 这三个参数，分别用于设置队列缓冲区、流量速度、模式 mode 用来控制三种模式的主要流程
	modeMask := FanIn | FanOut
	if mode > modeMask || modeMask&mode == 0 {
		return nil, errors.New("wrong flag")
	}
	if rate < minRate {
		rate = minRate
	}
	if size < minSize {
		size = minSize
	}
	f := &fanInOut{
		startTime: 0,
		rate:      rate,
		duration:  time.Second / time.Duration(rate),
		closed:    0,
		mode:      mode,
	}

	if FanIn&mode != 0 {
		// Fan-In队列
		f.queueIn = make(chan coroutine.Task, size)
	}
	if FanOut&mode != 0 {
		// Fan-Out队列
		f.queueOut = make(chan coroutine.Task, size)
	}
	if mode == modeMask {
		go f.exchange()
	}
	return f, nil
}

func (f *fanInOut) Push(t coroutine.Task) bool {
	// 给生产者
	if atomic.LoadInt64(&f.closed) == 1 {
		// 已关闭 不能推给生产者
		return false
	}
	f.RLock()
	defer f.RUnlock()

	if atomic.LoadInt64(&f.closed) == 1 {
		return false
	}
	if FanIn&f.mode != 0 {
		// FanIn模式 非阻塞推送任务到队列
		select {
		case f.queueIn <- t:
			return true
		default:
			return false
		}
	} else {
		// FanOut模式 需要一段时间，然后阻塞式推送到队列
		f.sleep()
		f.queueOut <- t
		return true
	}
}

func (f *fanInOut) Pop() (coroutine.Task, bool) {
	// 出队给消费
	if FanOut&f.mode != 0 {
		// FanOut模式 直接消费队列
		task, ok := <-f.queueOut
		return task, ok
	} else {
		// FanIn模式 需要等待一段时间 再消费队列
		f.sleep()
		task, ok := <-f.queueIn
		return task, ok
	}
}

func (f *fanInOut) Close() error {
	// 关闭限流器
	f.Lock()
	defer f.Unlock()
	if atomic.CompareAndSwapInt64(&f.closed, 0, 1) {
		if f.mode&FanIn != 0 {
			close(f.queueIn)
		} else if f.mode == FanOut {
			close(f.queueOut)
		}
	}
	return nil
}

func (f *fanInOut) exchange() {
	// Goroutine以固定速度消费Fan-In队列，将任务放到Fan-Out队列中
	for t := range f.queueIn {
		f.sleep()
		f.queueOut <- t
	}
	close(f.queueOut)
}

func (f *fanInOut) sleep() {
	now := time.Now().UnixNano()
	delta := f.duration - time.Duration(now-atomic.LoadInt64(&f.startTime))
	if delta > time.Millisecond {
		time.Sleep(delta)
	}
	atomic.StoreInt64(&f.startTime, now)
}
