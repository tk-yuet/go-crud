package config

import "fmt"

type MySqlConfig struct {
	DbDriver string
	*Config
}

func NewMySqlConfig(user string, pw string, dbname string) *MySqlConfig {
	defaultConf := NewConfig(user, pw, dbname)
	dbConf := &MySqlConfig{DbDriver: "mysql", Config: defaultConf}
	return dbConf
}

func (c *MySqlConfig) ConnectionString() string {
	connectStr := fmt.Sprintf("%s:%s@/%s", c.Config.User, c.Config.Password, c.Config.DbName)
	return connectStr
}
