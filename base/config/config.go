package config

import (
	"github.com/go-ini/ini"
	"sync"
)

var (
	mysqlConf defaultMysqlConfig
	redisConf defaultRedisConfig
	lock      sync.Mutex
)

func Init(filepath string) {
	var (
		err error
		cfg *ini.File
	)

	lock.Lock()
	defer lock.Unlock()
	if cfg, err = ini.Load(filepath); err != nil {
		panic(err)
	}
	if err = cfg.Section("mysql").MapTo(&mysqlConf); err != nil {
		panic(err)
	}
	if err = cfg.Section("redis").MapTo(&redisConf); err != nil {
		panic(err)
	}
}

func GetMysqlConfig() defaultMysqlConfig {
	return mysqlConf
}

func GetRedisConfig() defaultRedisConfig {
	return redisConf
}
