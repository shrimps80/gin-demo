package main

import (
	routeRegister "gin-demo/routes"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	r := gin.Default()
	
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "找不到该路由",
		})
	})
	
	routeRegister.RegisterApiRouter(r)
	
	return r
}
