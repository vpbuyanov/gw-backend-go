package postgresql

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vpbuyanov/gw-backend-go/internal/common"

	"github.com/vpbuyanov/gw-backend-go/internal/databases/postgres"
	"github.com/vpbuyanov/gw-backend-go/internal/models"
)

type userUC struct {
	repos postgres.Postgresql
}

type UserUC interface {
	SelectUserByID(ctx context.Context, id string) (*models.User, error)
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	UpdateUser(ctx context.Context, user models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id string) error
}

func NewUserUC(repos postgres.Postgresql) UserUC {
	return &userUC{
		repos: repos,
	}
}

func (u *userUC) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	query, err := u.repos.Query(ctx, postgres.CreateUser, user.Name, user.Email, common.HashFunction(user.HashPass))
	if err != nil {
		return nil, fmt.Errorf("can not create user: %w", err)
	}

	var res models.User
	val := query.RawValues()

	err = json.Unmarshal(val[0], &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user: %w", err)
	}

	return &res, nil
}

func (u *userUC) SelectUserByID(ctx context.Context, id string) (*models.User, error) {
	query, err := u.repos.Query(ctx, postgres.SelectUser, id)
	if err != nil {
		return nil, fmt.Errorf("select user: %w", err)
	}

	var res models.User
	val := query.RawValues()

	err = json.Unmarshal(val[0], &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user: %w", err)
	}

	return &res, nil
}

func (u *userUC) UpdateUser(ctx context.Context, user models.User) (*models.User, error) {
	return nil, nil
}

func (u *userUC) DeleteUser(ctx context.Context, id string) error {
	return nil
}
