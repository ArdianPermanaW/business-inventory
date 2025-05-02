// config/config.go
package config

import (
	"log"
	"os"
)

var (
	// Exported config values
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
)

func Init() {
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")

	if DBHost == "" || DBPort == "" || DBUser == "" || DBPassword == "" || DBName == "" {
		log.Println("Warning: One or more DB env vars are not set!")
	}
}
