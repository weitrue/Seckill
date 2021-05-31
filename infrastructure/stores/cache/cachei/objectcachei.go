/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 下午6:25
 * Description:
 **/

package cachei

type ObjCache interface {
	Cache
	// 缓存中value为结构体 活动信息
	Get(key string) (interface{}, bool)
	Set(key string, val interface{})
}
