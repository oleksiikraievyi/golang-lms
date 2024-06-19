package server

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Env  string `env:"ENV" envDefault:"development"`
	Port int    `env:"PORT" envDefault:"8080"`
}

func GetConfig() *Config {
	godotenv.Load(".env")

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}
	return &cfg
}
