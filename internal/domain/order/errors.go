/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2024/10/06
 * Description: Order 领域错误
 **/

package order

import "errors"

var (
	// ErrNotFound 订单不存在
	ErrNotFound = errors.New("order: not found")
	// ErrDuplicate 幂等命中(用户已在该活动+商品下单)
	ErrDuplicate = errors.New("order: duplicate uid+activity+goods")
	// ErrInvalidStatus 状态值不合法
	ErrInvalidStatus = errors.New("order: invalid status")
)
