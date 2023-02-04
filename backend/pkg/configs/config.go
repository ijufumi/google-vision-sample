package configs

import (
	"fmt"
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

type Config struct {
	Migration struct {
		Path string `env:"MIGRATION_PATH" envDefault:"migration"`
	}
	DB struct {
		User     string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		Port     uint   `env:"DB_PORT" envDefault:"5432"`
		Name     string `env:"DB_NAME"`
		Host     string `env:"DB_HOST"`
	}
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	c := Config{}
	err = env.Parse(&c)
	if err != nil {
		fmt.Println(err)
	}

	return &c
}
