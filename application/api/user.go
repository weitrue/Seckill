/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午12:54
 * Description:
 **/

package api

import (
	"net/http"

	"github.com/weitrue/Seckill/infrastructure/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type User struct {}

func (u User) Login(ctx *gin.Context) {
	// 登录

	// response
	response := &utils.Response{
		Code: 0,
		Data: nil,
		Msg:  "ok",
	}
	status := http.StatusOK
	logrus.Info("event list")
	ctx.JSON(status, response)
}
