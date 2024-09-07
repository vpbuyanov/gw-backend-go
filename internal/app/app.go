package app

import (
	"context"

	"github.com/gofiber/fiber/v2"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	mailerBroker "github.com/vpbuyanov/gw-backend-go/internal/broker/mailer"
	"github.com/vpbuyanov/gw-backend-go/internal/configs"
	"github.com/vpbuyanov/gw-backend-go/internal/configure"
	handle "github.com/vpbuyanov/gw-backend-go/internal/handlers/user"
	"github.com/vpbuyanov/gw-backend-go/internal/logger"
	"github.com/vpbuyanov/gw-backend-go/internal/middleware/log"
	"github.com/vpbuyanov/gw-backend-go/internal/middleware/token"
	"github.com/vpbuyanov/gw-backend-go/internal/models"
	"github.com/vpbuyanov/gw-backend-go/internal/service"
	"github.com/vpbuyanov/gw-backend-go/internal/storage/postgresql"
	"github.com/vpbuyanov/gw-backend-go/internal/storage/redis"
	"github.com/vpbuyanov/gw-backend-go/internal/usecase/mailer"
	redisUC "github.com/vpbuyanov/gw-backend-go/internal/usecase/redis"
	"github.com/vpbuyanov/gw-backend-go/internal/usecase/user"
)

const (
	sizeBufferChannel = 256
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
	app := fiber.New()
	app.Use(log.New())

	logger.InitLogger(a.cfg.Logger)

	// DB
	dbPool := configure.Postgres(ctx, a.cfg.Postgres)
	defer dbPool.Close()

	if err := a.cfg.Postgres.MigrationsUp(); err != nil {
		logger.Log.Errorf("can not up migration in postgres, err: %v", err)
	}

	redisDB := configure.Redis(a.cfg.Redis)

	// Channel
	channelMailer := make(chan models.Gmail, sizeBufferChannel)

	// Repos
	userRepos := postgresql.NewUserRepos(dbPool)
	redisRepos := redis.NewTokenRepos(redisDB)

	// UseCase
	ucUser := user.NewUCUser(userRepos)
	ucMailer := mailer.New(a.cfg.Mailer, channelMailer)
	ucRedis := redisUC.NewUCRepos(redisRepos)

	// Broker
	brokerMailer := mailerBroker.NewBrokerMailer(ucMailer, channelMailer)

	// Handler
	userHandler := handle.NewHandleUser(ucUser, ucRedis)

	// endpoint
	api := app.Group("/api")

	users := api.Group("/user")

	users.Post("/registration", userHandler.Registration, token.SignedToken)
	users.Post("/login", userHandler.Login, token.SignedToken)

	// Run
	service.RegisterBroker(brokerMailer.Run)

	err := app.Listen(a.cfg.Server.String())
	if err != nil {
		logger.Log.Errorf("err Listen: %v", err)

		return
	}
}
