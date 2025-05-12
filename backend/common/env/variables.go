package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
)

func Load() {
	godotenv.Load(".env")

	DB_HOST = os.Getenv("DB_HOST")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PORT = os.Getenv("DB_PORT")
}
