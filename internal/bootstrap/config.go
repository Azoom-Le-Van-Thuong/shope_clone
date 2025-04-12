package bootstrap

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	DBUrl                string
	Port                 string
	JWTAuthSecret        string
	JwtAccessTokenExpire time.Duration
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Load environment variables
	jwtAccessTokenExpireStr := getEnv("JWT_ACCESS_TOKEN_EXPIRE", "1h")

	// convert from 1h to 1h duration
	jwtAccessTokenExpire, err := time.ParseDuration(jwtAccessTokenExpireStr)
	if err != nil {
		log.Fatalf("invalid duration: %v", err)
	}
	return &Config{
		DBUrl:                getEnv("MYSQL_DATABASE_URL", "xx:xxx@tcp(127.0.0.1:3306)/shopee_clone"),
		Port:                 getEnv("PORT", "8080"),
		JWTAuthSecret:        getEnv("JWT_AUTH_SECRET", "secret"),
		JwtAccessTokenExpire: jwtAccessTokenExpire,
	}
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
