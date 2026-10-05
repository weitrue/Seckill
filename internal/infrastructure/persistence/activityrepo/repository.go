/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2024/10/06
 * Description: Activity 仓储组合层
 *   - MySQL 作为真实源
 *   - Redis Hash 作为预热副本/缓存
 *   - 读: 先读缓存,未命中回源 MySQL 并异步/同步回填缓存
 *   - 写: MySQL 写入成功后刷新缓存(写穿透);删除/更新同样
 *   - 实现 domain/activity.Repository
 **/

package activityrepo

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/weitrue/Seckill/internal/domain/activity"
)

// Repository 组合 MySQL + Redis 的 Activity 仓储
type Repository struct {
	mysql *mysqlRepo
	cache *cache
}

// 编译期类型断言: Repository 实现 activity.Repository
var _ activity.Repository = (*Repository)(nil)

/*New
 *@Description: 构造 Repository
 *@param db gorm.DB(本项目通过 infrastructure/stores/mysql.GetDB() 注入)
 *@param client *redis.Client(通过 infrastructure/stores/redis.GetRedisClient(db) 注入)
 *@param cacheTTL 缓存过期时间;<=0 用默认 7d
 *@return *Repository
 */
func New(db *gorm.DB, client *redis.Client, cacheTTL time.Duration) *Repository {
	return &Repository{
		mysql: newMySQLRepo(db),
		cache: newCache(client, cacheTTL),
	}
}

/*Create
 *@Description: 创建活动;MySQL 写入 → 刷新 Redis 缓存
 *@receiver r
 *@param ctx
 *@param a
 *@return error
 */
func (r *Repository) Create(ctx context.Context, a *activity.Activity) error {
	if err := r.mysql.Create(ctx, a); err != nil {
		return err
	}
	if err := r.cache.Set(ctx, a); err != nil {
		// 缓存失败不回滚 MySQL,仅记日志(后续读会回源 MySQL 并回填)
		logrus.Warnf("logType:ActivityRepoCreate, msg:cache set failed, id:%d, err:%s",
			a.ID, err.Error())
	}
	return nil
}

/*Update
 *@Description: 更新活动;MySQL 写入成功后主动删除缓存(或刷新)
 *@receiver r
 *@param ctx
 *@param a
 *@return error
 */
func (r *Repository) Update(ctx context.Context, a *activity.Activity) error {
	if err := r.mysql.Update(ctx, a); err != nil {
		return err
	}
	// 直接从 MySQL 读完整数据(含商品)后刷缓存
	full, err := r.mysql.Get(ctx, a.ID)
	if err != nil {
		logrus.Warnf("logType:ActivityRepoUpdate, msg:refresh cache failed on reload, id:%d, err:%s",
			a.ID, err.Error())
		_ = r.cache.Del(ctx, a.ID)
		return nil
	}
	if err = r.cache.Set(ctx, full); err != nil {
		logrus.Warnf("logType:ActivityRepoUpdate, msg:cache set failed, id:%d, err:%s",
			a.ID, err.Error())
	}
	return nil
}

/*UpdateStatus
 *@Description: 变更状态;MySQL 写入成功后删除缓存
 *@receiver r
 *@param ctx
 *@param id
 *@param status
 *@return error
 */
func (r *Repository) UpdateStatus(ctx context.Context, id int64, status activity.Status) error {
	if err := r.mysql.UpdateStatus(ctx, id, status); err != nil {
		return err
	}
	if err := r.cache.Del(ctx, id); err != nil {
		logrus.Warnf("logType:ActivityRepoUpdateStatus, msg:cache del failed, id:%d, err:%s",
			id, err.Error())
	}
	return nil
}

/*Get
 *@Description: 查活动;先读缓存,未命中回源 MySQL 并回填缓存
 *@receiver r
 *@param ctx
 *@param id
 *@return *activity.Activity
 *@return error
 */
func (r *Repository) Get(ctx context.Context, id int64) (*activity.Activity, error) {
	if a, err := r.cache.Get(ctx, id); err != nil {
		// 缓存出错不阻塞,记日志回源
		logrus.Warnf("logType:ActivityRepoGet, msg:cache get failed, id:%d, err:%s",
			id, err.Error())
	} else if a != nil {
		return a, nil
	}
	// 回源
	a, err := r.mysql.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	// 回填缓存
	if cacheErr := r.cache.Set(ctx, a); cacheErr != nil {
		logrus.Warnf("logType:ActivityRepoGet, msg:cache fill failed, id:%d, err:%s",
			id, cacheErr.Error())
	}
	return a, nil
}

/*ListOngoing
 *@Description: 查进行中/即将结束的活动;直接走 MySQL(需要范围查询,缓存不适合)
 *@receiver r
 *@param ctx
 *@param nowSec
 *@return []*activity.Activity
 *@return error
 */
func (r *Repository) ListOngoing(ctx context.Context, nowSec int64) ([]*activity.Activity, error) {
	return r.mysql.ListOngoing(ctx, nowSec)
}
