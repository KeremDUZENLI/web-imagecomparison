package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string

	SERVER_PORT string
)

func Load() {
	godotenv.Load(".env")

	DB_HOST = os.Getenv("DB_HOST")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PORT = os.Getenv("DB_PORT")

	SERVER_PORT = os.Getenv("SERVER_PORT")

	validateConfig()
}

func validateConfig() {
	required := map[string]string{
		"DB_HOST":     DB_HOST,
		"DB_USER":     DB_USER,
		"DB_PASSWORD": DB_PASSWORD,
		"DB_NAME":     DB_NAME,
		"DB_PORT":     DB_PORT,

		"SERVER_PORT": SERVER_PORT,
	}
	for key, val := range required {
		if val == "" {
			fmt.Fprintf(os.Stderr, "environment variable %s is required\n", key)
			os.Exit(1)
		}
	}
}
