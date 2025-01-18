package helpers

import (
	"log"

	"github.com/joho/godotenv"
)

var Env = map[string]string{}

func SetupConfig() {
	var err error
	Env, err = godotenv.Read(".env")
	if err != nil {
		log.Fatal(`failed to read env file : `, err)
	}
}

func GetEnv(key, defaultValue string) string {
	result := Env[key]
	if result == "" {
		result = defaultValue
	}
	return result
}
