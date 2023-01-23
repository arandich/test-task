package config

import (
	"os"
	"strconv"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DB       string
	Sslmode  string
}

func GetConfig() *Config {
	return &Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     getEnvAsInt("POSTGRES_PORT", 5432),
		User:     "postgres",
		Password: "admin",
		DB:       "tests",
		Sslmode:  os.Getenv("SSLMODE"),
	}
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
