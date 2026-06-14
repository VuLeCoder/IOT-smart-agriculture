package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
	SecretKey   string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := Config{
		Port:        os.Getenv("PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		SecretKey:   os.Getenv("SECRET_JWT_KEY"),
	}

	if err := validateConfig(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func validateConfig(config *Config) error {
	if config.Port == "" {
		return errors.New("missing PORT")
	}

	if config.DatabaseURL == "" {
		return errors.New("missing DATABASE_URL")
	}

	if config.SecretKey == "" {
		return errors.New("missing SECRET_JWT_KEY")
	}

	return nil
}
