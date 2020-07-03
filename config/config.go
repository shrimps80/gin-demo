package config

type Env struct {
	Debug      bool
	ServerPort string
	
	MysqlUser   string
	MysqlPasswd string
	MysqlAddr   string
	MysqlDBName string
	
	MaxIdleConns int
	MaxOpenConns int
	
	MongodbHost string
	MongodbName string
	
	RedisIp       string
	RedisPort     string
	RedisPassword string
	RedisDb       int
	
	EsServers   string
	EsIndexName string
	
	AccessLogDevice string
	AccessLogPath   string
}

func GetEnv() *Env {
	return &env
}
