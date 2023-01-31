package configs

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	DB struct {
		User     string `env:"DB_USER`
		Password string `env:"DB_PASSWORD`
		Port     uint   `env:"DB_PORT`
		Name     string `env:"DB_NAME`
		Host     string `env:"DB_HOST`
	}
}

func New() *Config {
	_ = godotenv.Load()
	c := Config{}
	_ = env.Parse(&c)

	return &c
}
