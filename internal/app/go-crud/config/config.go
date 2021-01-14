package config

type Config struct {
	User     string
	Password string
	DbName   string
}

func NewConfig(user string, pw string, dbName string) *Config {
	return &Config{
		User:     user,
		Password: pw,
		DbName:   dbName,
	}
}
