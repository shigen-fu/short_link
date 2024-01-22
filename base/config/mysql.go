package config

type defaultMysqlConfig struct {
	Address  string `ini:"address"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	DbName   string `ini:"dbName"`
}

type mysqlConfig interface {
	GetAddress() string
	GetUsername() string
	GetPassword() string
	getDbname() string
}

func (m defaultMysqlConfig) GetAddress() string {
	return m.Address
}

func (m defaultMysqlConfig) GetUsername() string {
	return m.Username
}

func (m defaultMysqlConfig) GetPassword() string {
	return m.Password
}

func (m defaultMysqlConfig) GetDbname() string {
	return m.DbName
}
