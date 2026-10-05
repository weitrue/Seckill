/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 上午9:36
 * Description: 库存领域接口
 *   所有方法接受 context.Context,为链路追踪 / 超时控制 / 请求级日志等预留入口
 *   具体实现放在 internal/infrastructure/persistence/stockrepo/{memstock,redisstock}
 **/

package stock

import "context"

// Stock 库存接口
type Stock interface {
	// Set 设置库存值
	Set(ctx context.Context, val int64, expire int64) error
	// Get 返回剩余库存
	Get(ctx context.Context) (int64, error)
	// Sub 扣减库存(传入请求用户 uid,内部可做幂等/去重判断)
	Sub(ctx context.Context, uid string) (int64, error)
	// Del 删除库存
	Del(ctx context.Context) error
	// GetActivityID 返回活动ID
	GetActivityID() string
	// GetGoodsID 返回商品ID
	GetGoodsID() string
}

// Factory 库存实例工厂签名
//
//	application 层通过该签名注入具体实现(内存预扣 / Redis 真扣),
//	实现侧(infrastructure/persistence/stockrepo/*)只需提供符合该签名的函数或 method value
type Factory func(activityID, goodsID string) (Stock, error)
