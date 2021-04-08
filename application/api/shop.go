/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午12:49
 * Description:
 **/

package api

import (
	"Seckill/infrastructure/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Shop struct {
	// 商品信息
}

func (s *Shop)AddCart(ctx *gin.Context)  {
	// 添加商品到购物车

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