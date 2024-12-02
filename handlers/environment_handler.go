package handlers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func LoadEnv(location ...string) {
	envPath := ".env"

	if len(location) > 0 && location[0] != "" {
		envPath = location[0]
	}

	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}
