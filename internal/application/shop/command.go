/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2024/10/06
 * Description: shop 用例的输入/输出 DTO 与业务码定义
 *   DTO 与 handler 的入参 JSON 结构解耦:handler 负责 bind HTTP body 后转换成 Command
 **/

package shop

// 业务结果码: 对接响应 json 的 code 字段
const (
	CodeOK                 = 0    // 成功
	CodeNoStock            = 1001 // 库存不足 / 已参与过
	CodeRedisErr           = 1002 // Redis 异常
	CodeTimeout            = 1003 // 超时 / 请求被取消
	CodeNoActivity         = 1004 // 活动不存在
	CodeActivityNotOngoing = 1005 // 活动未进行中
	CodeNoGoods            = 1006 // 商品不在活动内
	CodeOrderDuplicate     = 1007 // 已下过单(幂等)
	CodeOrderErr           = 1008 // 订单落库异常
	CodeBadRequest         = 1009 // 入参不合法
)

// AddCartCommand 下单(加购)命令
//
//	由 handler 层从鉴权上下文和 request body 中抽取拼装,传入 Service
type AddCartCommand struct {
	UID        string // 用户 ID(鉴权中间件注入)
	ActivityID string // 活动 ID(字符串形式,Service 内部转 int64 查 Activity)
	GoodsID    string // 商品 ID
}

// AddCartResult 下单(加购)结果
//
//	Service 的返回值,handler 层直接作为 json body 吐给前端
type AddCartResult struct {
	Code    int    `json:"code"`              // 业务码,见上方常量
	Msg     string `json:"msg"`               // 业务描述
	OrderID int64  `json:"order_id,omitempty"` // 下单成功时返回订单 ID
	Data    any    `json:"data,omitempty"`     // 预留扩展
}
