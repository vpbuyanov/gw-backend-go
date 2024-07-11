package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/vpbuyanov/gw-backend-go/internal/common"
	"github.com/vpbuyanov/gw-backend-go/internal/models"
)

type UCUser struct {
	user userRepos
}

func NewUCUser(user userRepos) *UCUser {
	return &UCUser{
		user: user,
	}
}

func (u *UCUser) Registration(ctx context.Context, request models.User) (*int, error) {
	pass, err := common.CreateHashPassword(request.HashPass)
	if err != nil {
		return nil, err
	}

	request.HashPass = pass

	id, err := u.user.InsertUser(ctx, request)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (u *UCUser) Login(ctx context.Context, email, password string) (*models.User, error) {
	user, err := u.user.SelectUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	err = common.CompareHashAndPassword(user.HashPass, password)
	if err != nil {
		return nil, errors.New("wrong password or email")
	}

	return user, nil
}
