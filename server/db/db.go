package db

import (
	"fmt"
	"myapp/global"

	"io/ioutil"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	m := global.GB_CONFIG.Mysql
	server := global.GB_CONFIG

	config, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmt.Print(err)
	}

	//yaml文件内容影射到结构体中
	yaml.Unmarshal(config, &server)
	m = server.Mysql

	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		fmt.Println("error in gorm.Open", err)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}