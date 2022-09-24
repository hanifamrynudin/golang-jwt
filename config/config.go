package config

import (
	"golang-jwt/models"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() models.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbUsername := os.Getenv("DB_USERNAME")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	serverPort := os.Getenv("SERVICE_PORT")
	config := models.Config{
		ServerPort: serverPort,
		Database: models.Database{
			Username: dbUsername,
			Password: dbPassword,
			Host:     dbHost,
			Port:     dbPort,
			Name:     dbName,
		},
	}
	return config
}