package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"short-link/base/config"
)

func InitMysql() {
	mysqlConf := config.GetMysqlConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConf.GetUsername(), mysqlConf.GetPassword(), mysqlConf.GetAddress(), mysqlConf.GetDbname())
	fmt.Printf("mysqlConf: %#v dsn:%s\n", mysqlConf, dsn)
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		panic(err)
	}
}
