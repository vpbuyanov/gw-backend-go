package redis

import (
	"context"

	"github.com/vpbuyanov/gw-backend-go/internal/entity"
	"github.com/vpbuyanov/gw-backend-go/internal/models"
)

type repos interface {
	SaveRefreshToken(ctx context.Context, userID int, refreshToken string)
	GetIDByRefreshToken(ctx context.Context, refreshToken string) *int
	DeleteRefreshToken(ctx context.Context, refreshToken string)
}

type userRepos interface {
	InsertUser(ctx context.Context, user entity.RegistrationUserRequest) (*int, error)
	SelectUserByID(ctx context.Context, id int) (*models.User, error)
	SelectUserByEmail(ctx context.Context, email string) (*models.User, error)
}
