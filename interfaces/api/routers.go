/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/8 下午2:21
 * Description:
 **/

package api

import (
	"github.com/weitrue/Seckill/application/api"
	"github.com/weitrue/Seckill/infrastructure/services/local/circuitbreaker"
	"github.com/weitrue/Seckill/infrastructure/services/local/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func initRouters(g *gin.Engine) {
	g.POST("/login", api.User{}.Login)
	// api接口路由注册到gin中
	logrus.Info("Init api routers")
	activityCB := circuitbreaker.NewCircuitBreaker(
		circuitbreaker.WithDuration(100),
		circuitbreaker.WithTotalLimit(20000), //  限制为 100 毫秒内最多处理 20000 次请求
		circuitbreaker.WithLatencyLimit(100), // 最大延迟 100 毫秒
		circuitbreaker.WithFailsLimit(5),     // 最大失败率 5%
	)
	activityCBMW := middlewares.NewCircuitBreakMiddleware(activityCB)
	// 熔断器 不使用身份校验中间件
	activity := g.Group("/activity").Use(activityCBMW, middlewares.NewAuthMiddleware(false))

	// 活动
	activityApp := api.Activity{}
	activity.GET("/list", activityApp.List)
	activity.GET("/info", activityApp.Info)

	sub := g.Group("/activity/subscribe").Use(middlewares.NewAuthMiddleware(true))
	sub.POST("/", activityApp.Subscribe)

	shopCB := circuitbreaker.NewCircuitBreaker(
		circuitbreaker.WithDuration(100),
		circuitbreaker.WithTotalLimit(1000),  //  限制为 100 毫秒内最多处理 1000 次请求
		circuitbreaker.WithLatencyLimit(200), // 最大延迟 200 毫秒
		circuitbreaker.WithFailsLimit(5),     // 最大失败率 5%
	)
	mdws := []gin.HandlerFunc{
		middlewares.NewCircuitBreakMiddleware(shopCB),
		middlewares.NewAuthMiddleware(true),
		middlewares.Blacklist,
	}
	// 商品
	shop := g.Group("/shop").Use(mdws...)
	shopApp := api.Shop{}
	shop.POST("/cart/add", shopApp.AddCart)
}
