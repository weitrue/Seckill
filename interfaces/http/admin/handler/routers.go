/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午5:37
 * Description: admin HTTP 路由注册入口
 *   现阶段只有 Activity 对接到 AdminService,Goods / Topic 占位
 **/

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	activityapp "github.com/weitrue/Seckill/internal/application/activity"
)

// Dependencies admin HTTP 路由依赖
type Dependencies struct {
	// ActivityAdminService 活动后台应用服务
	ActivityAdminService *activityapp.AdminService
	// TODO: 后续加 GoodsService / TopicService
}

/*InitRouters
 *@Description: 注册 admin HTTP 路由
 *@param g gin engine
 *@param deps 路由依赖
 */
func InitRouters(g *gin.Engine, deps Dependencies) {
	logrus.Info("Init admin routers")

	adminGroup := g.Group("/admin")

	// 专题(空壳,待业务)
	topicH := NewTopicHandler()
	adminGroup.GET("/topic/", topicH.List)
	adminGroup.GET("/topic/:id", topicH.Get)
	adminGroup.POST("/topic/", topicH.Add)
	adminGroup.PUT("/topic/:id", topicH.Update)
	adminGroup.PUT("/topic/:id/:status", topicH.Status)
	adminGroup.DELETE("/topic/:id", topicH.Delete)

	// 活动(对接 AdminService)
	activityH := NewAdminActivityHandler(deps.ActivityAdminService)
	adminGroup.GET("/activity/", activityH.List)
	adminGroup.GET("/activity/:id", activityH.Get)
	adminGroup.POST("/activity/", activityH.Add)
	adminGroup.PUT("/activity/:id", activityH.Update)
	adminGroup.PUT("/activity/:id/:status", activityH.Status)
	adminGroup.DELETE("/activity/:id", activityH.Delete)

	// 商品(空壳,待业务)
	goodsH := NewGoodsHandler()
	adminGroup.GET("/goods/", goodsH.List)
	adminGroup.GET("/goods/:id", goodsH.Get)
	adminGroup.POST("/goods/", goodsH.Add)
	adminGroup.POST("/goods/:id", goodsH.Update)
	adminGroup.DELETE("/goods/:id", goodsH.Delete)
}
