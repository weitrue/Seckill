/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午4:49
 * Description:
 **/

package admin

import (
	"net/http"

	"github.com/weitrue/Seckill/infrastructure/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Activity struct {

}


func (a *Activity) Add(ctx *gin.Context)  {
	// 创建秒杀活动

	// response
	response := &utils.Response{
		Code: 0,
		Data: nil,
		Msg:  "ok",
	}
	status := http.StatusOK
	logrus.Info("Topic Add")
	ctx.JSON(status, response)
}

func (a *Activity) List(ctx *gin.Context)  {
	// 获取秒杀活动列表

	// response
	response := &utils.Response{
		Code: 0,
		Data: nil,
		Msg:  "ok",
	}
	status := http.StatusOK
	logrus.Info("Topic List")
	ctx.JSON(status, response)
}

func (a *Activity) Get(ctx *gin.Context)  {
	// 获取某个秒杀活动

	// response
	response := &utils.Response{
		Code: 0,
		Data: nil,
		Msg:  "ok",
	}
	status := http.StatusOK
	logrus.Info("Topic Get")
	ctx.JSON(status, response)
}

func (a *Activity) Status(ctx *gin.Context)  {
	// 获取某个秒杀活动状态：上线/下线

	// response
	response := &utils.Response{
		Code: 0,
		Data: nil,
		Msg:  "ok",
	}
	status := http.StatusOK
	logrus.Info("Topic Status")
	ctx.JSON(status, response)
}

func (a *Activity) Update(ctx *gin.Context)  {
	// 修改某个秒杀活动

	// response
	response := &utils.Response{
		Code: 0,
		Data: nil,
		Msg:  "ok",
	}
	status := http.StatusOK
	logrus.Info("Topic Update")
	ctx.JSON(status, response)
}

func (a *Activity) Delete(ctx *gin.Context)  {
	// 删除某个秒杀活动

	// response
	response := &utils.Response{
		Code: 0,
		Data: nil,
		Msg:  "ok",
	}
	status := http.StatusOK
	logrus.Info("Topic Delete")
	ctx.JSON(status, response)
}

