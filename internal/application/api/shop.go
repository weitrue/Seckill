/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午12:49
 * Description:
 **/

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/weitrue/Seckill/internal/domain/shop"
	"github.com/weitrue/Seckill/internal/domain/stock/memstock"
	"github.com/weitrue/Seckill/internal/domain/user"
	"github.com/weitrue/Seckill/pkg/utils"
)

type Shop struct { // 商品信息
}

func (s *Shop) AddCart(ctx *gin.Context) {
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

	// 内存预扣: 快速过滤明显超额的请求,减轻下游 Redis 压力
	// 允许过度预扣(多实例独立 + 懒加载偏差),真实源以 Redis Lua 为准
	st, err := memstock.NewMemoryStock(params.ActivityID, params.GoodsID)
	if err != nil {
		response.Msg = "internal server error"
		status = http.StatusInternalServerError
		ctx.JSON(status, response)
		return
	}
	if s, subErr := st.Sub(userInfo.UID); subErr != nil {
		// 冷启动从 redis 拉库存失败,属于内部错误
		logrus.Errorf("logType:ShopAddCart, err:%s, step:memstock init, activityID:%s, goodsID:%s, uid:%s",
			subErr.Error(), params.ActivityID, params.GoodsID, userInfo.UID)
		response.Msg = "internal server error"
		status = http.StatusInternalServerError
		ctx.JSON(status, response)
		return
	} else if s < 0 {
		response.Code = shop.ErrNoStock
		response.Msg = "no stock"
		ctx.JSON(http.StatusOK, response)
		return
	}

	// Hijack方法
	conn, w, err := ctx.Writer.Hijack()
	if err != nil {
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
