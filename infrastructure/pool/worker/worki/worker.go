/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 下午10:51
 * Description:
 **/

package worki

import "github.com/weitrue/Seckill/infrastructure/pool/worker/taski"

type Worker interface {
	Push(t taski.Task) bool // 推送任务
	Close() error           // 关闭协程池
}
