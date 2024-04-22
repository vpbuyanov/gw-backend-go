package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/vpbuyanov/gw-backend-go/internal/models"
)

const (
	CreateUser        = `INSERT INTO "user" (name, email, hash_pass) VALUES($1, $2, $2) RETURNING *`
	SelectUserByID    = `SELECT * FROM "user" WHERE id=$1`
	SelectUserByEmail = `SELECT * FROM "user" WHERE email=$1`
	UpdateUser        = `UPDATE "user" SET name=$1, email=$2, hash_pass=$3, is_admin=$4 WHERE id=$5 RETURNING *`
	DeleteUser        = `DELETE FROM "user" WHERE id=$1 RETURNING *`
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
	query := u.db.QueryRow(ctx, CreateUser, user.Name, user.Email, user.HashPass)

	var res *models.User
	err := query.Scan(&res)
	if err != nil {
		return nil, fmt.Errorf("can not scan user for db: %w", err)
	}

	return res, nil
}

func (u *user) UpdateUser(ctx context.Context, user models.User, isAdmin bool) (*models.User, error) {
	query := u.db.QueryRow(ctx, UpdateUser, user.Name, user.Email, user.HashPass, isAdmin)

	var getUser *models.User
	err := query.Scan(&getUser)
	if err != nil {
		return nil, fmt.Errorf("can not scan user for update db: %w", err)
	}

	return getUser, nil
}

func (u *user) SelectUserByID(ctx context.Context, id string) (*models.User, error) {
	query := u.db.QueryRow(ctx, SelectUserByID, id)
	var getUser *models.User

	err := query.Scan(&getUser)
	if err != nil {
		return nil, fmt.Errorf("can not scan user for create db: %w", err)
	}

	return getUser, nil
}

func (u *user) SelectUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := u.db.QueryRow(ctx, SelectUserByEmail, email)
	var getUser *models.User
	err := query.Scan(&getUser)
	if err != nil {
		return nil, fmt.Errorf("can not select user by email: %w", err)
	}

	return getUser, nil
}

func (u *user) DeleteUser(ctx context.Context, id string) (*models.User, error) {
	query := u.db.QueryRow(ctx, DeleteUser, id)
	var getUser *models.User

	err := query.Scan(&getUser)
	if err != nil {
		return nil, fmt.Errorf("can not scan user for delete db: %w", err)
	}

	return getUser, nil
}
