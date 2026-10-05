/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2024/10/06
 * Description: Activity 仓储的 MySQL 实现(真实源)
 *   - 表: activity + activity_goods(SQL schema 见 docker/mysql/init/001_schema.sql)
 *   - 聚合根与 PO 之间手动转换,避免 gorm tag 污染 domain 层
 **/

package activityrepo

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/weitrue/Seckill/internal/domain/activity"
)

// activityPO activity 表 PO
type activityPO struct {
	ID         int64  `gorm:"column:id;primaryKey;autoIncrement"`
	Name       string `gorm:"column:name"`
	StartTime  int64  `gorm:"column:start_time"`
	EndTime    int64  `gorm:"column:end_time"`
	Status     int8   `gorm:"column:status"`
	CreateTime int64  `gorm:"column:create_time;autoCreateTime"`
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime"`
}

// TableName gorm 表名约定
func (activityPO) TableName() string { return "activity" }

// goodsPO activity_goods 表 PO
type goodsPO struct {
	ID            int64  `gorm:"column:id;primaryKey;autoIncrement"`
	ActivityID    int64  `gorm:"column:activity_id"`
	GoodsID       string `gorm:"column:goods_id"`
	Name          string `gorm:"column:name"`
	Price         int64  `gorm:"column:price"`
	ActivityPrice int64  `gorm:"column:activity_price"`
	Stock         int64  `gorm:"column:stock"`
	LimitPerUser  int    `gorm:"column:limit_per_user"`
	CreateTime    int64  `gorm:"column:create_time;autoCreateTime"`
	UpdateTime    int64  `gorm:"column:update_time;autoUpdateTime"`
}

// TableName gorm 表名约定
func (goodsPO) TableName() string { return "activity_goods" }

// mysqlRepo activity Repository 的 MySQL 实现(私有,通过组合层对外)
type mysqlRepo struct {
	db *gorm.DB
}

func newMySQLRepo(db *gorm.DB) *mysqlRepo {
	return &mysqlRepo{db: db}
}

// Create 创建活动(含商品);事务保证原子
func (r *mysqlRepo) Create(ctx context.Context, a *activity.Activity) error {
	if a == nil {
		return errors.New("activityrepo.mysql: nil activity")
	}
	now := time.Now().Unix()
	a.CreateTime = now
	a.UpdateTime = now
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		po := toActivityPO(a)
		if err := tx.Create(&po).Error; err != nil {
			return err
		}
		a.ID = po.ID
		if len(a.Goods) == 0 {
			return nil
		}
		gpos := make([]goodsPO, 0, len(a.Goods))
		for _, g := range a.Goods {
			g.ActivityID = po.ID
			g.CreateTime = now
			g.UpdateTime = now
			gpos = append(gpos, toGoodsPO(g))
		}
		if err := tx.Create(&gpos).Error; err != nil {
			return err
		}
		// 回填 ID
		for i := range a.Goods {
			a.Goods[i].ID = gpos[i].ID
			a.Goods[i].ActivityID = po.ID
		}
		return nil
	})
}

// Update 更新活动元数据(name / start / end / status),不改 goods 清单
func (r *mysqlRepo) Update(ctx context.Context, a *activity.Activity) error {
	if a == nil || a.ID <= 0 {
		return errors.New("activityrepo.mysql: invalid activity id")
	}
	return r.db.WithContext(ctx).Model(&activityPO{}).
		Where("id = ?", a.ID).
		Updates(map[string]interface{}{
			"name":        a.Name,
			"start_time":  a.StartTime,
			"end_time":    a.EndTime,
			"status":      int8(a.Status),
			"update_time": time.Now().Unix(),
		}).Error
}

// UpdateStatus 变更状态
func (r *mysqlRepo) UpdateStatus(ctx context.Context, id int64, status activity.Status) error {
	if id <= 0 || !status.IsValid() {
		return activity.ErrInvalidStatus
	}
	return r.db.WithContext(ctx).Model(&activityPO{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":      int8(status),
			"update_time": time.Now().Unix(),
		}).Error
}

