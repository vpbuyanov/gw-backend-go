package server

import (
	"context"

	"github.com/gofiber/fiber/v2"
	flogger "github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/vpbuyanov/gw-backend-go/internal/configs"
	"github.com/vpbuyanov/gw-backend-go/internal/handlers/http"
	"github.com/vpbuyanov/gw-backend-go/internal/logger"
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
	app.Use(flogger.New(flogger.Config{
		TimeZone:   "Europe/Moscow",
		TimeFormat: "2 Jan 2006 15:04:05",
	}))

	api := app.Group("/api")

	routes := http.New(s.userUC)
	routes.RegisterRoutes(api)

	err := app.Listen(s.config.Server.String())
	if err != nil {
		logger.Log.Errorf("err Listen: %v", err)
		return nil
	}

	return nil
}
