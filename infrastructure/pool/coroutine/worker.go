/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 下午10:51
 * Description:
 **/

package coroutine

type Worker interface {
	Push(t Task) bool // 推送任务
	Close() error     // 关闭协程池
}
