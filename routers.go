package main

import (
	routeRegister "gin-demo/routes"
	"github.com/gin-gonic/gin"
	"gin-demo/modules/response"
	"gin-demo/defs"
)

func initRouter() *gin.Engine {
	r := gin.Default()
	
	r.NoRoute(func(c *gin.Context) {
		response.ReturnErrorJson(c, defs.ErrorNotFound)
	})
	
	r.NoMethod(func(c *gin.Context) {
		response.ReturnErrorJson(c, defs.ErrorNotMethod)
	})
	
	routeRegister.RegisterApiRouter(r)
	
	return r
}
