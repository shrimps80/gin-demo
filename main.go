package main

import (
	"runtime"
	"gin-demo/config"
	"github.com/gin-gonic/gin"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	if config.GetEnv().Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	
	r := initRouter()
	
	r.Run(":" + config.GetEnv().ServerPort)
}
