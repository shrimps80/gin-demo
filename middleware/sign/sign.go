package sign

import (
	"github.com/gin-gonic/gin"
	"gin-demo/defs"
	"gin-demo/modules/response"
	"strconv"
	"time"
	"gin-demo/config"
	"strings"
	"gin-demo/modules/tools"
)

func SetUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		sign := c.Request.Header.Get("sign")
		if sign == "" {
			response.ReturnErrorJson(c, defs.MissingSignature)
			c.Abort()
			return
		}
		
		timestamp := c.Request.Header.Get("timestamp")
		if timestamp == "" {
			response.ReturnErrorJson(c, defs.MissingSignature)
			c.Abort()
			return
		}
		
		//判断time是否过期
		timestampInt, _ := strconv.ParseInt(timestamp, 10, 64)
		minute := time.Now().Unix() - timestampInt
		if minute > 300 { // 5分钟
			response.ReturnErrorJson(c, defs.MissingSignature)
			c.Abort()
			return
		}
		
		SignKey := config.GetEnv().AppSecret
		signArray := []string{
			timestamp,
			SignKey,
		}
		signMd5 := strings.ToUpper(tools.Md5Str(strings.Join(signArray, "&")))
		if signMd5 != sign {
			response.ReturnErrorJson(c, defs.MissingSignature)
			c.Abort()
			return
		}
	}
}
