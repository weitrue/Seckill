/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午4:49
 * Description: 专题管理 HTTP handler(当前为空壳)
 *   等后续补 application.topic.Service 时在此做协议适配
 **/

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/weitrue/Seckill/pkg/utils"
)

// TopicHandler 专题管理相关 HTTP handler
type TopicHandler struct {
	// TODO: 接入 application.topic.Service
}

/*NewTopicHandler
 *@Description: 构造 TopicHandler
 *@return *TopicHandler
 */
func NewTopicHandler() *TopicHandler {
	return &TopicHandler{}
}

func (h *TopicHandler) Add(c *gin.Context) {
	logrus.Info("Topic Add")
	c.JSON(http.StatusOK, &utils.Response{Code: 0, Msg: "ok"})
}

func (h *TopicHandler) List(c *gin.Context) {
	logrus.Info("Topic List")
	c.JSON(http.StatusOK, &utils.Response{Code: 0, Msg: "ok"})
}

func (h *TopicHandler) Get(c *gin.Context) {
	logrus.Info("Topic Get")
	c.JSON(http.StatusOK, &utils.Response{Code: 0, Msg: "ok"})
}

func (h *TopicHandler) Status(c *gin.Context) {
	logrus.Info("Topic Status")
	c.JSON(http.StatusOK, &utils.Response{Code: 0, Msg: "ok"})
}

func (h *TopicHandler) Update(c *gin.Context) {
	logrus.Info("Topic Update")
	c.JSON(http.StatusOK, &utils.Response{Code: 0, Msg: "ok"})
}

func (h *TopicHandler) Delete(c *gin.Context) {
	logrus.Info("Topic Delete")
	c.JSON(http.StatusOK, &utils.Response{Code: 0, Msg: "ok"})
}
