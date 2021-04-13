/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/13 下午12:11
 * Description:
 **/

package middlewares

import (
	"Seckill/domain/user"
	"Seckill/infrastructure/config"
	"Seckill/infrastructure/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
