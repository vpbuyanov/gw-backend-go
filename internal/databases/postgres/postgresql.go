package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/vpbuyanov/gw-backend-go/configs"
)

type postgresql struct {
	url    string
	client *pgx.Conn
}

type Postgresql interface {
	Connect(ctx context.Context) error
	Close(ctx context.Context) error
	Query(ctx context.Context, request string, args ...string) (pgx.Rows, error)
}

func NewReposPostgresql(configs configs.Postgres) Postgresql {
	return &postgresql{
		url: fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
			configs.User, configs.Password, configs.Host, configs.Port, configs.DbName),
	}
}

func (r *postgresql) Connect(ctx context.Context) error {
	client, err := pgx.Connect(ctx, r.url)
	if err != nil {
		return fmt.Errorf("could not connect to postgres: %w", err)
	}

	r.client = client
	return nil
}

func (r *postgresql) Close(ctx context.Context) error {
	err := r.client.Close(ctx)
	if err != nil {
		return fmt.Errorf("could not close postgres connection: %w", err)
	}

	return nil
}

func (r *postgresql) Query(ctx context.Context, request string, args ...string) (pgx.Rows, error) {
	query, err := r.client.Query(ctx, request, args)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %w", err)
	}

	if query == nil {
		return nil, errors.New("no rows returned")
	}

	return query, nil
}
