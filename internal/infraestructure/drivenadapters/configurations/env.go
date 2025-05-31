package configurations

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func GetInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		valueAsIn, err := strconv.Atoi(value)

		if err != nil {
			return fallback
		}

		return valueAsIn
	}
	return fallback
}

func GetBool(key string, fallback bool) bool {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	// Convert the string value to a boolean
	value, err := strconv.ParseBool(val)
	if err != nil {
		return fallback
	}
	return value
}
