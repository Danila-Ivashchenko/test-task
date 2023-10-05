package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

type config interface {
	GetRedisPort() string
	GetRedisHost() string
	GetRedisPassword() string
	GetRedisDB() int
}

type redisClient struct {
	port     string
	host     string
	password string
	db       int
}

func NewRedisClient(cfg config) *redisClient {
	return &redisClient{
		port:     cfg.GetRedisPort(),
		host:     cfg.GetRedisHost(),
		password: cfg.GetRedisPassword(),
		db:       cfg.GetRedisDB(),
	}
}

func (r redisClient) GetClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.host, r.port),
		Password: r.password,
		DB:       r.db,
	})
}
