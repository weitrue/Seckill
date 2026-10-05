/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午2:21
 * Description: HTTP 路由注册入口
 *   依赖通过 Dependencies 结构显式注入,方便单测和替换实现
 **/

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/weitrue/Seckill/interfaces/http/api/middleware"
	shopapp "github.com/weitrue/Seckill/internal/application/shop"
	"github.com/weitrue/Seckill/internal/infrastructure/services/local/circuitbreaker"
)

// Dependencies HTTP 路由依赖
//
//	在 interfaces/api/api.Run 中装配后传入
type Dependencies struct {
	// ShopService 下单应用服务(application.shop.Service)
	ShopService *shopapp.Service
}

/*InitRouters
 *@Description: 注册 HTTP 路由
 *@param g gin engine
 *@param deps 路由依赖(application 层服务等)
 */
func InitRouters(g *gin.Engine, deps Dependencies) {
	logrus.Info("Init api routers")

	// 用户
	userH := NewUserHandler()
	g.POST("/login", userH.Login)

	// 活动
	activityCB := circuitbreaker.NewCircuitBreaker(
		circuitbreaker.WithDuration(100),
		circuitbreaker.WithTotalLimit(20000), // 100 毫秒最多 20000 次请求
		circuitbreaker.WithLatencyLimit(100), // 最大延迟 100 毫秒
		circuitbreaker.WithFailsLimit(5),     // 最大失败率 5%
	)
	activityCBMW := middleware.CircuitBreak(activityCB)
	activityH := NewActivityHandler()
	// 熔断 + 鉴权非强制(允许匿名访问活动列表 / 详情)
	activityGroup := g.Group("/activity").Use(activityCBMW, middleware.Auth(false))
	activityGroup.GET("/list", activityH.List)
	activityGroup.GET("/info", activityH.Info)
	// 订阅活动需要登录
	subGroup := g.Group("/activity/subscribe").Use(middleware.Auth(true))
	subGroup.POST("/", activityH.Subscribe)

	// 商品
	shopCB := circuitbreaker.NewCircuitBreaker(
		circuitbreaker.WithDuration(100),
		circuitbreaker.WithTotalLimit(1000),  // 100 毫秒最多 1000 次请求
		circuitbreaker.WithLatencyLimit(200), // 最大延迟 200 毫秒
		circuitbreaker.WithFailsLimit(5),     // 最大失败率 5%
	)
	shopMdws := []gin.HandlerFunc{
		middleware.CircuitBreak(shopCB),
		middleware.Auth(true),
		middleware.Blacklist,
	}
	shopH := NewShopHandler(deps.ShopService)
	shopGroup := g.Group("/shop").Use(shopMdws...)
	shopGroup.POST("/cart/add", shopH.AddCart)
}
