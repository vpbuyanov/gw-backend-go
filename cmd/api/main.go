package main

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/gw-backend-go/configs"
	"github.com/vpbuyanov/gw-backend-go/internal/databases/postgres"
	"github.com/vpbuyanov/gw-backend-go/internal/databases/redis"
	"github.com/vpbuyanov/gw-backend-go/internal/server"
)

func main() {
	config := configs.LoadConfig()
	runner := server.GetServer(config)

	pg := postgres.NewReposPostgresql(config.Postgres)
	dbRedis := redis.NewReposRedis(config.Redis)

	err := pg.Connect()
	if err != nil {
		logrus.WithError(err).Println("can't connect to pg")
	}
	logrus.Println("pg connect")

	migrateUP(config)

	err = dbRedis.Connect()
	if err != nil {
		logrus.WithError(err).Println("can't connect to redis")
	}

	logrus.Println("redis connect")

	err = runner.Start()
	if err != nil {
		return
	}
}

func migrateUP(conf configs.Config) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conf.Postgres.User, conf.Postgres.Password, conf.Postgres.Host, conf.Postgres.Port, conf.Postgres.DbName)

	m, err := migrate.New(
		"file://migrations",
		url)
	if err != nil {
		logrus.Fatal(err)
	}

	err = m.Up()
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Println("migrations is up")
}
