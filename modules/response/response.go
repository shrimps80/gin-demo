package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-demo/defs"
)

func ReturnErrorJson(c *gin.Context, def defs.ErrResponse) {
	responseTo(c, gin.H{
		"code":    def.Error.ErrorCode,
		"message": def.Error.Error,
	}, def.HttpSC)
}

func ReturnHttpJsonData(c *gin.Context, val interface{}) {
	if val == nil {
		val = struct{}{}
	}
	responseTo(c, gin.H{
		"code": 0,
		"data": val,
	}, http.StatusOK)
}

func responseTo(c *gin.Context, data gin.H, statusCode int) {
	c.JSON(statusCode, data)
}
