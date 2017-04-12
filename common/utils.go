package common

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"encoding/json"
)

type configuration struct {
	Server, MongoDBHost, DBUser, DBPwd, Database string
}

// AppConfig holds the configuration values from config.json file
var AppConfig configuration

// Initialize AppConfig
func initConfig() {
	loadAppConfig()
}


func loadAppConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}
