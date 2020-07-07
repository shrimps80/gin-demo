package main

import (
	"runtime"
	"gin-demo/config"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "gin-demo/modules/cron"  //定时脚本
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
