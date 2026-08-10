package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load .env file")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}