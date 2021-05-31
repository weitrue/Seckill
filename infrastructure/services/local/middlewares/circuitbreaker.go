/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午5:37
 * Description:
 **/

package middlewares

import (
	"net/http"

	"github.com/weitrue/Seckill/infrastructure/services/local/circuitbreaker"

	"github.com/gin-gonic/gin"
)

func NewCircuitBreakMiddleware(cb *circuitbreaker.CircuitBreaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		allow := cb.Allow(func() bool {
			ctx.Next()
			if ctx.Writer.Status() >= http.StatusInternalServerError {
				// 服务不可用时，返回请求处理结果为false
				return false
			}
			return true
		})
		if !allow {
			ctx.AbortWithStatus(http.StatusServiceUnavailable)
		}
	}
}
