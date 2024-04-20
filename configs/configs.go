package configs

import (
	"os"
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
	serverPort := os.Getenv("PORT")

	pgUser := os.Getenv("POSTGRES_USER")
	pgPass := os.Getenv("POSTGRES_PASSWORD")
	pgDB := os.Getenv("POSTGRES_DB")

	rdbUser := os.Getenv("REDIS_USER")
	rdbPass := os.Getenv("REDIS_PASSWORD")

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
		Redis: Redis{
			Host:     "redis",
			Port:     "6379",
			User:     rdbUser,
			Password: rdbPass,
		},
	}
}
