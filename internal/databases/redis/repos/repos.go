package repos

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"

	"github.com/vpbuyanov/gw-backend-go/configs"
)

type reposRedis struct {
	url    string
	ctx    context.Context
	client *redis.Client
}

type Redis interface {
}

func NewReposRedis(configs configs.Redis) Redis {
	url := fmt.Sprintf("redis://%s:%s@%s:%s",
		configs.User, configs.Password, configs.Host, configs.Port)

	return &reposRedis{
		url: url,
		ctx: context.Background(),
	}
}

func (r *reposRedis) connect() error {
	opt, err := redis.ParseURL(r.url)
	if err != nil {
		panic(err)
	}

	r.client = redis.NewClient(opt)
	return nil
}

func (r *reposRedis) close() error {
	err := r.client.Close()
	if err != nil {
		return err
	}

	return nil
}
