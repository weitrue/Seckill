/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午12:52
 * Description: activity HTTP handler
 *   当前接口都是空壳,阶段 2 之后对接 application.activity.Service
 **/

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/weitrue/Seckill/pkg/utils"
)

// ActivityHandler 秒杀活动相关 HTTP handler
type ActivityHandler struct {
	// TODO: 阶段 2 接入 application.activity.Service
}

/*NewActivityHandler
 *@Description: 构造 ActivityHandler
 *@return *ActivityHandler
 */
func NewActivityHandler() *ActivityHandler {
	return &ActivityHandler{}
}

/*List
 *@Description: GET /activity/list 获取所有正在进行或者即将进行的活动
 *@receiver h
 *@param c gin 上下文
 */
func (h *ActivityHandler) List(c *gin.Context) {
	logrus.Info("Activity list")
	c.JSON(http.StatusOK, &utils.Response{Code: 0, Msg: "ok"})
}

/*Info
 *@Description: GET /activity/info 查看某个商品的秒杀活动信息
 *@receiver h
 *@param c gin 上下文
 */
func (h *ActivityHandler) Info(c *gin.Context) {
	logrus.Info("Activity Info")
	c.JSON(http.StatusOK, &utils.Response{Code: 0, Msg: "ok"})
}

/*Subscribe
 *@Description: POST /activity/subscribe/ 订阅某商品的活动开始通知
 *@receiver h
 *@param c gin 上下文
 */
func (h *ActivityHandler) Subscribe(c *gin.Context) {
	logrus.Info("Activity Subscribe")
	c.JSON(http.StatusOK, &utils.Response{Code: 0, Msg: "ok"})
}
