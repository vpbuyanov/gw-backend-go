package app

import (
	"context"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/vpbuyanov/gw-backend-go/internal/configs"
	"github.com/vpbuyanov/gw-backend-go/internal/configure"
	"github.com/vpbuyanov/gw-backend-go/internal/logger"
	"github.com/vpbuyanov/gw-backend-go/internal/server"
	"github.com/vpbuyanov/gw-backend-go/internal/storage/postgresql"
	"github.com/vpbuyanov/gw-backend-go/internal/usecase"
)

type App struct {
	cfg *configs.Config
}

func New(cfg *configs.Config) *App {
	return &App{
		cfg: cfg,
	}
}

func (a *App) Run(ctx context.Context) {
	dbPool := configure.Postgres(ctx, a.cfg.Postgres)
	defer dbPool.Close()

	if err := a.cfg.Postgres.MigrationsUp(); err != nil && err.Error() != "no change" {
		panic(err)
	}

	logger.InitLogger(a.cfg.Logger)

	repos := postgresql.NewUserRepos(dbPool)
	userUC := usecase.NewUserUC(repos)

	runner := server.New(
		a.cfg,
		userUC,
	)

	err := runner.Start(ctx)
	if err != nil {
		logger.Log.Panicf("failed to start server: %v", err)
	}
}
