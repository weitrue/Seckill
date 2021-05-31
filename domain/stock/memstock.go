/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/10 上午9:49
 * Description:
 **/

package stock

import (
	"errors"
	"fmt"

	"github.com/weitrue/Seckill/domain/stock/stocki"
	"github.com/weitrue/Seckill/infrastructure/utils"
)

type memoryStack struct {
	activityID string // 活动ID
	goodsID    string // 商品ID
	key        string //库存key
}

func NewMemoryStock(activityID, goodsID string) (stocki.Stock, error) {
	if activityID == "" || goodsID == "" {
		return nil, errors.New("invalid event id or goods id")
	}
	oneStack := &memoryStack{
		activityID: activityID,
		goodsID:    goodsID,
		key:        fmt.Sprintf("%s:%s:%s", utils.GetServiceName(), activityID, goodsID),
	}
	return oneStack, nil
}

func (m memoryStack) Set(val int64, expire int64) error {
	panic("implement me")
}

func (m memoryStack) Get() (int64, error) {
	panic("implement me")
}

func (m memoryStack) Sub(uid string) (int64, error) {
	panic("implement me")
}

func (m memoryStack) Del() error {
	panic("implement me")
}

func (m memoryStack) GetActivityID() string {
	panic("implement me")
}

func (m memoryStack) GetGoodsID() string {
	panic("implement me")
}