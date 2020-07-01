package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterApiRouter(router *gin.Engine) {
	apiRouter := router.Group("api")
	{
		apiRouter.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}
