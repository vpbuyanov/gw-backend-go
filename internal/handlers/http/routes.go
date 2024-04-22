package http

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/vpbuyanov/gw-backend-go/internal/models"
	"github.com/vpbuyanov/gw-backend-go/internal/usecase"
	"net/http"
)

type Routes struct {
	ctx context.Context
	usecase.UserUC
}

func New(ctx context.Context, uc usecase.UserUC) Routes {
	return Routes{
		ctx:    ctx,
		UserUC: uc,
	}
}

func (r *Routes) RegisterRoutes(app fiber.Router) {
	app.Get("/ping", r.Ping)
	app.Post("/create_user", r.CreateUser)
	app.Get("/get_user")
}

func (r *Routes) Ping(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
}

func (r *Routes) CreateUser(ctx *fiber.Ctx) error {
	var user models.User
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(struct {
				Error string `json:"error"`
			}{"can not parse body"})
	}

	createUser, err := r.UserUC.CreateUser(r.ctx, user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(struct {
			Error string `json:"error"`
		}{err.Error()})
	}

	return ctx.JSON(createUser)
}

func (r *Routes) GetUser(ctx *fiber.Ctx) error {
	var user models.GetUser
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(struct {
				Error string `json:"error"`
			}{"can not parse body"})
	}

	getUser, err := r.UserUC.GetUser(r.ctx, user.ID)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(struct {
			Error string `json:"error"`
		}{err.Error()})
	}

	if getUser == nil || getUser.UUID == "" {
		return ctx.Status(http.StatusNoContent).JSON(struct{}{})
	}

	return ctx.JSON(getUser)
}
