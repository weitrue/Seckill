/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午4:35
 * Description:
 **/

package circuitbreaker

import (
	"sync/atomic"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	minDuration  = 100
	minTotal     = 1000
	minFailsRate = 2
)

type CircuitBreaker struct {
	totalCounter Counter // 请求计数器
	failsCounter Counter //失败请求计数器

	duration       int64 // 时间窗口
	latencyLimit   int64 // 最大延迟限制
	totalLimit     int64 // 最大请求数限制
	failsRateLimit int64 // 最大失败率限制

	recoverFailsRate int64 // 恢复请求的最低失败率
	startTime        int64 // 时间窗口开始时间
	allow            int64 // 当前是否允许请求执行, 0表示不允许，1表示允许
}

type CBOption func(cb *CircuitBreaker)

func WithDuration(duration int64) CBOption {
	return func(cb *CircuitBreaker) {
		cb.duration = duration
	}
}

func WithLatencyLimit(latencyLimit int64) CBOption {
	return func(cb *CircuitBreaker) {
		cb.latencyLimit = latencyLimit
	}
}

func WithFailsLimit(failsRateLimit int64) CBOption {
	return func(cb *CircuitBreaker) {
		cb.failsRateLimit = failsRateLimit
	}
}

func WithTotalLimit(totalLimit int64) CBOption {
	return func(cb *CircuitBreaker) {
		cb.totalLimit = totalLimit
	}
}

func NewCircuitBreaker(opts ...CBOption) *CircuitBreaker {
	cb := &CircuitBreaker{
		totalCounter:     0,
		failsCounter:     0,
		duration:         0,
		latencyLimit:     0,
		totalLimit:       0,
		failsRateLimit:   0,
		recoverFailsRate: 0,
		startTime:        0,
		allow:            1,
	}
	// 基于CBOption实现可扩展的变参
	for _, opt := range opts {
		opt(cb)
	}
	if cb.duration < minDuration {
		cb.duration = minDuration
	}
	if cb.totalLimit < minTotal {
		cb.totalLimit = minTotal
	}
	if cb.failsRateLimit < minFailsRate {
		cb.failsRateLimit = minFailsRate
	}
	cb.recoverFailsRate = cb.recoverFailsRate / 2
	return cb
}

func (cb *CircuitBreaker) Allow(f func() bool) bool {
	fails := cb.failsCounter.Load()
	total := cb.totalCounter.Load()
	start := time.Now().UnixNano() / int64(time.Millisecond)
	if start > cb.startTime+cb.duration {
		// 更新请求计数器和时间窗口
		atomic.StoreInt64(&cb.startTime, start)
		cb.failsCounter.Reset()
		cb.totalCounter.Reset()
		atomic.StoreInt64(&cb.allow, 1)
	}
	cb.totalCounter.Add()
	// 根据请求失败率判断当前是否需要熔断 0表示需要， 1表示不需要 (fails/total >= failsRateLimit% || total >= totalLimit)
	allow := !(total > 0 && fails*100/cb.failsRateLimit >= total || total >= cb.totalLimit)
	if atomic.LoadInt64(&cb.allow) == 0 {
		if fails*100/cb.recoverFailsRate > total { // fails/total > recoverFailsRate%
			allow = false
		} else if !allow {
			atomic.StoreInt64(&cb.allow, 1)
		}
	} else if !allow {
		atomic.StoreInt64(&cb.allow, 0)
	}
	if !allow {
		logrus.Error("not allowed")
		return false
	}
	// 处理请求的函数
	ok := f()
	// 请求处理结束 统计窗口内请求失败率、延迟不达标率、请求数
	end := time.Now().UnixNano() / int64(time.Millisecond)
	if (cb.latencyLimit > 0 && end-start >= cb.latencyLimit) || !ok {
		cb.failsCounter.Add()
	}
	return true
}
