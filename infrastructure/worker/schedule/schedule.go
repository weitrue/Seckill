/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 下午11:09
 * Description: 调度式工作机制：按照优先级将任务分配到对应的协程上
 **/

package schedule

import (
	"math"
	"sync/atomic"

	"github.com/weitrue/Seckill/infrastructure/worker"
	"github.com/weitrue/Seckill/infrastructure/worker/grab"
)

const (
	MinBufferSize = 10
	MinNumber
)

type priorityWorker struct {
	priorities int             // 优先级数
	number     int             // 每个优先级的协程数
	size       int             //  每个优先级的任务池大小
	closed     int32           // 关闭状态：1表示关闭，0表示未关闭
	workers    []worker.Worker // 各个优先级，没有优先级都是抢占式工作机制
}

func NewPriorityWorker(number, size, priorities int) worker.Worker {
	priorities = int(math.Max(float64(priorities), float64(MinNumber)))
	number = (number - 1 + priorities) / priorities
	if number < MinNumber {
		number = MinNumber
	}
	size = (size - 1 + priorities) / priorities
	if size < MinBufferSize {
		size = MinBufferSize
	}
	w := &priorityWorker{
		priorities: priorities,
		number:     number,
		size:       size,
		closed:     0,
		workers:    make([]worker.Worker, priorities),
	}
	for i := 0; i < priorities; i++ {
		// 每个优先级分别创建一个抢占式协程池
		w.workers[i] = grab.NewWorker(number, size)
	}
	return w
}

func (pw priorityWorker) Push(t worker.Task) bool {
	if pw.isClosed() {
		return false
	}

	if pt, ok := t.(worker.PriorityTask); ok {
		// 获取任务中的优先级
		priority := pt.Priority()
		if priority < 0 {
			// 任务优先级 < 0 取最小优先级
			priority = 0
		} else if priority >= pw.priorities {
			// 如果任务优先级 > 优先级数 取最大优先级
			priority = pw.priorities - 1
		}
		//  推送任务到对应优先级的抢占式协程池
		return pw.workers[priority].Push(t)
	} else {
		return pw.workers[pw.priorities-1].Push(t)
	}
}

func (pw priorityWorker) Close() error {
	if !pw.isClosed() && atomic.CompareAndSwapInt32(&pw.closed, 0, 1) {
		for _, w := range pw.workers {
			// 关闭所有协程池
			_ = w.Close()
		}
		return nil
	}
	return nil
}

func (pw *priorityWorker) isClosed() bool {
	return atomic.LoadInt32(&pw.closed) == 1
}
