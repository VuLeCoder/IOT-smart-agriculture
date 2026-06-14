package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	DatabaseURL    string
	JWTSecretKey   string
	JWTExpireHours int
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	expireHours, err := strconv.Atoi(os.Getenv("JWT_EXPIRE_HOURS"))
	if err != nil {
		return nil, err
	}

	cfg := Config{
		Port:           os.Getenv("PORT"),
		DatabaseURL:    os.Getenv("DATABASE_URL"),
		JWTSecretKey:   os.Getenv("SECRET_JWT_KEY"),
		JWTExpireHours: expireHours,
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

	if config.JWTSecretKey == "" {
		return errors.New("missing SECRET_JWT_KEY")
	}

	if config.JWTExpireHours <= 0 {
		return errors.New("JWT_EXPIRE_HOURS must be positive")
	}

	return nil
}
