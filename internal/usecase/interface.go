package usecase

import (
	"context"

	"github.com/vpbuyanov/gw-backend-go/internal/models"
)

type UserUC interface {
	CreateUser(ctx context.Context, user models.User) error
	CreateAdmin(ctx context.Context, id string) error
}
