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

type Goods struct {

}

func (g *Goods) Add(ctx *gin.Context)  {
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

func (g *Goods) List(ctx *gin.Context)  {
	// 获取商品列表

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

func (g *Goods) Get(ctx *gin.Context)  {
	// 获取某个商品

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

func (g *Goods) Update(ctx *gin.Context)  {
	// 修改某个商品

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

func (g *Goods) Delete(ctx *gin.Context)  {
	// 删除某个商品

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