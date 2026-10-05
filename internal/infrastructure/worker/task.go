/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 下午10:55
 * Description:
 **/

package worker

type Task interface {
	Do() // 执行任务
}

type QueueTaskFunc func()

func (tf QueueTaskFunc) Do() {
	tf()
}
