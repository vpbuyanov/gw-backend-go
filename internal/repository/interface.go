package repository

import (
	"context"

	"github.com/vpbuyanov/gw-backend-go/internal/models"
)

type UserRepos interface {
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	UpdateUser(ctx context.Context, user models.User, isAdmin bool) (*models.User, error)
	SelectUserByID(ctx context.Context, id string) (*models.User, error)
	SelectUserByEmail(ctx context.Context, username string) (*models.User, error)
	DeleteUser(ctx context.Context, id string) (*models.User, error)
}

type TopicRepos interface {
}

type CommentRepos interface {
}
