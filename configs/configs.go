package configs

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
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
		logrus.Fatal("Error loading .env file")
	}

	serverPort := os.Getenv("PORT")

	err = godotenv.Load(".postgres.env")
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	pgUser := os.Getenv("POSTGRES_USER")
	pgPass := os.Getenv("POSTGRES_PASSWORD")
	pgDB := os.Getenv("POSTGRES_DB")

	err = godotenv.Load(".redis.env")
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

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
