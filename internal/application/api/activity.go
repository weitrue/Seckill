/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午12:52
 * Description:
 **/

package api

import (
	"net/http"

	"github.com/weitrue/Seckill/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Activity struct { // 秒杀活动
}

func (e *Activity) List(ctx *gin.Context) {
	// 获取所有正在进行或者即将进行的活动

	// response
	response := &utils.Response{
		Code: 0,
		Data: nil,
		Msg:  "ok",
	}
	status := http.StatusOK
	logrus.Info("Activity list")
	ctx.JSON(status, response)
}

func (e *Activity) Info(ctx *gin.Context) {
	// 产看某个商品的秒杀活动信息

	// response
	response := &utils.Response{
		Code: 0,
		Data: nil,
		Msg:  "ok",
	}
	status := http.StatusOK
	logrus.Info("Activity Info")
	ctx.JSON(status, response)
}

func (e *Activity) Subscribe(ctx *gin.Context) {
	// 订阅某商品的活动开始通知

	// response
	response := &utils.Response{
		Code: 0,
		Data: nil,
		Msg:  "ok",
	}
	status := http.StatusOK
	logrus.Info("Activity Subscribe")
	ctx.JSON(status, response)
}
