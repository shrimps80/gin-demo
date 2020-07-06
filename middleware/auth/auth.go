package auth

import (
	"github.com/gin-gonic/gin"
	"strings"
	"gin-demo/services"
	"gin-demo/modules/tools"
	"gin-demo/modules/database/redis"
	"github.com/astaxie/beego/logs"
	"gin-demo/modules/response"
	"gin-demo/defs"
)

func SetUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		token = strings.Replace(token, "Bearer ", "", -1)
		if token == "" {
			response.ReturnErrorJson(c, defs.LoginAuthFail)
			c.Abort()
		}
		
		user, err := services.ValidateToken(token)
		if err != nil {
			response.ReturnErrorJson(c, defs.LoginAuthFail)
			c.Abort()
		}
		
		uToken, err := redis.Client.Get(tools.UserTokenKey(user.Id))
		if err != nil {
			response.ReturnErrorJson(c, defs.LoginAuthFail)
			c.Abort()
		}
		
		if len(uToken) == 0 || uToken != token {
			logs.Error("validate token, lose !!")
			response.ReturnErrorJson(c, defs.LoginAuthFail)
			c.Abort()
		}
		
		c.Set("user_id", user.Id)
	}
}
