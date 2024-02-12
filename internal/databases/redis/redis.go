package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"

	"github.com/vpbuyanov/gw-backend-go/configs"
)

type redisDatabases struct {
	url    string
	ctx    context.Context
	client *redis.Client
}

type Redis interface {
	Connect() error
	Close() error
}

func NewReposRedis(configs configs.Redis) Redis {
	url := fmt.Sprintf("redis://%s:%s@%s:%s",
		configs.User, configs.Password, configs.Host, configs.Port)

	return &redisDatabases{
		url: url,
		ctx: context.Background(),
	}
}

func (r *redisDatabases) Connect() error {
	opt, err := redis.ParseURL(r.url)
	if err != nil {
		return err
	}

	r.client = redis.NewClient(opt)
	return nil
}

func (r *redisDatabases) Close() error {
	err := r.client.Close()
	if err != nil {
		return err
	}

	return nil
}
