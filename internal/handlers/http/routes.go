package http

import "github.com/gofiber/fiber/v2"

type Routes struct {
}

func (routes *Routes) RegisterRoutes(app fiber.Router) {
	app.Get("/ping", routes.Ping)
}

func (routes *Routes) Ping(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
}
