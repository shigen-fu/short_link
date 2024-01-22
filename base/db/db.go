package db

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"sync"
)

var (
	err  error
	db   *gorm.DB
	rdb  *redis.Client
	lock sync.Mutex
)

func Init() {
	lock.Lock()
	defer lock.Unlock()
	InitMysql()
	InitRedis()
}

func GetMysqlDb() *gorm.DB {
	return db
}

func GetRedisDb() *redis.Client {
	return rdb
}
