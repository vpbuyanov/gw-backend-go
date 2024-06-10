package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/vpbuyanov/gw-backend-go/internal/middleware"
	"github.com/vpbuyanov/gw-backend-go/internal/usecase"
)

type Routes struct {
	userUC *usecase.UserUC
}

func New(userUC *usecase.UserUC) Routes {
	return Routes{
		userUC: userUC,
	}
}

func (r *Routes) RegisterRoutes(app fiber.Router) {
	app.Get("/ping", r.Ping)

	user := app.Group("/user")
	admin := app.Group("/admin")

	user.Post("/login", r.Login)
	user.Post("/registration", r.Registration)

	authUser := user.Group("/auth", middleware.CompareToken)
	authUser.Post("/create_admin", r.CreateAdmin)
	authUser.Get("/get_user_by_id", r.GetUserByID)

	admin.Post("/login")

	authAdmin := admin.Group("/auth", middleware.CompareAdminToken)
	authAdmin.Post("/update_user")
}

func (r *Routes) Ping(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
}
