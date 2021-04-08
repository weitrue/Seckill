/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午5:37
 * Description:
 **/

package admin

import (
	"Seckill/application/admin"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitRouters(g *gin.Engine)  {

	// admin接口路由注册到gin中
	logrus.Info("Init api routers")
	adminGroup := g.Group("/admin")
	topicApp := admin.Topic{}
	activityApp := admin.Activity{}
	goodsApp := admin.Goods{}

	adminGroup.GET("/topic/", topicApp.List)
	adminGroup.GET("/topic/:id", topicApp.Get)
	adminGroup.POST("/topic/", topicApp.Add)
	adminGroup.PUT("/topic/:id", topicApp.Update)
	adminGroup.PUT("/topic/:id/:status", topicApp.Status)
	adminGroup.DELETE("/topic/:id", topicApp.Delete)

	adminGroup.GET("/activity/", activityApp.List)
	adminGroup.GET("/activity/:id", activityApp.Get)
	adminGroup.POST("/activity/", activityApp.Add)
	adminGroup.PUT("/activity/:id", activityApp.Update)
	adminGroup.PUT("/activity/:id/:status", activityApp.Status)
	adminGroup.DELETE("/activity/:id", activityApp.Delete)

	adminGroup.GET("/goods/", goodsApp.List)
	adminGroup.GET("/goods/:id", goodsApp.Get)
	adminGroup.POST("/goods/", goodsApp.Add)
	adminGroup.POST("/goods/:id", goodsApp.Update)
	adminGroup.DELETE("/goods/:id", goodsApp.Delete)

}
