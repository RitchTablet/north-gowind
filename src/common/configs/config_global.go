package config_global

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVar(envVar string) string {
	return os.Getenv(envVar)
}

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		erroMessage := "Error loading .env file"
		log.Fatal(erroMessage)
		panic(erroMessage)
	}
}
