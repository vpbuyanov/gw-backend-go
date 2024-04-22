package configs

import (
	"os"
)

type (
	Config struct {
		Server   Server
		Postgres Postgres
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
)

func LoadConfig() Config {
	serverPort := os.Getenv("PORT")

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
