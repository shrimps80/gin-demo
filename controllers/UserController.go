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
	"gin-demo/modules/database/elasticsearch"
	"gin-demo/filter/User"
	"gin-demo/defs"
	"github.com/gookit/validate"
)

func GetUserInfo(c *gin.Context) {
	// validate
	u := User.UserForm{}
	
	s, e := tools.Bind(&u, c)
	if e != nil {
		response.ReturnErrorJson(c, defs.ErrorLostParams)
		return
	}
	
	v := validate.Struct(s)
	if !v.Validate() {
		fmt.Println(v.Errors) // 所有的错误消息
		//fmt.Println(v.Errors.One()) // 返回随机一条错误消息
		//fmt.Println(v.Errors.Field("Name")) // 返回该字段的错误消息
		response.ReturnErrorJson(c, defs.ValidateErr(v.Errors.One()))
		return
	}
	
	id := tools.String2Int64(c.Param("id"))
	row, err := Users.GetOneById(id)
	if err != nil {
		log.Error(err.Error())
	}
	if row != nil {
		redisKey := fmt.Sprintf("user:%d", row.Id)
		r := redis.Client
		r.Set(redisKey, row.Name, 60*time.Second)
		
		//es
		e := elasticsearch.Client
		mapping := `{
			"settings":{
				"number_of_shards":1,
				"number_of_replicas":0
			},
			"mappings":{
				"properties":{
					"id":{
						"type":"long"
					},
					"name":{
						"type":"text"
					}
				}
			}
		}`
		idStr := tools.Int64ToString(row.Id)
		e.IndexExists(mapping)
		e.SetIndex(idStr, row)
		ecData := e.GetIndex(idStr)
		e.DelIndex(idStr)
		response.ReturnHttpJsonData(c, ecData)
		
		return
	}
	response.ReturnHttpJsonData(c, nil)
}