// Get 查活动(含商品)
func (r *mysqlRepo) Get(ctx context.Context, id int64) (*activity.Activity, error) {
	if id <= 0 {
		return nil, activity.ErrNotFound
	}
	var apo activityPO
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&apo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, activity.ErrNotFound
		}
		return nil, err
	}
	var gpos []goodsPO
	if err = r.db.WithContext(ctx).Where("activity_id = ?", id).Find(&gpos).Error; err != nil {
		return nil, err
	}
	return fromActivityPO(&apo, gpos), nil
}

// ListOngoing 查"结束时间 >= now 且状态上线"的活动,用于启动预热
func (r *mysqlRepo) ListOngoing(ctx context.Context, nowSec int64) ([]*activity.Activity, error) {
	var apos []activityPO
	err := r.db.WithContext(ctx).
		Where("status = ? AND end_time >= ?", int8(activity.StatusOnline), nowSec).
		Find(&apos).Error
	if err != nil {
		return nil, err
	}
	if len(apos) == 0 {
		return nil, nil
	}
	// 批量取所有 activity 的 goods,避免循环查
	ids := make([]int64, 0, len(apos))
	for i := range apos {
		ids = append(ids, apos[i].ID)
	}
	var gpos []goodsPO
	if err = r.db.WithContext(ctx).Where("activity_id IN ?", ids).Find(&gpos).Error; err != nil {
		return nil, err
	}
	// 按 activity_id 分桶
	goodsByAct := make(map[int64][]goodsPO, len(apos))
	for _, g := range gpos {
		goodsByAct[g.ActivityID] = append(goodsByAct[g.ActivityID], g)
	}
	out := make([]*activity.Activity, 0, len(apos))
	for i := range apos {
		out = append(out, fromActivityPO(&apos[i], goodsByAct[apos[i].ID]))
	}
	return out, nil
}

// --- PO <-> 聚合 转换 ---

func toActivityPO(a *activity.Activity) activityPO {
	return activityPO{
		ID:         a.ID,
		Name:       a.Name,
		StartTime:  a.StartTime,
		EndTime:    a.EndTime,
		Status:     int8(a.Status),
		CreateTime: a.CreateTime,
		UpdateTime: a.UpdateTime,
	}
}

func toGoodsPO(g *activity.Goods) goodsPO {
	return goodsPO{
		ID:            g.ID,
		ActivityID:    g.ActivityID,
		GoodsID:       g.GoodsID,
		Name:          g.Name,
		Price:         g.Price,
		ActivityPrice: g.ActivityPrice,
		Stock:         g.Stock,
		LimitPerUser:  g.LimitPerUser,
		CreateTime:    g.CreateTime,
		UpdateTime:    g.UpdateTime,
	}
}

func fromActivityPO(apo *activityPO, gpos []goodsPO) *activity.Activity {
	a := &activity.Activity{
		ID:         apo.ID,
		Name:       apo.Name,
		StartTime:  apo.StartTime,
		EndTime:    apo.EndTime,
		Status:     activity.Status(apo.Status),
		CreateTime: apo.CreateTime,
		UpdateTime: apo.UpdateTime,
	}
	a.Goods = make([]*activity.Goods, 0, len(gpos))
	for _, g := range gpos {
		a.Goods = append(a.Goods, &activity.Goods{
			ID:            g.ID,
			ActivityID:    g.ActivityID,
			GoodsID:       g.GoodsID,
			Name:          g.Name,
			Price:         g.Price,
			ActivityPrice: g.ActivityPrice,
			Stock:         g.Stock,
			LimitPerUser:  g.LimitPerUser,
			CreateTime:    g.CreateTime,
			UpdateTime:    g.UpdateTime,
		})
	}
	return a
}
