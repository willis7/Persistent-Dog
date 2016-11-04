package common

import (
	"github.com/joho/godotenv"
	"log"
)

// AppConfig holds the configuration values from config.json file
var AppConfig map[string]string

// Initialize AppConfig
func initConfig() {
	loadAppConfig()
}


func loadAppConfig() {
	AppConfig, err := godotenv.Read()
	if err != nil {
		log.Fatalf("Failed to read dotenv: %v", err)
	}
}
