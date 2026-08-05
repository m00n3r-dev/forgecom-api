package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string

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
		DBHost:     getEnv("DBHost", "localhost"),
		DBPort:     getEnv("DBPort", "5432"),
		DBUser:     getEnv("DBUser", "root"),
		DBPassword: getEnv("DBPassword", "0612"),
		DBName:     getEnv("DBName", "forgecom"),
		DBSSLMode:  getEnv("DBSSLMode", "disable"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
