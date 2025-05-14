package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	ServerPort string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load(".env")

	cfg := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}

	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func validateConfig(cfg *Config) error {
	if cfg.DBHost == "" ||
		cfg.DBUser == "" ||
		cfg.DBPassword == "" ||
		cfg.DBName == "" ||
		cfg.DBPort == "" ||
		cfg.ServerPort == "" {
		return fmt.Errorf("missing required environment variables")
	}
	return nil
}
