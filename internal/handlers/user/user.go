package user

import (
	"github.com/vpbuyanov/gw-backend-go/internal/models"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/vpbuyanov/gw-backend-go/internal/entity"
)

type Handle struct {
	userUC  userUC
	redisUC redisUC
}

func NewHandleUser(userUC userUC, redisUC redisUC) *Handle {
	return &Handle{
		userUC:  userUC,
		redisUC: redisUC,
	}
}

func (r *Handle) Registration(ctx *fiber.Ctx) error {
	var data registrationUserRequest
	err := ctx.BodyParser(&data)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			entity.ErrorsRequest{
				Error:   err.Error(),
				Message: entity.ErrorParseBody,
				Status:  http.StatusBadRequest,
			})
	}

	id, err := r.userUC.Registration(ctx.Context(), models.User{
		Phone:    data.Phone,
		Email:    data.Email,
		HashPass: data.Password,
		Name:     data.Name,
		Surname:  data.Surname,
	})

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(entity.ErrorsRequest{
			Error:   err.Error(),
			Message: "can not create user",
			Status:  http.StatusBadRequest,
		})
	}

	refreshToken := r.redisUC.CreateRefreshToken(ctx.Context(), *id)

	ctx.Locals("UserID", id)
	ctx.Locals("RefreshToken", refreshToken)

	return ctx.Next()
}

func (r *Handle) Login(ctx *fiber.Ctx) error {
	var user loginUserRequest
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(entity.ErrorsRequest{
			Error:   err.Error(),
			Message: entity.ErrorParseBody,
			Status:  http.StatusBadRequest,
		})
	}

	getUser, err := r.userUC.Login(ctx.Context(), user.Email, user.Password)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(entity.ErrorsRequest{
			Error:   err.Error(),
			Message: "can not login user",
			Status:  http.StatusBadRequest,
		})
	}

	refreshToken := r.redisUC.CreateRefreshToken(ctx.Context(), getUser.ID)

	ctx.Locals("UserID", getUser.ID)
	ctx.Locals("RefreshToken", refreshToken)

	return ctx.Next()
}
