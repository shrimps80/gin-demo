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
