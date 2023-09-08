package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var mysqlLogger logger.Interface

func init() {
	username := "root"
	password := "root"
	host := "127.0.0.1"
	port := 3306
	DbName := "chat"
	timeOut := "10s"
	mysqlLogger = logger.Default.LogMode(logger.Info)

	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		username, password, host, port, DbName, timeOut)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "",//表名前缀
			SingularTable: true, //是否为单数表名
			//NoLowerCase:true,   //不要小写转换
		},
	})
	if err != nil {
		panic("连接数据库失败,err" + err.Error())

	}

	//连接成功
	DB = db
	DB = db.Session(&gorm.Session{
		Logger: mysqlLogger,
	})
}
