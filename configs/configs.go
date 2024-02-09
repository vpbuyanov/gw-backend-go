package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		Server Server
	}

	Server struct {
		Port string
	}
)

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	return Config{
		Server: Server{
			Port: port,
		},
	}
}
