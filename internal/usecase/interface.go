package usecase

import (
	"context"

	"github.com/vpbuyanov/gw-backend-go/internal/models"
)

type UserUC interface {
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	CreateAdmin(ctx context.Context, id string) (*models.User, error)
	GetUser(ctx context.Context, id string) (*models.User, error)
	Login(ctx context.Context, email string, password string) (*models.User, error)
}
