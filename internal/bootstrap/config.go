package bootstrap

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBUrl string
	Port  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		DBUrl: getEnv("DATABASE_URL", "root:thuong123@tcp(127.0.0.1:3306)/shopee_clone"),
		Port:  getEnv("PORT", "8080"),
	}
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
