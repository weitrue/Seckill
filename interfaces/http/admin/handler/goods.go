/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午4:49
 * Description: 商品管理 HTTP handler(当前为空壳)
 *   等后续补 application.goods.Service 时在此做协议适配
 **/

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/weitrue/Seckill/pkg/utils"
)

// GoodsHandler 商品管理 HTTP handler
type GoodsHandler struct {
	// TODO: 接入 application.goods.Service
}

/*NewGoodsHandler
 *@Description: 构造 GoodsHandler
 *@return *GoodsHandler
 */
func NewGoodsHandler() *GoodsHandler {
	return &GoodsHandler{}
}

func (h *GoodsHandler) Add(c *gin.Context) {
	logrus.Info("Goods Add")
	c.JSON(http.StatusOK, &utils.Response{Code: 0, Msg: "ok"})
}

func (h *GoodsHandler) List(c *gin.Context) {
	logrus.Info("Goods List")
	c.JSON(http.StatusOK, &utils.Response{Code: 0, Msg: "ok"})
}

func (h *GoodsHandler) Get(c *gin.Context) {
	logrus.Info("Goods Get")
	c.JSON(http.StatusOK, &utils.Response{Code: 0, Msg: "ok"})
}

func (h *GoodsHandler) Update(c *gin.Context) {
	logrus.Info("Goods Update")
	c.JSON(http.StatusOK, &utils.Response{Code: 0, Msg: "ok"})
}

func (h *GoodsHandler) Delete(c *gin.Context) {
	logrus.Info("Goods Delete")
	c.JSON(http.StatusOK, &utils.Response{Code: 0, Msg: "ok"})
}
