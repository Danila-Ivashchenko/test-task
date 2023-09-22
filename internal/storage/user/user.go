package user

import (
	"context"
	"go-kafka/internal/domain/model"

	"github.com/go-pg/pg"
)

type psqlClient interface {
	GetDb() *pg.DB
}

type userStorage struct {
	psqlClient psqlClient
}

func NewUserStorage(pc psqlClient) *userStorage {
	return &userStorage{
		psqlClient: pc,
	}
}

func (s userStorage) AddUser(ctx context.Context, dto *model.User) error {
	db := s.psqlClient.GetDb()
	defer db.Close()

	_, err := db.Model(dto).Insert()
	if err != nil {
		return err
	}

	return nil
}

func (s userStorage) GetUsers(context.Context, *model.GetUsersDTO) {

}
