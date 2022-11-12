/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午4:49
 * Description:
 **/

package admin

import (
	"net/http"

	"github.com/weitrue/Seckill/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Topic struct{}

func (t *Topic) Add(ctx *gin.Context) {
	// 创建专题

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

func (t *Topic) List(ctx *gin.Context) {
	// 获取专题列表

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

func (t *Topic) Get(ctx *gin.Context) {
	// 获取某个专题

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

func (t *Topic) Status(ctx *gin.Context) {
	// 获取某个专题状态：上线/下线

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

func (t *Topic) Update(ctx *gin.Context) {
	// 修改某个专题

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

func (t *Topic) Delete(ctx *gin.Context) {
	// 删除某个专题

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
