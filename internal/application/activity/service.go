/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2024/10/06
 * Description: api 侧 Activity 应用服务
 *   - Get: 走 Repository(缓存优先,未命中回源 MySQL)
 *   - ListOngoing: 查进行中活动,供预热和 C 端列表接口
 *   - Prewarm: 服务启动预热,把 MySQL 的进行中活动 + 库存写入 Redis(供下单路径使用)
 **/

package activity

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/weitrue/Seckill/internal/domain/activity"
	"github.com/weitrue/Seckill/internal/domain/stock"
)

// Service api 侧活动应用服务
type Service struct {
	repo activity.Repository
	// stockFactory 用于预热时初始化 Redis 库存(通常传 redisstock.NewRedisStock)
	stockFactory stock.Factory
}

/*NewService
 *@Description: 构造 Service
 *@param repo
 *@param stockFactory
 *@return *Service
 */
func NewService(repo activity.Repository, stockFactory stock.Factory) *Service {
	return &Service{repo: repo, stockFactory: stockFactory}
}

/*Get
 *@Description: 按 ID 查活动
 *@receiver s
 *@param ctx
 *@param id
 *@return *activity.Activity
 *@return error
 */
func (s *Service) Get(ctx context.Context, id int64) (*activity.Activity, error) {
	if id <= 0 {
		return nil, activity.ErrNotFound
	}
	return s.repo.Get(ctx, id)
}

/*ListOngoing
 *@Description: 列出进行中/未结束的活动
 *@receiver s
 *@param ctx
 *@return []*activity.Activity
 *@return error
 */
func (s *Service) ListOngoing(ctx context.Context) ([]*activity.Activity, error) {
	return s.repo.ListOngoing(ctx, time.Now().Unix())
}

/*Prewarm
 *@Description: 启动预热
 *   1. 查 MySQL 中进行中的活动
 *   2. 遍历每个活动的商品,通过 stockFactory 初始化 Redis 库存
 *   3. 回源 Get 一次,触发 Redis 元数据回填(Repository.Get 内自带回填逻辑)
 *@receiver s
 *@param ctx
 *@return error 预热失败返回 err,调用方可选择不阻塞启动
 */
func (s *Service) Prewarm(ctx context.Context) error {
	logType := "ActivityPrewarm"
	if s.stockFactory == nil {
		return errors.New("activity.Service: stockFactory is nil, cannot prewarm")
	}
	list, err := s.repo.ListOngoing(ctx, time.Now().Unix())
	if err != nil {
		logrus.Errorf("logType:%s, err:%s, step:list ongoing", logType, err.Error())
		return err
	}
	logrus.Infof("logType:%s, msg:found ongoing activities, count:%d", logType, len(list))

	for _, a := range list {
		// 元数据回填(Get 内自带 cache 回填)
		aidStr := fmt.Sprintf("%d", a.ID)
		for _, g := range a.Goods {
			rs, factoryErr := s.stockFactory(aidStr, g.GoodsID)
			if factoryErr != nil {
				logrus.Errorf("logType:%s, err:%s, step:stockFactory, activityID:%d, goodsID:%s",
					logType, factoryErr.Error(), a.ID, g.GoodsID)
				continue
			}
			// Redis 库存写入;TTL 用活动剩余时间 + 1 天缓冲
			ttlSec := a.EndTime - time.Now().Unix() + 24*3600
			if ttlSec <= 0 {
				ttlSec = 3600
			}
			if err := rs.Set(ctx, g.Stock, ttlSec); err != nil {
				logrus.Errorf("logType:%s, err:%s, step:stock.Set, activityID:%d, goodsID:%s, stock:%d",
					logType, err.Error(), a.ID, g.GoodsID, g.Stock)
				continue
			}
			logrus.Infof("logType:%s, msg:stock prewarmed, activityID:%d, goodsID:%s, stock:%d, ttlSec:%d",
				logType, a.ID, g.GoodsID, g.Stock, ttlSec)
		}
		// Repository.Get 内部会回填 Redis 元数据缓存
		if _, err := s.repo.Get(ctx, a.ID); err != nil {
			logrus.Warnf("logType:%s, msg:meta cache fill failed, id:%d, err:%s",
				logType, a.ID, err.Error())
		}
	}
	return nil
}
