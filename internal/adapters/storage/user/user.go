package user

import (
	"context"
	"fmt"
	custom_errors "go-kafka/internal/domain/errors"
	"go-kafka/internal/domain/model"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
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

func (s userStorage) AddUser(ctx context.Context, dto *model.User) (*model.User, error) {
	db := s.psqlClient.GetDb()
	defer db.Close()

	_, err := db.ModelContext(ctx, dto).Insert()
	if err != nil {
		return nil, custom_errors.ErrorToAddUser
	}

	return dto, nil
}

func (s userStorage) GetUsers(ctx context.Context, dto *model.GetUsersDTO) (*model.Users, error) {
	db := s.psqlClient.GetDb()
	defer db.Close()

	users := &model.Users{}
	whereCase := dto.ExtractSQL()
	query := &orm.Query{}
	if whereCase != "" {
		query = db.ModelContext(ctx, &users.UsersArr).Where(whereCase).Order("id DESC")
	} else {
		query = db.ModelContext(ctx, &users.UsersArr).Order("id DESC")
	}

	if dto.Limit != 0 {
		query.Limit(dto.Limit)
	}
	if dto.Offset != 0 {
		fmt.Println(dto.Offset)
		query.Offset(dto.Offset)
	}
	err := query.Select()
	if err != nil {
		return nil, err
	}
	users.Count = int64(len(users.UsersArr))
	if users.Count == 0 {
		return nil, custom_errors.ErrorNoSuchUsers
	}
	return users, err
}

func (s userStorage) GetUserById(ctx context.Context, dto *model.IdDTO) (*model.User, error) {
	db := s.psqlClient.GetDb()
	defer db.Close()

	user := model.User{Id: dto.Id}
	query := db.ModelContext(ctx, &user).Where("id=?", dto.Id)

	err := query.Select()
	if err != nil {
		return nil, custom_errors.ErrorNoSuchUser
	}

	return &user, err
}

func (s userStorage) DeleteUser(ctx context.Context, dto *model.IdDTO) error {
	db := s.psqlClient.GetDb()
	defer db.Close()

	user := &model.User{Id: dto.Id}
	result, err := db.ModelContext(ctx, user).WherePK().Delete()

	if result.RowsAffected() == 0 {
		return custom_errors.ErrorNoSuchUser
	}
	return err
}

func (s userStorage) UpdateUser(ctx context.Context, dto *model.User) error {
	db := s.psqlClient.GetDb()
	defer db.Close()

	result, err := db.ModelContext(ctx, dto).Set(dto.ExtractSQLToUpdate()).WherePK().Update()
	if err != nil {
		return custom_errors.ErrorNoSuchUser
	}
	if result.RowsAffected() == 0 {
		return custom_errors.ErrorNoSuchUser
	}
	return err
}
