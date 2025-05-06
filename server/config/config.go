package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	ServerPort     string
	DBHost         string
	DBPort         string
	DBName         string
	DBUser         string
	DBPassword     string
	DBSSLMode      string
	AgifyURL       string
	GenderizeURL   string
	NationalizeURL string
	LogLevel       string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	return &Config{
		ServerPort:     getEnv("SERVER_PORT", "8081"),
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5433"),
		DBName:         getEnv("DB_NAME", "person_db"),
		DBUser:         getEnv("DB_USER", "person_user"),
		DBPassword:     getEnv("DB_PASSWORD", "person_password"),
		DBSSLMode:      getEnv("DB_SSL_MODE", "disable"),
		AgifyURL:       getEnv("AGIFY_URL", "https://api.agify.io"),
		GenderizeURL:   getEnv("GENDERIZE_URL", "https://api.genderize.io"),
		NationalizeURL: getEnv("NATIONALIZE_URL", "https://api.nationalize.io"),
		LogLevel:       getEnv("LOG_LEVEL", "debug"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
