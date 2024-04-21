package repository

import (
	"context"

	"github.com/vpbuyanov/gw-backend-go/internal/models"
)

type UserRepos interface {
	CreateUser(ctx context.Context, user models.User) (*models.User, error)

	CreateAdmin(ctx context.Context, user models.User) (*models.User, error)
	SelectUserByID(ctx context.Context, id string) (*models.User, error)
}

type TopicRepos interface {
}

type CommentRepos interface {
}
