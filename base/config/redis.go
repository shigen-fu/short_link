package config

type defaultRedisConfig struct {
	Address  string `ini:"address"`
	Password string `ini:"password"`
	Db       int    `ini:"db"`
}

type redisConfig interface {
	GetAddress() string
	GetPassword() string
	getDb() int
}

func (r defaultRedisConfig) GetAddress() string {
	return r.Address
}

func (r defaultRedisConfig) GetPassword() string {
	return r.Password
}

func (r defaultRedisConfig) GetDb() int {
	return r.Db
}
