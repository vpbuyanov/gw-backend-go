package usecase

import (
	"context"
	"fmt"

	"github.com/vpbuyanov/gw-backend-go/internal/common"
	"github.com/vpbuyanov/gw-backend-go/internal/entity"
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

func (u *UserUC) CreateUser(ctx context.Context, user entity.RegistrationUserRequest) (*models.User, error) {
	const funcName = "SelectUser UserUC"

	pass, err := common.CreateHashPassword(user.HashPass)
	if err != nil {
		return nil, fmt.Errorf("[%v] can not create hash pass: %w", funcName, err)
	}

	user.HashPass = pass

	getUsers, err := u.repos.InsertUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("[%v] can not create user: %w", funcName, err)
	}

	return getUsers, nil
}

func (u *UserUC) UpdateUserToAdmin(ctx context.Context, id string) (*models.User, error) {
	const funcName = "UpdateUserToAdmin UserUC"

	user, err := u.repos.SelectUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("[%v] can not select user in databases || user not found: %w", funcName, err)
	}

	user.IsAdmin = true

	res, err := u.repos.UpdateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("[%v] can not update user to admin: %w", funcName, err)
	}

	return res, nil
}

func (u *UserUC) GetUser(ctx context.Context, id string) (*models.User, error) {
	const funcName = "GetUser UserUC"

	user, err := u.repos.SelectUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("[%v] select user: %w", funcName, err)
	}

	return user, nil
}

func (u *UserUC) Login(ctx context.Context, email, password string) (bool, error) {
	const funcName = "Login UserUC"

	user, err := u.repos.SelectUserByEmail(ctx, email)
	if err != nil {
		return false, fmt.Errorf("[%v] can not select user by email: %w", funcName, err)
	}

	ok, err := common.CompareHashAndPassword(user.HashPass, password)
	if err != nil {
		return false, fmt.Errorf("[%v] can not select user by email: %w", funcName, err)
	}

	return ok, nil
}
