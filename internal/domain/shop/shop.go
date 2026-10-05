/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午8:32
 * Description: shop 领域包
 *   目前只作为领域包占位,后续阶段补齐:
 *     - Order 聚合(id/uid/activityID/goodsID/status/createAt + Create 规则)
 *     - 下单领域错误 / 领域事件(OrderPlaced / OrderCanceled)
 *     - OrderRepository 接口(实现放 infrastructure/persistence/orderrepo)
 *
 *   注意:
 *     - 协议细节(gin.Context/net.Conn/http.Response)已搬到 interfaces/api/handler
 *     - 队列/worker/消费循环已搬到 internal/application/shop
 *     - 业务码/DTO 定义在 internal/application/shop(command.go)
 **/

package shop
