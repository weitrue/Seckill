/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2024/10/06
 * Description: shop HTTP handler
 *   职责:
 *     - 协议细节: bind HTTP body / 从 gin.Context 取鉴权用户 / ctx.JSON 响应
 *     - 调用协议无关的 application.shop.Service.AddCart
 *   不含业务逻辑
 **/

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/weitrue/Seckill/interfaces/http/api/middleware"
	shopapp "github.com/weitrue/Seckill/internal/application/shop"
	"github.com/weitrue/Seckill/internal/domain/user"
)

// ShopHandler 商品相关 HTTP handler
type ShopHandler struct {
	svc *shopapp.Service
}

/*NewShopHandler
 *@Description: 构造 ShopHandler,依赖 application.shop.Service
 *@param svc
 *@return *ShopHandler
 */
func NewShopHandler(svc *shopapp.Service) *ShopHandler {
	return &ShopHandler{svc: svc}
}

// addCartRequest 入参 JSON 结构,仅用于协议 bind
type addCartRequest struct {
	GoodsID    string `json:"goods_id"`
	ActivityID string `json:"event_id"`
}

/*AddCart
 *@Description: POST /shop/cart/add
 *   bind body → 取鉴权用户 → 组装 Command → 调 Service → ctx.JSON
 *@receiver h
 *@param c gin 上下文
 */
func (h *ShopHandler) AddCart(c *gin.Context) {
	logType := "HttpShopAddCart"
	var req addCartRequest
	if err := c.BindJSON(&req); err != nil || req.ActivityID == "" || req.GoodsID == "" {
		logrus.Warnf("logType:%s, msg:bad request, err:%v, body:%+v", logType, err, req)
		c.JSON(http.StatusBadRequest, shopapp.AddCartResult{
			Code: http.StatusBadRequest,
			Msg:  "bad request",
		})
		return
	}

	userInfo := extractUser(c)
	if userInfo == nil {
		c.JSON(http.StatusUnauthorized, shopapp.AddCartResult{
			Code: http.StatusUnauthorized,
			Msg:  "unauthorized",
		})
		return
	}

	cmd := shopapp.AddCartCommand{
		UID:        userInfo.UID,
		ActivityID: req.ActivityID,
		GoodsID:    req.GoodsID,
	}

	// 走 Request 的 ctx,中间件挂的 trace span 后续可以从这里继续
	result, _ := h.svc.AddCart(c.Request.Context(), cmd)
	c.JSON(http.StatusOK, result)
}

// extractUser 从 gin.Context 中取鉴权中间件注入的用户信息
func extractUser(c *gin.Context) *user.Info {
	v, ok := c.Get(middleware.CtxKeyUserInfo)
	if !ok {
		return nil
	}
	u, _ := v.(*user.Info)
	return u
}
