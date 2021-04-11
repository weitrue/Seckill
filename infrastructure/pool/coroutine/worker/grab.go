/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 下午11:08
 * Description: 抢占式工作机制：多个协程共享一个任务池，所有任务放到共享任务池中，每个协程区共享任务池中抢任务并执行
 **/

package worker

import (
	"Seckill/infrastructure/pool/coroutine"
	"math"
	"sync"
	"sync/atomic"

	"github.com/sirupsen/logrus"
)

const (
	minBufferSize = 10
	minNumber
)

type worker struct {
	number   int                 // 协程数
	size     int                 // 任务池大小
	closed   int32               // 关闭状态：1表示关闭，0表示未关闭
	taskPool chan coroutine.Task // 任务池
	wg       sync.WaitGroup      // 状态同步
}

func NewWorker(number, size int) coroutine.Worker {
	number = int(math.Max(float64(number), float64(minNumber)))
	size = int(math.Max(float64(size), float64(minBufferSize)))
	w := &worker{
		number:   number,
		size:     size,
		taskPool: make(chan coroutine.Task, size),
	}
	w.wg.Add(number)
	for i := 0; i < number; i++ {
		go w.run()
	}
	return w
}

func (w *worker) Push(t coroutine.Task) bool {
	if w.isClosed() {
		return false
	}
	// 阻塞式
	w.taskPool <- t
	//// 非阻塞式
	//select {
	//case w.taskPool <- t:
	//	break
	//default:
	//	logrus.Error(fmt.Sprintf("Push Task %v failed", t))
	//}
	return true
}

func (w *worker) Close() error {
	if !w.isClosed() && atomic.CompareAndSwapInt32(&w.closed, 0, 1) {
		close(w.taskPool)
		w.wg.Wait()
	}
	return nil
}

func (w *worker) run() {
	defer w.wg.Done()
	for task := range w.taskPool {
		w.process(task)
	}
}

func (w *worker) process(t coroutine.Task) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Error(err)
		}
	}()
	t.Do()
}

func (w *worker) isClosed() bool {
	return atomic.LoadInt32(&w.closed) == 1
}
