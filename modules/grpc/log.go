package grpc

import (
	"os"
	"log"
	"fmt"
	"context"
	"google.golang.org/grpc"
	"encoding/json"
	"gin-demo/modules/tools"
	"gin-demo/config"
	"gin-demo/modules/database/mongo"
)

var grpcChannel = make(chan string, 100)

func ClientInterceptor() grpc.UnaryClientInterceptor {
	
	go handleGrpcChannel()
	
	return func(ctx context.Context, method string,
		req, reply interface{}, cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		
		// 开始时间
		startTime := tools.GetCurrentMilliUnix()
		
		err := invoker(ctx, method, req, reply, cc, opts...)
		
		// 结束时间
		endTime := tools.GetCurrentMilliUnix()
		
		// 日志格式
		grpcLogMap := make(map[string]interface{})
		
		grpcLogMap["request_time"] = startTime
		grpcLogMap["request_data"] = req
		grpcLogMap["request_method"] = method
		
		grpcLogMap["response_data"] = reply
		grpcLogMap["response_error"] = err
		
		grpcLogMap["cost_time"] = fmt.Sprintf("%vms", endTime-startTime)
		
		byteStr, _ := json.Marshal(grpcLogMap)
		grpcLogJson := string(byteStr)
		grpcChannel <- grpcLogJson
		
		return err
	}
}

func handleGrpcChannel() {
	if config.GetEnv().AppLogDevice == "file" {
		if f, err := os.OpenFile(config.GetEnv().GrpcLogPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err != nil {
			log.Println(err)
		} else {
			for accessLog := range grpcChannel {
				_, _ = f.WriteString(accessLog + "\n")
			}
		}
	}
	if config.GetEnv().AppLogDevice == "mongodb" {
		name := fmt.Sprintf("grpc_log_%s", tools.GetToday())
		for accessLog := range grpcChannel {
			logMap := make(map[string]interface{})
			json.Unmarshal([]byte(accessLog), &logMap)
			mongo.Client.InsertOne(name, logMap)
		}
	}
	return
}
