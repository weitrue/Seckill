/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2024/10/06
 * Description: admin 侧 Activity 应用服务
 *   - 后台 CRUD + 上线/下线
 *   - 上线时顺便初始化 Redis 库存(通过 stockFactory)
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

// AdminService admin 侧活动应用服务
type AdminService struct {
	repo         activity.Repository
	stockFactory stock.Factory // 上线时初始化 Redis 库存
}

/*NewAdminService
 *@Description: 构造 AdminService
 *@param repo
 *@param stockFactory
 *@return *AdminService
 */
func NewAdminService(repo activity.Repository, stockFactory stock.Factory) *AdminService {
	return &AdminService{repo: repo, stockFactory: stockFactory}
}

/*Create
 *@Description: 创建活动
 *@receiver s
 *@param ctx
 *@param cmd
 *@return *activity.Activity 带 ID 的 Activity
 *@return error
 */
func (s *AdminService) Create(ctx context.Context, cmd CreateCommand) (*activity.Activity, error) {
	if cmd.Name == "" {
		return nil, errors.New("activity: name required")
	}
	if cmd.StartTime <= 0 || cmd.EndTime <= cmd.StartTime {
		return nil, activity.ErrInvalidTime
	}
	status := cmd.Status
	if status == 0 {
		status = activity.StatusPrepared
	}
	if !status.IsValid() {
		return nil, activity.ErrInvalidStatus
	}
	a := &activity.Activity{
		Name:      cmd.Name,
		StartTime: cmd.StartTime,
		EndTime:   cmd.EndTime,
		Status:    status,
	}
	a.Goods = make([]*activity.Goods, 0, len(cmd.Goods))
	for _, g := range cmd.Goods {
		a.Goods = append(a.Goods, &activity.Goods{
			GoodsID:       g.GoodsID,
			Name:          g.Name,
			Price:         g.Price,
			ActivityPrice: g.ActivityPrice,
			Stock:         g.Stock,
			LimitPerUser:  g.LimitPerUser,
		})
	}
	if err := s.repo.Create(ctx, a); err != nil {
		return nil, err
	}
	return a, nil
}

/*Update
 *@Description: 更新活动元数据
 *@receiver s
 *@param ctx
 *@param cmd
 *@return error
 */
func (s *AdminService) Update(ctx context.Context, cmd UpdateCommand) error {
	if cmd.ID <= 0 {
		return activity.ErrNotFound
	}
	if cmd.Status != 0 && !cmd.Status.IsValid() {
		return activity.ErrInvalidStatus
	}
	a := &activity.Activity{
		ID:        cmd.ID,
		Name:      cmd.Name,
		StartTime: cmd.StartTime,
		EndTime:   cmd.EndTime,
		Status:    cmd.Status,
	}
	return s.repo.Update(ctx, a)
}

/*Online
 *@Description: 上线活动;MySQL 状态置 StatusOnline + 遍历 goods 初始化 Redis 库存
 *@receiver s
 *@param ctx
 *@param id
 *@return error
 */
func (s *AdminService) Online(ctx context.Context, id int64) error {
	logType := "AdminActivityOnline"
	a, err := s.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if err = s.repo.UpdateStatus(ctx, id, activity.StatusOnline); err != nil {
		return err
	}
	if s.stockFactory == nil {
		return nil
	}
	aidStr := fmt.Sprintf("%d", a.ID)
	for _, g := range a.Goods {
		rs, factoryErr := s.stockFactory(aidStr, g.GoodsID)
		if factoryErr != nil {
			logrus.Errorf("logType:%s, err:%s, step:stockFactory, activityID:%d, goodsID:%s",
				logType, factoryErr.Error(), a.ID, g.GoodsID)
			continue
		}
		ttlSec := a.EndTime - time.Now().Unix() + 24*3600
		if ttlSec <= 0 {
			ttlSec = 3600
		}
		if err = rs.Set(ctx, g.Stock, ttlSec); err != nil {
			logrus.Errorf("logType:%s, err:%s, step:stock.Set, activityID:%d, goodsID:%s",
				logType, err.Error(), a.ID, g.GoodsID)
		}
	}
	return nil
}

/*Offline
 *@Description: 下线活动
 *@receiver s
 *@param ctx
 *@param id
 *@return error
 */
func (s *AdminService) Offline(ctx context.Context, id int64) error {
	if id <= 0 {
		return activity.ErrNotFound
	}
	return s.repo.UpdateStatus(ctx, id, activity.StatusEnded)
}
