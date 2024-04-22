package usecase

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/gw-backend-go/internal/common"
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

func (u *userUC) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	const funcName = "SelectUser UserUC"

	pass, err := common.CreateHashPassword(user.HashPass)
	if err != nil {
		return nil, err
	}

	user.HashPass = pass

	res, err := u.repos.CreateUser(ctx, user)
	if err != nil {
		u.log.Errorf("[%v] can not create user: %v", funcName, err.Error())
		return nil, err
	}

	if res == nil || len(res.UUID) == 0 {
		u.log.Errorf("[%v] not user returning", funcName)
		return nil, err
	}

	return res, nil
}

func (u *userUC) CreateAdmin(ctx context.Context, id string) (*models.User, error) {
	const funcName = "CreateAdmin UserUC"

	user, err := u.repos.SelectUserByID(ctx, id)
	if err != nil {
		u.log.Errorf("[%v] can not select user in databases || user not found: %v", funcName, err.Error())
		return nil, err
	}

	res, err := u.repos.UpdateUser(ctx, *user, true)
	if err != nil {
		u.log.Errorf("[%v] create admin: %v", funcName, err.Error())
		return nil, err
	}

	return res, nil
}

func (u *userUC) GetUser(ctx context.Context, id string) (*models.User, error) {
	const funcName = "SelectUser UserUC"

	user, err := u.repos.SelectUserByID(ctx, id)
	if err != nil {
		u.log.Errorf("[%v] select user: %v", funcName, err.Error())
		return nil, err
	}

	return user, nil
}

func (u *userUC) Login(ctx context.Context, email string, password string) (*models.User, error) {
	const funcName = "Login UserUC"

	user, err := u.repos.SelectUserByEmail(ctx, email)
	if err != nil {
		u.log.Errorf("[%v] can not select user by email: %v", funcName, err.Error())
		return nil, err
	}

	ok, err := common.CompareHashAndPassword(user.HashPass, password)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
