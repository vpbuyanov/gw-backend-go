package http

import (
	"github.com/vpbuyanov/gw-backend-go/internal/models"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/vpbuyanov/gw-backend-go/internal/entity"
)

func (r *Routes) Login(ctx *fiber.Ctx) error {
	var user entity.LoginUserRequest
	err := ctx.BodyParser(&user)
	if err != nil {
		entity.BindByError(ctx, http.StatusBadRequest, entity.ErrorParseBody, err)
		return nil
	}

	ok, err := r.userUC.Login(ctx.Context(), user.Email, user.HashPass)
	if err != nil {
		entity.BindByError(ctx, http.StatusBadRequest, "can not login user", err)
		return nil
	}

	return ctx.JSON(struct {
		OK bool `json:"ok"`
	}{ok})
}

func (r *Routes) Registration(ctx *fiber.Ctx) error {
	var data entity.RegistrationUserRequest
	err := ctx.BodyParser(&data)
	if err != nil {
		entity.BindByError(ctx, http.StatusBadRequest, entity.ErrorParseBody, err)
		return nil
	}

	var user models.User
	user.Name = data.Name
	user.Email = data.Email
	user.Phone = data.Phone
	user.HashPass = data.HashPass
	user.IsAdmin = false
	user.IsBanned = false

	return nil
}

func (r *Routes) CreateAdmin(ctx *fiber.Ctx) error {
	// TODO: create admin
	return nil
}

func (r *Routes) GetUserByID(ctx *fiber.Ctx) error {
	// TODO: get user
	return nil
}
