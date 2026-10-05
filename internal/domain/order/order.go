/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2024/10/06
 * Description: Order 聚合
 *   - Order 是下单结果的聚合根
 *   - MySQL 唯一键 (uid, activity_id, goods_id) 做兜底幂等,
 *     Redis Lua 的 history 字段是第一道去重,Order 落库出现 duplicate 视为"已下单"
 **/

package order

// Status 订单状态
type Status int8

const (
	StatusPending  Status = 1 // 待支付
	StatusPaid     Status = 2 // 已支付
	StatusCanceled Status = 3 // 已取消
)

// IsValid 判断状态是否合法
func (s Status) IsValid() bool {
	switch s {
	case StatusPending, StatusPaid, StatusCanceled:
		return true
	}
	return false
}

// Order 秒杀订单聚合根
type Order struct {
	ID         int64
	UID        string
	ActivityID int64
	GoodsID    string
	Price      int64 // 分
	Status     Status
	CreateTime int64
	UpdateTime int64
}
