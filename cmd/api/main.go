package main

import (
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
