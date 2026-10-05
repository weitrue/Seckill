/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 上午9:36
 * Description: 库存
 **/

package stock

type Stock interface {
	// Set 设置库存值
	Set(val int64, expire int64) error
	// Get 返回剩余库存
	Get() (int64, error)
	// Sub 减少库存
	Sub(uid string) (int64, error)
	// Del 删除库存
	Del() error
	// GetActivityID 返回活动ID
	GetActivityID() string
	// GetGoodsID 返回商品ID
	GetGoodsID() string
}
