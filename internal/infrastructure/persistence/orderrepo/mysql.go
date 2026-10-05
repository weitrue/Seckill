/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2024/10/06
 * Description: Order 仓储的 MySQL 实现
 *   - 表: seckill_order(SQL schema 见 docker/mysql/init/001_schema.sql)
 *   - uk_user_activity_goods 唯一键做兜底幂等,Duplicate 返回 order.ErrDuplicate
 **/

package orderrepo

import (
	"context"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/weitrue/Seckill/internal/domain/order"
)

// orderPO seckill_order 表 PO
type orderPO struct {
	ID         int64  `gorm:"column:id;primaryKey;autoIncrement"`
	UID        string `gorm:"column:uid"`
	ActivityID int64  `gorm:"column:activity_id"`
	GoodsID    string `gorm:"column:goods_id"`
	Price      int64  `gorm:"column:price"`
	Status     int8   `gorm:"column:status"`
	CreateTime int64  `gorm:"column:create_time"`
	UpdateTime int64  `gorm:"column:update_time"`
}

// TableName gorm 表名约定
func (orderPO) TableName() string { return "seckill_order" }

// Repository Order 仓储 MySQL 实现
type Repository struct {
	db *gorm.DB
}

// 编译期类型断言
var _ order.Repository = (*Repository)(nil)

/*New
 *@Description: 构造 Repository
 *@param db *gorm.DB
 *@return *Repository
 */
func New(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

/*Create
 *@Description: 创建订单;唯一键冲突返回 order.ErrDuplicate
 *@receiver r
 *@param ctx
 *@param o
 *@return error
 */
func (r *Repository) Create(ctx context.Context, o *order.Order) error {
	if o == nil {
		return errors.New("orderrepo: nil order")
	}
	now := time.Now().Unix()
	if o.CreateTime == 0 {
		o.CreateTime = now
	}
	o.UpdateTime = now
	if !o.Status.IsValid() {
		o.Status = order.StatusPending
	}
	po := toPO(o)
	if err := r.db.WithContext(ctx).Create(&po).Error; err != nil {
		if isDuplicateErr(err) {
			return order.ErrDuplicate
		}
		return err
	}
	o.ID = po.ID
	return nil
}

/*Get
 *@Description: 按 ID 查订单
 *@receiver r
 *@param ctx
 *@param id
 *@return *order.Order
 *@return error
 */
func (r *Repository) Get(ctx context.Context, id int64) (*order.Order, error) {
	if id <= 0 {
		return nil, order.ErrNotFound
	}
	var po orderPO
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&po).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, order.ErrNotFound
		}
		return nil, err
	}
	return fromPO(&po), nil
}

/*GetByUserGoods
 *@Description: 按 uid+activity+goods 查订单,用于幂等检查
 *@receiver r
 *@param ctx
 *@param uid
 *@param activityID
 *@param goodsID
 *@return *order.Order
 *@return error
 */
func (r *Repository) GetByUserGoods(ctx context.Context, uid string, activityID int64, goodsID string) (*order.Order, error) {
	if uid == "" || activityID <= 0 || goodsID == "" {
		return nil, order.ErrNotFound
	}
	var po orderPO
	err := r.db.WithContext(ctx).
		Where("uid = ? AND activity_id = ? AND goods_id = ?", uid, activityID, goodsID).
		First(&po).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, order.ErrNotFound
		}
		return nil, err
	}
	return fromPO(&po), nil
}

// isDuplicateErr 识别 MySQL 唯一键冲突(1062 Duplicate entry)
//
//	go-sql-driver 返回的 *mysql.MySQLError.Number 是 1062,但为了不引入强依赖,
//	这里用错误消息判断("Error 1062")
func isDuplicateErr(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	return strings.Contains(msg, "1062") || strings.Contains(msg, "Duplicate entry")
}

// --- PO <-> 聚合 转换 ---

func toPO(o *order.Order) orderPO {
	return orderPO{
		ID:         o.ID,
		UID:        o.UID,
		ActivityID: o.ActivityID,
		GoodsID:    o.GoodsID,
		Price:      o.Price,
		Status:     int8(o.Status),
		CreateTime: o.CreateTime,
		UpdateTime: o.UpdateTime,
	}
}

func fromPO(po *orderPO) *order.Order {
	return &order.Order{
		ID:         po.ID,
		UID:        po.UID,
		ActivityID: po.ActivityID,
		GoodsID:    po.GoodsID,
		Price:      po.Price,
		Status:     order.Status(po.Status),
		CreateTime: po.CreateTime,
		UpdateTime: po.UpdateTime,
	}
}
