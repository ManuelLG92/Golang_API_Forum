package helpers

import (
	"github.com/joho/godotenv"
	"log"
)

var envs map[string]string

func LoadEnvs() {
	values, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	envs = values
}

func Get(key string) string {
	return envs[key]
}
