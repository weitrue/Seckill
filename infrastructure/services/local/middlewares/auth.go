/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 上午11:33
 * Description:
 **/

package middlewares

import (
	"Seckill/domain/user"
	"Seckill/infrastructure/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func NewAuthMiddleware(redirect bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var info *user.Info
		// 获取请求头token
		token := ctx.Request.Header.Get(user.TokenHeader)
		if token != "" && strings.Contains(token, user.TokenPrefix) {
			token = strings.Trim(token, user.TokenPrefix)
			token = strings.TrimSpace(token)
			// 校验信息
			info = user.Auth(token)
		}
		if info == nil {
			ctx.Set("UserInfo", info)
		} else if redirect {
			utils.Abort(ctx, http.StatusUnauthorized, "need login")
			return
		}
		ctx.Next()
	}
}
