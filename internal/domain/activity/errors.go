/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2024/10/06
 * Description: Activity 领域错误
 **/

package activity

import "errors"

var (
	// ErrNotFound 活动不存在
	ErrNotFound = errors.New("activity: not found")
	// ErrInvalidStatus 状态值不合法
	ErrInvalidStatus = errors.New("activity: invalid status")
	// ErrInvalidTime 时间区间不合法(start >= end 或 negative)
	ErrInvalidTime = errors.New("activity: invalid time range")
	// ErrGoodsNotInActivity 商品不在活动内
	ErrGoodsNotInActivity = errors.New("activity: goods not in activity")
	// ErrNotOngoing 活动不在进行中
	ErrNotOngoing = errors.New("activity: not ongoing")
)
