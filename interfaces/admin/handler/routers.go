/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午5:37
 * Description:
 **/

package handler

import (
	_ "github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/weitrue/Seckill/application/admin"
)

func InitRouters(g *gin.Engine) {
	// g.HandleFunc("/debug/pprof/", pprof.Index)
	// g.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	// g.HandleFunc("/debug/pprof/profile", pprof.Profile)
	// g.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	// g.HandleFunc("/debug/pprof/trace", pprof.Trace)

	// admin接口路由注册到gin中
	logrus.Info("Init admin routers")
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