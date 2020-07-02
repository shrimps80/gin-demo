package controllers

import (
	"fmt"
	"time"
	"gin-demo/models/Users"
	
	"github.com/gin-gonic/gin"
	
	"gin-demo/modules/response"
	"gin-demo/modules/tools"
	"github.com/silenceper/log"
	"gin-demo/modules/database/redis"
)

func GetUserInfo(c *gin.Context) {
	id := tools.String2Int64(c.Param("id"))
	row, err := Users.GetOneById(id)
	if err != nil {
		log.Error(err.Error())
	}
	if row != nil {
		redisKey := fmt.Sprintf("user:%d", row.Id)
		r := redis.Client
		r.Set(redisKey, row.Name, 60*time.Second)
	}
	response.ReturnHttpJsonData(c, row)
}
