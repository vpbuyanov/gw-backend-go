package usecase

import (
	"context"
	"fmt"

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
	const funcName = "SelectUser UserUC"

	_, err := u.repos.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("[%v] create user: %w", funcName, err)
	}
	return nil
}

func (u *userUC) CreateAdmin(ctx context.Context, id string) error {
	const funcName = "CreateAdmin UserUC"

	user, err := u.repos.SelectUserByID(ctx, id)
	if err != nil {
		u.log.Errorf("[%v] create user: %v", funcName, err.Error())
		return err
	}

	_, err = u.repos.CreateAdmin(ctx, *user)
	if err != nil {
		u.log.Errorf("[%v] create admin: %v", funcName, err.Error())
		return err
	}

	return nil
}

func (u *userUC) SelectUser(ctx context.Context, id string) (*models.User, error) {
	const funcName = "SelectUser UserUC"

	user, err := u.repos.SelectUserByID(ctx, id)
	if err != nil {
		u.log.Errorf("[%v] select user: %v", funcName, err.Error())
		return nil, err
	}

	return user, nil
}
