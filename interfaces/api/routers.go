/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午2:21
 * Description:
 **/

package api

import (
	"Seckill/application/api"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func initRouters(g *gin.Engine) {
	// api接口路由注册到gin中
	logrus.Info("Init api routers")

	activity := g.Group("/activity")
	// 活动
	activityApp := api.Activity{}
	activity.GET("/list", activityApp.List)
	activity.GET("/info", activityApp.Info)
	activity.GET("/subscribe", activityApp.Subscribe)

	// 商品
	shop := g.Group("/shop")
	shopApp := api.Shop{}
	shop.POST("/cart/add", shopApp.AddCart)
}


