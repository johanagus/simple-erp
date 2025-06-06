package config

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

func LoadEnv() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }
}

func GetEnv(key, fallback string) string {
    value := os.Getenv(key)
    if value == "" {
        return fallback
    }
    return value
}
