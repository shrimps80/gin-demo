package grpc

import (
	"fmt"
	"time"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func CreateServiceConn(c *gin.Context) *grpc.ClientConn {
	return createGrpcConn("127.0.0.1:50051", c)
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
