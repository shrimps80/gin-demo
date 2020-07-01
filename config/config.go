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
	
	RedisIp       string
	RedisPort     string
	RedisPassword string
	RedisDb       int
}

func GetEnv() *Env {
	return &env
}