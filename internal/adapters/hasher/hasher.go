package hasher

import (
	"encoding/json"
	"fmt"
	"go-kafka/internal/domain/errors"
	"go-kafka/internal/domain/model"
	"time"

	"github.com/go-redis/redis"
)

type redisClienter interface {
	GetClient() *redis.Client
}

type configer interface {
	GetRedisTtl() time.Duration
}

type hasher struct {
	redisCliente redisClienter
	ttl          time.Duration
}

func New(cfg configer, r redisClienter) *hasher {
	return &hasher{
		redisCliente: r,
		ttl:          cfg.GetRedisTtl(),
	}
}

func (h hasher) Set(user *model.User) error {
	client := h.redisCliente.GetClient()
	defer client.Close()

	return client.Set(fmt.Sprintf("%d", user.Id), user, h.ttl).Err()
}

func (h hasher) Get(id int64) (*model.User, error) {
	client := h.redisCliente.GetClient()
	defer client.Close()

	result, err := client.Get(fmt.Sprintf("%d", id)).Result()
	if err != nil {
		return nil, errors.ErrorNoSuchUser
	}
	user := &model.User{}
	err = json.Unmarshal([]byte(result), &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (h hasher) Delete(id int64) error {
	client := h.redisCliente.GetClient()
	defer client.Close()

	return client.Del(fmt.Sprintf("%d", id)).Err()
}

func (h hasher) Update(user *model.User) error {
	client := h.redisCliente.GetClient()
	defer client.Close()

	_, err := client.Get(fmt.Sprintf("%d", user.Id)).Result()
	if err != nil {
		return errors.ErrorNoSuchUser
	}

	return h.Set(user)
}
