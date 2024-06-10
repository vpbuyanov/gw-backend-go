package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/vpbuyanov/gw-backend-go/internal/entity"
	"github.com/vpbuyanov/gw-backend-go/internal/models"
)

const (
	SelectUserByID    = `SELECT * FROM "user" WHERE id=$1`
	SelectUserByEmail = `SELECT * FROM "user" WHERE email=$1`
	DeleteUser        = `DELETE FROM "user" WHERE id=$1 RETURNING *`
)

type UserRepos struct {
	db *pgxpool.Pool
}

func NewUserRepos(db *pgxpool.Pool) *UserRepos {
	return &UserRepos{
		db: db,
	}
}

func (u *UserRepos) InsertUser(ctx context.Context, user entity.RegistrationUserRequest) (*models.User, error) {
	const (
		query = `
			insert into "user" (name, email, phone, hash_pass) 
			values($1, $2, $3, $4) returning *`
	)

	transaction, err := u.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = transaction.Rollback(ctx)
	}()

	var res models.User
	err = transaction.QueryRow(ctx, query, user.Name, user.Email, user.Phone, user.HashPass).
		Scan(
			&res.UUID,
			&res.Name,
			&res.Email,
			&res.HashPass,
			&res.IsAdmin)

	if err != nil {
		return nil, fmt.Errorf("can not scan UserRepos for db: %w", err)
	}

	if err = transaction.Commit(ctx); err != nil {
		return nil, err
	}

	return &res, nil
}

func (u *UserRepos) UpdateUser(ctx context.Context, user models.User) (*models.User, error) {
	const (
		query = `
			UPDATE "user" SET name=$1, email=$2,
			phone=$3, hash_pass=$4, is_admin=$5, is_blocked=$6 
			WHERE id=$7 RETURNING *`
	)

	var getUser models.User
	err := u.db.QueryRow(ctx, query, user.Name, user.Email, user.Phone, user.HashPass, user.IsAdmin, user.IsBanned).
		Scan(
			&getUser.UUID,
			&getUser.Name,
			&getUser.Email,
			&getUser.Phone,
			&getUser.HashPass,
			&getUser.IsAdmin,
			&getUser.IsBanned)

	if err != nil {
		return nil, fmt.Errorf("can not scan UserRepos for update db: %w", err)
	}

	return &getUser, nil
}

func (u *UserRepos) SelectUserByID(ctx context.Context, id string) (*models.User, error) {
	var getUser models.User
	err := u.db.QueryRow(ctx, SelectUserByID, id).
		Scan(
			&getUser.UUID,
			&getUser.Name,
			&getUser.Email,
			&getUser.HashPass,
			&getUser.IsAdmin)

	if err != nil {
		return nil, fmt.Errorf("can not scan UserRepos for create db: %w", err)
	}

	return &getUser, nil
}

func (u *UserRepos) SelectUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var getUser models.User
	err := u.db.QueryRow(ctx, SelectUserByEmail, email).Scan(
		&getUser.UUID,
		&getUser.Name,
		&getUser.Email,
		&getUser.HashPass,
		&getUser.IsAdmin)

	if err != nil {
		return nil, fmt.Errorf("can not select UserRepos by email: %w", err)
	}

	return &getUser, nil
}

func (u *UserRepos) DeleteUser(ctx context.Context, id string) (*models.User, error) {
	var getUser models.User
	err := u.db.QueryRow(ctx, DeleteUser, id).
		Scan(
			&getUser.UUID,
			&getUser.Name,
			&getUser.Email,
			&getUser.HashPass,
			&getUser.IsAdmin)

	if err != nil {
		return nil, fmt.Errorf("can not scan UserRepos for delete db: %w", err)
	}

	return &getUser, nil
}

func (u *UserRepos) IsAdmin(ctx context.Context, id string) (bool, error) {
	const (
		query = `select exists (select 1 from "user" where id = $1 and is_admin = true and is_blocked = false)`
	)

	var isAdmin bool
	err := u.db.QueryRow(ctx, query, id).Scan(&isAdmin)
	if err != nil {
		return false, err
	}

	return isAdmin, nil
}
