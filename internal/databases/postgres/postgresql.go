package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/vpbuyanov/gw-backend-go/configs"
)

type postgresql struct {
	url    string
	client *pgx.Conn
	ctx    context.Context
}

type Postgresql interface {
	Connect() error
	Close() error
}

func NewReposPostgresql(configs configs.Postgres) Postgresql {
	return &postgresql{
		ctx: context.Background(),
		url: fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
			configs.User, configs.Password, configs.Host, configs.Port, configs.DbName),
	}
}

func (r *postgresql) Connect() error {
	client, err := pgx.Connect(r.ctx, r.url)
	if err != nil {
		return err
	}

	r.client = client
	return nil
}

func (r *postgresql) Close() error {
	err := r.client.Close(r.ctx)
	if err != nil {
		return err
	}

	return nil
}
