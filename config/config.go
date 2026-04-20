package config

import (
	"log"
	"os"

	"github.com/dotenv-org/godotenvvault"
)

type Config struct {
	Port    string
	Debug   bool
	AppName string
}

func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func LoadConfig() *Config {
	err := godotenvvault.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &Config{
		Port:    getEnvWithDefault("PORT", "8080"),
		Debug:   getEnvWithDefault("DEBUG", "false") == "true",
		AppName: getEnvWithDefault("APP_NAME", "Cloud Duplicate File Finder"),
	}
	return config
}
