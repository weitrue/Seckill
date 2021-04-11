/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 下午11:02
 * Description:
 **/

package coroutine

type PriorityTask interface {
	Priority() int
	Task
}
