package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/vpbuyanov/gw-backend-go/configs"
	"github.com/vpbuyanov/gw-backend-go/internal/handlers/http"
	"github.com/vpbuyanov/gw-backend-go/internal/usecase"
)

type server struct {
	config configs.Config
	userUC usecase.UserUC
}

type Server interface {
	Start(ctx context.Context) error
}

func GetServer(config configs.Config, userUC usecase.UserUC) Server {
	return &server{
		config: config,
		userUC: userUC,
	}
}

func (s *server) Start(ctx context.Context) error {
	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api")

	routes := http.New(ctx, s.userUC)
	routes.RegisterRoutes(api)

	return app.Listen(s.config.Server.Port)
}
