/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午12:11
 * Description: 黑名单中间件
 *   必须放在 Auth(true) 之后使用,依赖 ctx 里的 userInfo
 **/

package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/weitrue/Seckill/internal/domain/user"
	"github.com/weitrue/Seckill/internal/infrastructure/config"
	"github.com/weitrue/Seckill/pkg/utils"
)

/*Blacklist
 *@Description: 黑名单检查
 *   ctx 没有 userInfo(未登录) → 401
 *   命中黑名单 → 403
 *   其他 → 放行
 *@param c gin 上下文
 */
func Blacklist(c *gin.Context) {
	v, ok := c.Get(CtxKeyUserInfo)
	if !ok {
		utils.Abort(c, http.StatusUnauthorized, "need login")
		return
	}
	info, ok := v.(*user.Info)
	if !ok || info == nil {
		utils.Abort(c, http.StatusUnauthorized, "need login")
		return
	}
	if config.InBlacklist(info.UID) {
		utils.Abort(c, http.StatusForbidden, "blocked")
		return
	}
	c.Next()
}
