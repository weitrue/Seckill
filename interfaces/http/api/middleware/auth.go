/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 上午11:33
 * Description: 鉴权中间件
 *   - 从 HTTP header Authorization 取 "Bearer <token>" 格式的 token
 *   - 调 infrastructure/auth.Verify 解出 *user.Info
 *   - 成功: 塞到 ctx 的 "userInfo" key(key 与 handler 侧的 c.Get 一致)
 *   - 失败:
 *       required = true  → Abort 401 "need login"
 *       required = false → 放行(匿名访问)
 **/

package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/weitrue/Seckill/internal/infrastructure/auth"
	"github.com/weitrue/Seckill/pkg/utils"
)

// HTTP 鉴权相关常量
const (
	// TokenHeader 鉴权 header 名
	TokenHeader = "Authorization"
	// TokenPrefix Bearer 前缀(空格后面才是真正的 token)
	TokenPrefix = "Bearer "
	// CtxKeyUserInfo gin.Context 中保存 *user.Info 的 key,与 handler 侧 c.Get("userInfo") 对齐
	CtxKeyUserInfo = "userInfo"
)

/*Auth
 *@Description: 鉴权中间件
 *@param required 是否强制登录;true 则未登录 Abort 401,false 则允许匿名放行
 *@return gin.HandlerFunc
 */
func Auth(required bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		logType := "MiddlewareAuth"
		token := extractToken(c)
		if token == "" {
			if required {
				utils.Abort(c, http.StatusUnauthorized, "need login")
				return
			}
			c.Next()
			return
		}

		info, err := auth.Verify(c.Request.Context(), token)
		if err != nil || info == nil {
			if required {
				logrus.Warnf("logType:%s, msg:verify failed, err:%v, path:%s",
					logType, err, c.Request.URL.Path)
				utils.Abort(c, http.StatusUnauthorized, "need login")
				return
			}
			c.Next()
			return
		}

		c.Set(CtxKeyUserInfo, info)
		c.Next()
	}
}

// extractToken 从 Authorization header 抽出去掉 Bearer 前缀的 token
func extractToken(c *gin.Context) string {
	raw := strings.TrimSpace(c.Request.Header.Get(TokenHeader))
	if raw == "" {
		return ""
	}
	// 大小写不敏感匹配 Bearer 前缀
	if len(raw) >= len(TokenPrefix) &&
		strings.EqualFold(raw[:len(TokenPrefix)], TokenPrefix) {
		return strings.TrimSpace(raw[len(TokenPrefix):])
	}
	return raw
}
