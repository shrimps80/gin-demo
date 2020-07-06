package grpc

import (
	"fmt"
	"time"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gin-demo/config"
)

func CreateServiceConn(c *gin.Context) *grpc.ClientConn {
	return createGrpcConn(config.GetEnv().GrpcClient, c)
}

func createGrpcConn(serviceAddress string, c *gin.Context) *grpc.ClientConn {
	
	var conn *grpc.ClientConn
	var err error
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()
	
	conn, err = grpc.DialContext(
		ctx,
		serviceAddress,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	
	if err != nil {
		fmt.Println(serviceAddress, "grpc conn err:", err)
	}
	return conn
}
