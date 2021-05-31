/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 下午6:27
 * Description:
 **/

package cachei

type IntCache interface {
	Cache
	// 缓存中value为int 黑白名单 库存信息用户抢购次数等
	Get(key string) (int64, bool)
	Set(key string, val int64)
	Add(key string, delta int64) int64
}
