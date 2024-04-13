package redis

import (
	"fmt"

	"github.com/go-redis/redis"

	"github.com/vpbuyanov/gw-backend-go/configs"
)

type redisDatabases struct {
	url    string
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
	}
}

func (r *redisDatabases) Connect() error {
	opt, err := redis.ParseURL(r.url)
	if err != nil {
		return fmt.Errorf("could not parse redis url %s: %w", r.url, err)
	}

	r.client = redis.NewClient(opt)
	return nil
}

func (r *redisDatabases) Close() error {
	err := r.client.Close()
	if err != nil {
		return fmt.Errorf("could not close redis client: %w", err)
	}

	return nil
}
