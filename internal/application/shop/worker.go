/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2024/10/06
 * Description: shop 的消费端装配
 *   queue(漏桶) → consumeLoop(单协程搬运) → workers(协程池并发执行 task)
 *   queue 和 workers 由外部构造并注入 Service,Service 只负责:
 *     1. 启动 consumeLoop
 *     2. Close 时关 queue(让 Consume 返回 err 退出 loop)和 workers
 **/

package shop

import (
	"github.com/sirupsen/logrus"
)

/*Start
 *@Description: 启动消费 goroutine;幂等,重复调用无副作用
 *@receiver s
 *@return error
 */
func (s *Service) Start() error {
	logType := "ShopServiceStart"
	s.startMu.Lock()
	defer s.startMu.Unlock()
	if s.started {
		logrus.Warnf("logType:%s, msg:already started, skip", logType)
		return nil
	}
	go s.consumeLoop()
	s.started = true
	logrus.Infof("logType:%s, msg:consume loop started", logType)
	return nil
}

/*Close
 *@Description: 关闭消费 goroutine 和依赖的 queue/workers
 *@receiver s
 *@return error 首个关闭错误(关闭过程不中断,尽量清理所有资源)
 */
func (s *Service) Close() error {
	logType := "ShopServiceClose"
	s.startMu.Lock()
	defer s.startMu.Unlock()
	if !s.started {
		return nil
	}
	s.started = false

	// 通知 consumeLoop 退出(目前 consumeLoop 依赖 queue.Consume 的 err 退出,
	// stopCh 作为兜底信号,后续如果 queue 支持 ctx 可改为 ctx cancel)
	close(s.stopCh)

	var firstErr error
	if err := s.deps.Queue.Close(); err != nil {
		logrus.Errorf("logType:%s, err:%s, step:close queue", logType, err.Error())
		firstErr = err
	}
	if err := s.deps.Workers.Close(); err != nil {
		logrus.Errorf("logType:%s, err:%s, step:close workers", logType, err.Error())
		if firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

/*consumeLoop
 *@Description: 从漏桶队列拉任务,转给 worker 池并发执行
 *   - 单 goroutine 搬运,协程数量可控
 *   - 外层 recover 兜底消费循环本身;每个 task 的 recover 已在 grab.worker.process 内
 *   - queue.Consume 返回 err 表示队列已关闭,自然退出
 */
func (s *Service) consumeLoop() {
	logType := "ShopServiceConsumeLoop"
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("logType:%s, panic:%v", logType, r)
		}
	}()
	for {
		select {
		case <-s.stopCh:
			logrus.Infof("logType:%s, msg:stop signal received", logType)
			return
		default:
		}
		task, err := s.deps.Queue.Consume()
		if err != nil {
			logrus.Infof("logType:%s, msg:queue closed, err:%s", logType, err.Error())
			return
		}
		if ok := s.deps.Workers.Push(task); !ok {
			logrus.Warnf("logType:%s, msg:worker pool closed, drop task and exit loop", logType)
			return
		}
	}
}
