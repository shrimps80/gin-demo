package main

import (
	routeRegister "gin-demo/routes"
	"github.com/gin-gonic/gin"
	"gin-demo/modules/response"
	"gin-demo/defs"
	_ "gin-demo/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func initRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	r.NoRoute(func(c *gin.Context) {
		response.ReturnErrorJson(c, defs.ErrorNotFound)
	})
	
	r.NoMethod(func(c *gin.Context) {
		response.ReturnErrorJson(c, defs.ErrorNotMethod)
	})
	
	routeRegister.RegisterApiRouter(r)
	
	return r
}
