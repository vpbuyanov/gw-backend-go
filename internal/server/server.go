package server

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/vpbuyanov/gw-backend-go/internal/configs"
	"github.com/vpbuyanov/gw-backend-go/internal/handlers/http"
	"github.com/vpbuyanov/gw-backend-go/internal/usecase"
)

type Server struct {
	config *configs.Config
	userUC *usecase.UserUC
}

func New(config *configs.Config, userUC *usecase.UserUC) *Server {
	return &Server{
		config: config,
		userUC: userUC,
	}
}

func (s *Server) Start(ctx context.Context) error {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		TimeZone:   "Europe/Moscow",
		TimeFormat: "2 Jan 2006 15:04:05",
	}))

	api := app.Group("/api")

	routes := http.New(s.userUC)
	routes.RegisterRoutes(api)

	return app.Listen(s.config.Server.String())
}
