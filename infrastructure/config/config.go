package config

import (
	"github.com/caarlos0/env/v7"
)

type EnvConfig struct {
	DBLogin string `env:"DB_LOGIN"`
	DBPass  string `env:"DB_PASS"`
	DBName  string `env:"DB_NAME"`
	DBHost 	string `env:"DB_HOST"`
	DBPort  string `env:"DB_PORT"`
}

func GetEnvConfig() (EnvConfig, error) {
	cfg := EnvConfig{}
	err := env.Parse(&cfg)

	return cfg, err
}