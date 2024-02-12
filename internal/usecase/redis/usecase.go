package redis

import (
	"github.com/vpbuyanov/gw-backend-go/internal/databases/redis"
)

type redisUS struct {
	repos *redis.Redis
}

type USRedis interface {
}

func NewRedisUseCase(repos *redis.Redis) USRedis {
	return &redisUS{repos: repos}
}
