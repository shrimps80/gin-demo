package logger

import (
	"os"
	"log"
	"fmt"
	"bytes"
	"github.com/gin-gonic/gin"
	"gin-demo/modules/tools"
	"gin-demo/modules/response"
	"encoding/json"
	"gin-demo/config"
	"gin-demo/modules/database/mongo"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

var accessChannel = make(chan string, 100)

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func SetUp() gin.HandlerFunc {
	
	go handleAccessChannel()
	
	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		
		// 开始时间
		startTime := tools.GetCurrentMilliUnix()
		
		// 处理请求
		c.Next()
		
		responseBody := bodyLogWriter.body.String()
		
		var responseCode int
		var responseMsg string
		var responseData interface{}
		
		if responseBody != "" {
			res := response.Response{}
			err := json.Unmarshal([]byte(responseBody), &res)
			if err == nil {
				responseCode = res.Code
				responseMsg = res.Message
				responseData = res.Data
			}
		}
		
		// 结束时间
		endTime := tools.GetCurrentMilliUnix()
		
		if c.Request.Method == "POST" {
			_ = c.Request.ParseForm()
		}
		
		// 日志格式
		accessLogMap := make(map[string]interface{})
		
		accessLogMap["request_time"] = startTime
		accessLogMap["request_method"] = c.Request.Method
		accessLogMap["request_uri"] = c.Request.RequestURI
		accessLogMap["request_proto"] = c.Request.Proto
		accessLogMap["request_ua"] = c.Request.UserAgent()
		accessLogMap["request_referer"] = c.Request.Referer()
		accessLogMap["request_post_data"] = c.Request.PostForm.Encode()
		accessLogMap["request_client_ip"] = c.ClientIP()
		
		accessLogMap["response_time"] = endTime
		accessLogMap["response_code"] = responseCode
		accessLogMap["response_msg"] = responseMsg
		accessLogMap["response_data"] = responseData
		
		accessLogMap["cost_time"] = fmt.Sprintf("%vms", endTime-startTime)
		
		byteStr, _ := json.Marshal(accessLogMap)
		accessLogJson := string(byteStr)
		accessChannel <- accessLogJson
	}
}

func handleAccessChannel() {
	if config.GetEnv().AppLogDevice == "file" {
		if f, err := os.OpenFile(config.GetEnv().AccessLogPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err != nil {
			log.Println(err)
		} else {
			for accessLog := range accessChannel {
				_, _ = f.WriteString(accessLog + "\n")
			}
		}
	}
	
	if config.GetEnv().AppLogDevice == "mongodb" {
		name := fmt.Sprintf("access_log_%s", tools.GetToday())
		for accessLog := range accessChannel {
			logMap := make(map[string]interface{})
			json.Unmarshal([]byte(accessLog), &logMap)
			mongo.Client.InsertOne(name, logMap)
		}
	}
	return
}
