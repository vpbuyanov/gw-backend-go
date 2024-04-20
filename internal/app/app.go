package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/gw-backend-go/configs"
	"github.com/vpbuyanov/gw-backend-go/internal/repository"
	"github.com/vpbuyanov/gw-backend-go/internal/server"
	"github.com/vpbuyanov/gw-backend-go/internal/usecase"
)

type App struct {
	log *logrus.Logger
	cfg configs.Config
}

func New(log *logrus.Logger, cfg configs.Config) *App {
	return &App{
		log: log,
		cfg: cfg,
	}
}

func (a *App) Run(ctx context.Context) {
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		a.cfg.Postgres.User, a.cfg.Postgres.Password, a.cfg.Postgres.Host, a.cfg.Postgres.Port, a.cfg.Postgres.DbName)

	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		a.log.Panicf("failed to connect to postgres: %v", err)
	}
	a.migrations(url)

	repos := repository.New(pool)
	userUC := usecase.NewUserUC(a.log, repos)

	runner := server.GetServer(
		a.cfg,
		userUC,
	)

	err = runner.Start()
	if err != nil {
		a.log.Panicf("failed to start server: %v", err)
	}
}

func (a *App) migrations(url string) {
	m, err := migrate.New("file://migrations", url)
	if err != nil {
		a.log.Panicf("failed to init migrations: %v", err)
	}

	err = m.Up()
	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			a.log.Panicf("failed to run migrations: %v", err)
		}
		a.log.Info("no change in migrations")
		return
	}

	a.log.Info("migrations is up")
}
