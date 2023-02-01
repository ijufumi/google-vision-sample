package configs

import (
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

type Config struct {
	Migration struct {
		Path string `env:"MIGRATION_PATH"`
	}
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
