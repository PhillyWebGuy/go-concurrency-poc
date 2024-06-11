package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	//This line reads the .env file and loads the environment variables into the system.
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
