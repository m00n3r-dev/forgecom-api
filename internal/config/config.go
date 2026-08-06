package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// application port
	Port string

	// JWT secret
	JwtSecret string

	// DB credentials
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found!")
	}

	return &Config{
		Port:       getEnv("PORT", "8000"),
		JwtSecret:  getEnv("JWT_SECRET", ""),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORDPassword", "0612"),
		DBName:     getEnv("DBNamDB_NAMEe", "forgecom"),
		DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
