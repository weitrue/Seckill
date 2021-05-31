/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午12:49
 * Description:
 **/

package api

import (
	"net/http"

	"github.com/weitrue/Seckill/domain/shop"
	"github.com/weitrue/Seckill/domain/stock"
	"github.com/weitrue/Seckill/domain/user"
	"github.com/weitrue/Seckill/infrastructure/utils"

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

	params := struct {
		GoodsID    string `json:"goods_id"`
		ActivityID string `json:"event_id"`
	}{}
	var userInfo *user.Info
	if v, ok := ctx.Get("userInfo"); ok {
		userInfo, _ = v.(*user.Info)
	}

	err := ctx.BindJSON(&params)
	if err != nil || params.ActivityID == "" || params.GoodsID == "" || userInfo == nil {
		response.Msg = "bad request"
		status = http.StatusBadRequest
		ctx.JSON(status, response)
		return
	}
	logrus.Info(params)

	// 扣减内存缓存中的库存库存 用于初步获取资格
	st, _ := stock.NewRedisStock(params.ActivityID, params.GoodsID)
	if s, _ := st.Sub(userInfo.UID); s < 0 {
		response.Code = shop.ErrNoStock
		response.Msg = "no stock"
		ctx.JSON(http.StatusOK, response)
		return
	}

	// Hijack方法
	conn, w, err1 := ctx.Writer.Hijack()
	if err1 != nil {
		response.Msg = "bad request"
		status = http.StatusBadRequest
		ctx.JSON(status, response)
		return
	}
	logrus.Info("shop add cart")
	shopCtx := &shop.Context{
		Request: ctx.Request,
		Conn:    conn,
		Writer:  w,
		GoodsID: params.GoodsID,
		EventID: params.ActivityID,
		UID:     userInfo.UID,
	}
	shop.Handle(shopCtx)
}