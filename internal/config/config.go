package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	User string
	Password string
	Database string
}

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
    	log.Fatalf("Error loading .env file")
  	}
}

func GetDBConfig() DBConfig {
	config := DBConfig {
		User: os.Getenv("db.user"),
		Password: os.Getenv("db.password"),
		Database: os.Getenv("db.database"),
	}
	return config
}

func GetPort() string {
	return os.Getenv("server.port")
}