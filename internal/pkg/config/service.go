package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Get retrieves the value of the environment variable identified by the given key.
// If the environment variable is not found, an empty string is returned.
func Get(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic("env variable not found")
	}

	return os.Getenv(key)
}

// GetWithDefault retrieves the value of the environment variable identified by the given key.
// If the environment variable is not found or its value is an empty string, the defaultValue is returned.
func GetWIthDefault(key string, defaultValue string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic("env variable not found")
	}

	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
