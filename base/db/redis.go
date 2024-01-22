package db

import (
	"github.com/redis/go-redis/v9"
	"short-link/base/config"
)

func InitRedis() {
	redisConf := config.GetRedisConfig()
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisConf.GetAddress(),
		Password: redisConf.GetPassword(),
		DB:       redisConf.GetDb(),
	})
}
