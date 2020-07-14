package config

var env = Env{
	Debug: true,

	ServerPort: "8000",

	MysqlUser:   "root",
	MysqlPasswd: "123456",
	MysqlAddr:   "192.168.99.100:3306",
	MysqlDBName: "shop",

	MaxIdleConns: 50,
	MaxOpenConns: 100,

	MongodbHost: "192.168.99.100:27017",
	MongodbName: "pms",

	RedisIp:       "192.168.99.100",
	RedisPort:     "6379",
	RedisPassword: "",
	RedisDb:       11,

	EsServers:   "http://192.168.99.100:9200/",
	EsIndexName: "subject",

	AppLogDevice:  "file", //mongodb or file
	AccessLogPath: "storage/logs/access.log",
	GrpcLogPath:   "storage/logs/grpc.log",

	GrpcClient: "127.0.0.1:50051",

	AppSecret: "abc",

	MqServers:  "192.168.99.100:5672",
	MQName:     "admin",
	MQPassword: "admin",
}
