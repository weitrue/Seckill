/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2024/10/06
 * Description: Order 仓储接口
 **/

package order

import "context"

// Repository 订单仓储接口
type Repository interface {
	// Create 创建订单;碰到 (uid, activity_id, goods_id) 唯一键冲突返回 ErrDuplicate
	Create(ctx context.Context, o *Order) error
	// Get 按 ID 查订单
	Get(ctx context.Context, id int64) (*Order, error)
	// GetByUserGoods 按 uid+activity+goods 查订单,用于幂等检查
	GetByUserGoods(ctx context.Context, uid string, activityID int64, goodsID string) (*Order, error)
}
