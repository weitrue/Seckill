/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午12:54
 * Description: user HTTP handler
 *   当前接口是空壳,阶段 3 后对接 application.user.Service
 **/

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/weitrue/Seckill/pkg/utils"
)

// UserHandler 用户相关 HTTP handler
type UserHandler struct {
	// TODO: 阶段 3 接入 application.user.Service
}

/*NewUserHandler
 *@Description: 构造 UserHandler
 *@return *UserHandler
 */
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

/*Login
 *@Description: POST /login
 *@receiver h
 *@param c gin 上下文
 */
func (h *UserHandler) Login(c *gin.Context) {
	logrus.Info("user login")
	c.JSON(http.StatusOK, &utils.Response{Code: 0, Msg: "ok"})
}
