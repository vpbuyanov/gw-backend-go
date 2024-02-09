package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/vpbuyanov/gw-backend-go/configs"
	"github.com/vpbuyanov/gw-backend-go/internal/handlers/http"
)

type server struct {
	config configs.Config
}

type Server interface {
	Start() error
}

func GetServer(config configs.Config) Server {
	return &server{config: config}
}

func (s *server) Start() error {
	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api")

	routes := http.Routes{}
	routes.RegisterRoutes(api)

	return app.Listen(s.config.Server.Port)
}
