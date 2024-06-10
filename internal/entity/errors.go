package entity

import (
	"github.com/gofiber/fiber/v2"

	"github.com/vpbuyanov/gw-backend-go/internal/logger"
)

const (
	ErrorSendRequest = "can not send request, err: %v"
	ErrorParseBody   = "can not parse body"
)

type errorsRequest struct {
	Error   error  `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func BindByError(ctx *fiber.Ctx, status int, message string, customError error) {
	err := ctx.Status(status).JSON(errorsRequest{
		Error:   customError,
		Message: message,
		Status:  status,
	})

	if err != nil {
		logger.Log.Errorf(ErrorSendRequest, err)
		return
	}
}
