/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2024/10/06
 * Description: Activity 聚合
 *   - Activity 是秒杀活动的聚合根,内部组合多个活动商品(Goods)
 *   - 领域规则: IsOngoing(当前时间在活动期内 + 已上线) / FindGoods(查商品是否在活动内)
 *   - 持久化在 MySQL(activity + activity_goods),Redis Hash 作为预热副本/缓存
 **/

package activity

import "time"

// Status 活动状态
type Status int8

const (
	StatusPrepared Status = 1 // 筹备中(未上线)
	StatusOnline   Status = 2 // 上线/进行中
	StatusEnded    Status = 3 // 已结束
)

// IsValid 判断状态是否合法
func (s Status) IsValid() bool {
	switch s {
	case StatusPrepared, StatusOnline, StatusEnded:
		return true
	}
	return false
}

// Activity 秒杀活动聚合根
type Activity struct {
	ID         int64    // 活动 ID
	Name       string   // 活动名称
	StartTime  int64    // 开始时间(秒)
	EndTime    int64    // 结束时间(秒)
	Status     Status   // 状态
	Goods      []*Goods // 活动商品(聚合内实体)
	CreateTime int64
	UpdateTime int64
}

// Goods 活动商品(Activity 聚合内部实体)
type Goods struct {
	ID            int64  // 自增 ID
	ActivityID    int64  // 活动 ID
	GoodsID       string // 商品 ID(业务侧字符串标识)
	Name          string // 商品名
	Price         int64  // 原价(分)
	ActivityPrice int64  // 活动价(分)
	Stock         int64  // 秒杀库存
	LimitPerUser  int    // 单用户限购数
	CreateTime    int64
	UpdateTime    int64
}

/*IsOngoing
 *@Description: 判断当前时间是否在活动期内(状态为上线 + 当前时间在 [start, end) 内)
 *@receiver a
 *@param now
 *@return bool
 */
func (a *Activity) IsOngoing(now time.Time) bool {
	if a == nil {
		return false
	}
	sec := now.Unix()
	return a.Status == StatusOnline && a.StartTime <= sec && sec < a.EndTime
}

/*FindGoods
 *@Description: 按业务侧 goodsID 查找活动内商品;找不到返回 nil
 *@receiver a
 *@param goodsID
 *@return *Goods
 */
func (a *Activity) FindGoods(goodsID string) *Goods {
	if a == nil {
		return nil
	}
	for _, g := range a.Goods {
		if g.GoodsID == goodsID {
			return g
		}
	}
	return nil
}
