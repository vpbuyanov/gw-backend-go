package usecase

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/gw-backend-go/internal/models"
	"github.com/vpbuyanov/gw-backend-go/internal/repository"
)

type userUC struct {
	log   *logrus.Logger
	repos repository.UserRepos
}

func NewUserUC(log *logrus.Logger, repos repository.UserRepos) UserUC {
	return &userUC{
		log:   log,
		repos: repos,
	}
}

func (u *userUC) CreateUser(ctx context.Context, user models.User) error {
	_, err := u.repos.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUC) CreateAdmin(ctx context.Context, id string) error {

	return nil
}
