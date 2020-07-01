package controllers

import (
	"gin-demo/models/Users"
	
	"github.com/gin-gonic/gin"
	
	"gin-demo/modules/response"
	"gin-demo/modules/tools"
)

func GetUserInfo(c *gin.Context) {
	id := tools.String2Int64(c.Param("id"))
	row, _ := Users.GetOneById(id)
	response.ReturnHttpJsonData(c, row)
}
