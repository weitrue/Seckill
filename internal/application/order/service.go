/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2024/10/06
 * Description: Order 应用服务
 *   - Create: 落库;遇到唯一键冲突返回 order.ErrDuplicate(调用方按幂等处理)
 **/

package order

import (
	"context"
	"errors"

	"github.com/weitrue/Seckill/internal/domain/order"
)

// Service 订单应用服务
type Service struct {
	repo order.Repository
}

/*NewService
 *@Description: 构造 Service
 *@param repo
 *@return *Service
 */
func NewService(repo order.Repository) *Service {
	return &Service{repo: repo}
}

/*Create
 *@Description: 创建订单
 *@receiver s
 *@param ctx
 *@param cmd
 *@return *order.Order 带 ID
 *@return error ErrDuplicate 表示幂等命中
 */
func (s *Service) Create(ctx context.Context, cmd CreateCommand) (*order.Order, error) {
	if cmd.UID == "" || cmd.ActivityID <= 0 || cmd.GoodsID == "" {
		return nil, errors.New("order: invalid create command")
	}
	o := &order.Order{
		UID:        cmd.UID,
		ActivityID: cmd.ActivityID,
		GoodsID:    cmd.GoodsID,
		Price:      cmd.Price,
		Status:     order.StatusPending,
	}
	if err := s.repo.Create(ctx, o); err != nil {
		return nil, err
	}
	return o, nil
}

/*Get
 *@Description: 按 ID 查订单
 *@receiver s
 *@param ctx
 *@param id
 *@return *order.Order
 *@return error
 */
func (s *Service) Get(ctx context.Context, id int64) (*order.Order, error) {
	return s.repo.Get(ctx, id)
}
