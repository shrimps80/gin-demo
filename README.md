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
- [x] gRPC
    - [x] 客户端
    - [x] 日志记录

## Proto
```apple js
protoc -I . --go_out=plugins=grpc:. ./helloworld.proto
```