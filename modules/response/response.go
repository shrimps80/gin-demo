package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-demo/defs"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func ReturnErrorJson(c *gin.Context, def defs.ErrResponse) {
	responseTo(c, Response{
		Code:    def.Error.ErrorCode,
		Message: def.Error.Error,
		Data: struct {
		}{},
	}, def.HttpSC)
}

func ReturnHttpJsonData(c *gin.Context, val interface{}) {
	if val == nil {
		val = struct{}{}
	}
	responseTo(c, Response{
		Code:    0,
		Message: "ok",
		Data:    val,
	}, http.StatusOK)
}

func responseTo(c *gin.Context, data Response, statusCode int) {
	c.JSON(statusCode, data)
	return
}
