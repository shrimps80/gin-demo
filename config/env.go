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
	
	AccessLogDevice: "mongodb", //mongodb or file
	AccessLogPath:   "storage/logs/access.log",
}
