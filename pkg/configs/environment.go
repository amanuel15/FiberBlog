package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DB_NAME string
var DB_HOST string
var DB_PORT string
var DB_USER string
var DB_PASSWORD string
var PORT string
var JWT_SECRET []byte

func SetupEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DB_HOST = os.Getenv("DB_HOST")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PORT = os.Getenv("DB_PORT")
	PORT = os.Getenv("PORT")
	JWT_SECRET = []byte(os.Getenv("JWT_SECRET"))
}
