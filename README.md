# gin-demo

## Features
- [x] 路由中间件
    - [x] api签名
    - [x] 日志记录
    - [x] Jwt
- [x] 参数验证
    - [x] 模型绑定和验证
    - [x] 自定义验证器
- [x] 存储
    - [x] MySQL
    - [x] Redis
    - [x] MongoDB
    - [x] ES
    - [x] RabbitMq
- [x] gRPC
    - [x] 客户端
    - [x] 日志记录
- [x] cron定时任务
- [x] swagger

## Proto
```
protoc -I . --go_out=plugins=grpc:. ./helloworld.proto
```

## grpc-server
```
go run grpc-server/main.go
```


## Skip sign
```
skip_debug=18120080
```

## govendor 
初始化vendor目录
```
govendor init
```

将GOPATH中本工程使用到的依赖包自动移动到vendor目录中
```
govendor add +external
```

## swagger
http://127.0.0.1:8000/swagger/index.html
```
swag init
```

## docker
```
$ cd docker
$ docker-compose up -d
```

## docker ip
```
$ docker-machine ip
```

