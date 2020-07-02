package routes

import (
	"github.com/gin-gonic/gin"
	"gin-demo/modules/response"
	"gin-demo/controllers"
	"gin-demo/middleware/logger"
)

func RegisterApiRouter(router *gin.Engine) {
	apiRouter := router.Group("api")
	
	apiRouter.Use(logger.SetUp())
	
	apiRouter.GET("/ping", func(c *gin.Context) {
		response.ReturnHttpJsonData(c, "pong")
	})
	
	apiRouter.GET("/user/:id", controllers.GetUserInfo)
}
