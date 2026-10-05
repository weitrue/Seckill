/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2024/10/06
 * Description: Activity 用例的输入 DTO
 **/

package activity

import "github.com/weitrue/Seckill/internal/domain/activity"

// CreateCommand 后台创建活动命令
type CreateCommand struct {
	Name      string         // 活动名
	StartTime int64          // 开始(秒)
	EndTime   int64          // 结束(秒)
	Status    activity.Status // 初始状态(通常 StatusPrepared)
	Goods     []GoodsItem    // 活动商品
}

// GoodsItem 创建命令里的商品项
type GoodsItem struct {
	GoodsID       string
	Name          string
	Price         int64
	ActivityPrice int64
	Stock         int64
	LimitPerUser  int
}

// UpdateCommand 后台更新活动元数据命令
type UpdateCommand struct {
	ID        int64
	Name      string
	StartTime int64
	EndTime   int64
	Status    activity.Status
}
