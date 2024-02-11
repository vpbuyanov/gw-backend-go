package repos

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/vpbuyanov/gw-backend-go/configs"
)

type reposPg struct {
	url    string
	client *pgx.Conn
	ctx    context.Context
}

type Pg interface {
}

func NewReposPg(configs configs.Postgres) Pg {
	return &reposPg{
		ctx: context.Background(),
		url: fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
			configs.User, configs.Password, configs.Host, configs.Port, configs.DbName),
	}
}

func (r *reposPg) connect() error {
	client, err := pgx.Connect(r.ctx, r.url)
	if err != nil {
		return err
	}

	r.client = client
	return nil
}

func (r *reposPg) close() error {
	err := r.client.Close(r.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *reposPg) InsertUser() error {
	err := r.connect()
	if err != nil {
		return err
	}

	defer func(r *reposPg) {
		err = r.close()
		if err != nil {
			return
		}
	}(r)

	return nil
}
