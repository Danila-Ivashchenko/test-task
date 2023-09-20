package storage

import (
	"context"
	"go-kafka/internal/domain/model"
)

type UserStorage interface {
	AddUser(context.Context, *model.FullUser) error
	GetUsers(context.Context, *model.GetUsersDTO)
}
