package config

var env = Env{
	Debug: true,
	
	ServerPort: "8000",
	
	MysqlUser:   "root",
	MysqlPasswd: "123",
	MysqlAddr:   "127.0.0.1:3306",
	MysqlDBName: "shop",
	
	MaxIdleConns: 50,
	MaxOpenConns: 100,
	
	MongodbHost: "10.32.5.87:27017",
	MongodbName: "pms",
	
	RedisIp:       "127.0.0.1",
	RedisPort:     "6379",
	RedisPassword: "",
	RedisDb:       11,
	
	EsServers:   "http://127.0.0.1:9200/",
	EsIndexName: "subject",
	
	AppLogDevice:  "file", //mongodb or file
	AccessLogPath: "storage/logs/access.log",
	GrpcLogPath:   "storage/logs/grpc.log",
	
	
	GrpcClient: "127.0.0.1:50051",
	
	AppSecret: "abc",
}
