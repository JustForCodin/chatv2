package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	MysqlDSN string `env:"MYSQL_DSN"`
}

var (
	cfg *Config
)

func GetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	c := Config{}
	if err := env.Parse(&c); err != nil {
		panic(err.Error())
	}

	cfg = &c

	return *cfg
}
