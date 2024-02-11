package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		Server   Server
		Postgres Postgres
		Redis    Redis
	}

	Server struct {
		Port string
	}

	Postgres struct {
		Port     string
		Host     string
		User     string
		Password string
		DbName   string
	}

	Redis struct {
		Port     string
		Host     string
		User     string
		Password string
	}
)

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serverPort := os.Getenv("PORT")

	err = godotenv.Load(".postgres.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pgUser := os.Getenv("POSTGRES_USER")
	pgPass := os.Getenv("POSTGRES_PASSWORD")
	pgDB := os.Getenv("POSTGRES_DB")

	return Config{
		Server: Server{
			Port: serverPort,
		},
		Postgres: Postgres{
			Port:     "5432",
			Host:     "postgres",
			Password: pgPass,
			User:     pgUser,
			DbName:   pgDB,
		},
	}
}
