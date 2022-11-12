/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午5:37
 * Description:
 **/

package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/weitrue/Seckill/infrastructure/services/local/circuitbreaker"
)

// CircuitBreakMiddleware 熔断器中间件
func CircuitBreakMiddleware(cb *circuitbreaker.CircuitBreaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		allow := cb.Allow(func() bool {
			ctx.Next()
			// 服务不可用时，返回请求处理结果为false
			return ctx.Writer.Status() < http.StatusInternalServerError
		})
		if !allow {
			ctx.AbortWithStatus(http.StatusServiceUnavailable)
		}
	}
}
