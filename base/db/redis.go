package db

import (
	"github.com/redis/go-redis/v9"
	"short-link/base/config"
	"time"
)

func InitRedis() {
	redisConf := config.GetRedisConfig()
	rdb = redis.NewClient(&redis.Options{
		Addr:           redisConf.Address,
		Password:       redisConf.Password,
		DB:             redisConf.Db,
		DialTimeout:    time.Second * 3,
		PoolSize:       10,
		MinIdleConns:   2,
		MaxActiveConns: 10,
	})
}
