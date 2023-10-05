package storage

import (
	"context"
	"go-kafka/internal/domain/model"
)

type UserStorage interface {
	AddUser(context.Context, *model.User) (*model.User, error)
	GetUsers(context.Context, *model.GetUsersDTO) (*model.Users, error)
	GetUserById(context.Context, *model.IdDTO) (*model.User, error)
	DeleteUser(context.Context, *model.IdDTO) error
	UpdateUser(context.Context, *model.User) error
}
