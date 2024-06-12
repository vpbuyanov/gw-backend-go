package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/vpbuyanov/gw-backend-go/internal/entity"
	"github.com/vpbuyanov/gw-backend-go/internal/logger"
)

func (r *Routes) Login(ctx *fiber.Ctx) error {
	var user entity.LoginUserRequest
	err := ctx.BodyParser(&user)
	if err != nil {
		entity.ErrorWithAbort(ctx, http.StatusBadRequest, entity.ErrorParseBody, err)
		return nil
	}

	ok, err := r.userUC.Login(ctx.Context(), user.Email, user.HashPass)
	if err != nil {
		entity.ErrorWithAbort(ctx, http.StatusBadRequest, "can not login user", err)
		return nil
	}

	err = ctx.JSON(struct {
		OK bool `json:"ok"`
	}{ok})

	if err != nil {
		logger.Log.Errorf(entity.ErrorSendRequest, err)
		return nil
	}

	return nil
}

func (r *Routes) Registration(ctx *fiber.Ctx) error {
	var data entity.RegistrationUserRequest
	err := ctx.BodyParser(&data)
	if err != nil {
		entity.ErrorWithAbort(ctx, http.StatusBadRequest, entity.ErrorParseBody, err)
		return nil
	}

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
