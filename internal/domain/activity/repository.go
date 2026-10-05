/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2024/10/06
 * Description: Activity 仓储接口
 *   - 实现放在 infrastructure/persistence/activityrepo(组合层: MySQL 作为源 + Redis Hash 缓存)
 **/

package activity

import "context"

// Repository 活动仓储接口
type Repository interface {
	// Create 创建活动(含商品),返回带 ID 的 Activity
	Create(ctx context.Context, a *Activity) error
	// Update 更新活动元数据(name/start/end 等);goods 清单的更新建议走单独方法,先保持最小
	Update(ctx context.Context, a *Activity) error
	// UpdateStatus 变更状态(上线/下线等)
	UpdateStatus(ctx context.Context, id int64, status Status) error
	// Get 根据 ID 查活动(含商品)
	Get(ctx context.Context, id int64) (*Activity, error)
	// ListOngoing 查询"进行中 + 即将开始"的活动,用于启动预热
	//
	//	withinSec: 结束时间不小于当前秒;可以把 start_time > now 的"即将开始"也包含进来
	ListOngoing(ctx context.Context, nowSec int64) ([]*Activity, error)
}
