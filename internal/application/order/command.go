/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2024/10/06
 * Description: Order 用例 DTO
 **/

package order

// CreateCommand 创建订单命令
type CreateCommand struct {
	UID        string
	ActivityID int64
	GoodsID    string
	Price      int64 // 分
}
