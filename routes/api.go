package routes

import (
	"github.com/gin-gonic/gin"
	"gin-demo/modules/response"
	"gin-demo/controllers"
)

func RegisterApiRouter(router *gin.Engine) {
	apiRouter := router.Group("api")
	{
		apiRouter.GET("/ping", func(c *gin.Context) {
			response.ReturnHttpJsonData(c, "pong")
		})
		
		apiRouter.GET("/user/:id", controllers.GetUserInfo)
	}
}
