/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午5:37
 * Description: 熔断器中间件
 *   熔断算法本身留在 infrastructure/services/local/circuitbreaker,这里只做 gin wrapper
 **/

package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/weitrue/Seckill/internal/infrastructure/services/local/circuitbreaker"
)

/*CircuitBreak
 *@Description: 熔断中间件
 *   每次请求通过 cb.Allow 包裹 c.Next(),返回状态码 >= 500 视为失败,熔断器据此统计
 *   拒绝时返回 503
 *@param cb 熔断器实例
 *@return gin.HandlerFunc
 */
func CircuitBreak(cb *circuitbreaker.CircuitBreaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		allow := cb.Allow(func() bool {
			c.Next()
			return c.Writer.Status() < http.StatusInternalServerError
		})
		if !allow {
			c.AbortWithStatus(http.StatusServiceUnavailable)
		}
	}
}
