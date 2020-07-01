package mysql

import (
	"fmt"
	"flag"
	"time"
	"github.com/jinzhu/gorm"
	"gin-demo/config"
)

var SqlDB *gorm.DB

func init() {
	dbConStr := fmt.Sprintf("%s:%s@tcp(%s)/%s%s?charset=utf8&parseTime=True&loc=Local",
		config.GetEnv().MysqlUser,
		config.GetEnv().MysqlPasswd,
		config.GetEnv().MysqlAddr,
		config.GetEnv().MysqlDBName, getConnectDbName())
	
	db, err := gorm.Open("mysql", dbConStr)
	
	SqlDB = db
	if err != nil {
		panic(fmt.Errorf("Connect %s error: %s", dbConStr, err))
	}
	
	// 设置数据库最大连接 减少timewait 正式环境调大
	SqlDB.DB().SetMaxIdleConns(config.GetEnv().MaxIdleConns) // 连接池连接数 = mysql最大连接数/2
	SqlDB.DB().SetMaxOpenConns(config.GetEnv().MaxOpenConns) // 最大打开连接 = mysql最大连接数
	
	// 设置链接重置时间
	SqlDB.DB().SetConnMaxLifetime(80 * time.Second)
	SqlDB.LogMode(true)
	SqlDB.SingularTable(true)
}

func getConnectDbName() string {
	dbName := flag.String("db", "", "")
	flag.Parse()
	
	return *dbName
}
