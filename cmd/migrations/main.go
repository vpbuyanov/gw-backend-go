package main

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/gw-backend-go/configs"
)

func main() {
	config := configs.LoadConfig()

	migrateUp(config)
}

func migrateUp(conf configs.Config) {
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
