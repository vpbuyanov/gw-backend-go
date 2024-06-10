package usecase

import (
	"context"

	"github.com/vpbuyanov/gw-backend-go/internal/common"
	"github.com/vpbuyanov/gw-backend-go/internal/logger"
	"github.com/vpbuyanov/gw-backend-go/internal/models"
	"github.com/vpbuyanov/gw-backend-go/internal/storage/postgresql"
)

type UserUC struct {
	repos *postgresql.UserRepos
}

func NewUserUC(repos *postgresql.UserRepos) *UserUC {
	return &UserUC{
		repos: repos,
	}
}

func (u *UserUC) CreateUser(ctx context.Context, user models.User) error {
	const funcName = "SelectUser UserUC"

	pass, err := common.CreateHashPassword(user.HashPass)
	if err != nil {
		return err
	}

	user.HashPass = pass

	err = u.repos.InsertUser(ctx, user)
	if err != nil {
		logger.Log.Errorf("[%v] can not create user: %v", funcName, err.Error())
		return err
	}

	return nil
}

func (u *UserUC) UpdateUserToAdmin(ctx context.Context, id string) (*models.User, error) {
	const funcName = "UpdateUserToAdmin UserUC"

	user, err := u.repos.SelectUserByID(ctx, id)
	if err != nil {
		logger.Log.Errorf("[%v] can not select user in databases || user not found: %v", funcName, err.Error())
		return nil, err
	}

	user.IsAdmin = true

	res, err := u.repos.UpdateUser(ctx, *user)
	if err != nil {
		logger.Log.Errorf("[%v] create admin: %v", funcName, err.Error())
		return nil, err
	}

	return res, nil
}

func (u *UserUC) GetUser(ctx context.Context, id string) (*models.User, error) {
	const funcName = "GetUser UserUC"

	user, err := u.repos.SelectUserByID(ctx, id)
	if err != nil {
		logger.Log.Errorf("[%v] select user: %v", funcName, err.Error())
		return nil, err
	}

	return user, nil
}

func (u *UserUC) Login(ctx context.Context, email, password string) (bool, error) {
	const funcName = "Login UserUC"

	user, err := u.repos.SelectUserByEmail(ctx, email)
	if err != nil {
		logger.Log.Errorf("[%v] can not select user by email: %v", funcName, err.Error())
		return false, err
	}

	ok, err := common.CompareHashAndPassword(user.HashPass, password)
	if err != nil {
		return false, err
	}

	return ok, nil
}
