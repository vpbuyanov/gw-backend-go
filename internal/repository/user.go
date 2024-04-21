package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/vpbuyanov/gw-backend-go/internal/common"
	"github.com/vpbuyanov/gw-backend-go/internal/models"
)

const (
	CreateUser = `INSERT INTO "user" (name, email, hash_pass) VALUES(?, ?, ?) RETURNING *`
	SelectUser = `SELECT * FROM "user" WHERE id=?`
	UpdateUser = `UPDATE "user" SET name=?, email=?, hash_pass=?, is_admin=? WHERE id=? RETURNING *`
	DeleteUser = `DELETE FROM "user" WHERE id=? RETURNING *`
)

type user struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) UserRepos {
	return &user{
		db: db,
	}
}

func (u *user) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	query := u.db.QueryRow(ctx, CreateUser, user.Name, user.Email, common.CreateHash(user.HashPass))

	var res *models.User
	err := query.Scan(&res)
	if err != nil {
		return nil, fmt.Errorf("can not scan user for db: %w", err)
	}

	if res != nil {
		return res, nil
	}

	return nil, errors.New("can not create user")
}

func (u *user) CreateAdmin(ctx context.Context, user models.User) (*models.User, error) {
	query := u.db.QueryRow(ctx, UpdateUser, user.Name, user.Email, user.HashPass, "true")

	var getUser *models.User
	err := query.Scan(&getUser)
	if err != nil {
		return nil, fmt.Errorf("can not scan user for update db: %w", err)
	}

	return getUser, nil
}

func (u *user) SelectUserByID(ctx context.Context, id string) (*models.User, error) {
	query := u.db.QueryRow(ctx, SelectUser, id)
	var getUser *models.User

	err := query.Scan(&getUser)
	if err != nil {
		return nil, fmt.Errorf("can not scan user for create db: %w", err)
	}

	return getUser, nil
}

func (u *user) DeleteUser(ctx context.Context, id string) error {
	query := u.db.QueryRow(ctx, DeleteUser, id)
	var getUser *models.User

	err := query.Scan(&getUser)
	if err != nil {
		return fmt.Errorf("can not scan user for delete db: %w", err)
	}

	return getUser, nil
}
