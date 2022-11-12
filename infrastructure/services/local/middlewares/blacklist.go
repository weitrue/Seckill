/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午12:11
 * Description:
 **/

package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/weitrue/Seckill/domain/user"
	"github.com/weitrue/Seckill/infrastructure/config"
	"github.com/weitrue/Seckill/pkg/utils"
)

// Blacklist 黑名单中间件
func Blacklist(ctx *gin.Context) {
	data, _ := ctx.Get("UserInfo")
	info, ok := data.(*user.Info)
	if !ok {
		utils.Abort(ctx, http.StatusUnauthorized, "need login")
		return
	}
	if config.InBlacklist(info.UID) {
		utils.Abort(ctx, http.StatusForbidden, "blocked")
		return
	}
	ctx.Next()
}
