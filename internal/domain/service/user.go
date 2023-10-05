package service

import (
	"context"
	"go-kafka/internal/domain/errors"
	"go-kafka/internal/domain/model"
	"go-kafka/internal/domain/ports/enricher"
	"go-kafka/internal/domain/ports/hasher"
	"go-kafka/internal/domain/ports/storage"
	"sync"

	"golang.org/x/exp/slog"
)

type validor interface {
	ValidateRawUser(dto *model.RawUser) error
	ValidateUserToUpdate(dto *model.User) error
	ValidateId(id int64) error
	ValidateGetUsersDTO(dto *model.GetUsersDTO) error
}

type userService struct {
	storage  storage.UserStorage
	enricher enricher.Enricher
	validor  validor
	hasher   hasher.Hasher
	logger   *slog.Logger
}

func NewUserService(s storage.UserStorage, e enricher.Enricher, v validor, h hasher.Hasher, l *slog.Logger) *userService {
	return &userService{
		storage:  s,
		enricher: e,
		validor:  v,
		hasher:   h,
		logger:   l,
	}
}

func (s userService) AddUser(ctx context.Context, dto *model.UserInsertRequest) error {
	toHash := dto.ToHash
	user := &model.RawUser{Name: dto.Name, Surname: dto.Surname, Patronymic: dto.Patronymic}

	err := s.validor.ValidateRawUser(user)
	if err != nil {
		return err
	}
	errCh := make(chan error)
	age := &model.AgeDTO{}
	gender := &model.GenderDTO{}
	nationalities := &model.NationalitiesDTO{}

	go func() {
		defer close(errCh)
		wg := &sync.WaitGroup{}
		wg.Add(3)

		go func() {
			defer wg.Done()
			result, err := s.enricher.GetAge(ctx, user.Name)
			if err != nil {
				select {
				case errCh <- err:
					errCh <- err
				default:
				}
				return
			}
			age = result
		}()

		go func() {
			defer wg.Done()
			result, err := s.enricher.GetGender(ctx, user.Name)
			if err != nil {
				select {
				case errCh <- err:
					errCh <- err
				default:
				}
				return
			}
			gender = result
		}()

		go func() {
			defer wg.Done()
			result, err := s.enricher.GetNationalities(ctx, user.Name)
			if err != nil {
				select {
				case errCh <- err:
					errCh <- err
				default:
				}
				return
			}
			nationalities = result
		}()
		wg.Wait()

		fullUser := &model.User{
			RawUser:     *user,
			Age:         age.Age,
			Gender:      gender.Gender,
			Nationality: extractNationality(nationalities),
		}

		user, err := s.storage.AddUser(ctx, fullUser)
		if err != nil {
			select {
			case errCh <- err:
				errCh <- err
				return
			default:
				return
			}
		}

		if toHash {
			err = s.hasher.Set(user)
			if err != nil {
				s.logger.Error(slog.String("redis error", err.Error()).String())
			} else {
				s.logger.Debug(slog.Int64("redis: success to insert user with id", user.Id).String())
			}
		}

	}()
	select {
	case <-ctx.Done():
		return errors.ErrorTimeOut
	case err := <-errCh:
		return err
	}
}

func extractNationality(nts *model.NationalitiesDTO) string {
	if len(nts.Country) == 0 {
		return ""
	}
	maxProb := nts.Country[0].Probability
	nat := nts.Country[0].CountryId

	for i := 1; i < len(nts.Country); i++ {
		if nts.Country[i].Probability > maxProb {
			maxProb = nts.Country[i].Probability
			nat = nts.Country[i].CountryId
		}
	}

	return nat
}

func (s userService) GetUsers(ctx context.Context, dto *model.GetUsersDTO) (*model.Users, error) {
	err := s.validor.ValidateGetUsersDTO(dto)
	if err != nil {
		return nil, err
	}
	errCh := make(chan error)
	resultCh := make(chan *model.Users)

	go func() {
		defer close(errCh)
		defer close(resultCh)

		result, err := s.storage.GetUsers(ctx, dto)
		if err != nil {
			errCh <- err
			return
		}
		resultCh <- result
	}()

	select {
	case <-ctx.Done():
		return nil, errors.ErrorTimeOut
	case err := <-errCh:
		return nil, err
	case result := <-resultCh:
		return result, nil
	}
}

func (s userService) GetUserById(ctx context.Context, dto *model.IdDTO) (*model.User, error) {
	err := s.validor.ValidateId(dto.Id)
	if err != nil {
		return nil, err
	}
	errCh := make(chan error)
	resultCh := make(chan *model.User)

	go func() {
		defer close(errCh)
		defer close(resultCh)

		result, err := s.hasher.Get(dto.Id)
		if err == nil {
			resultCh <- result
			s.logger.Debug(slog.Int64("success to get user from redis with id", result.Id).String())
			return
		}

		result, err = s.storage.GetUserById(ctx, dto)
		if err != nil {
			errCh <- err
			return
		}
		resultCh <- result
	}()

	select {
	case <-ctx.Done():
		return nil, errors.ErrorTimeOut
	case err := <-errCh:
		return nil, err
	case result := <-resultCh:
		return result, nil
	}
}

func (s userService) DeleteUser(ctx context.Context, dto *model.IdDTO) error {
	err := s.validor.ValidateId(dto.Id)
	if err != nil {
		return err
	}
	errCh := make(chan error)

	go func() {
		defer close(errCh)
		go func() {
			err := s.hasher.Delete(dto.Id)
			if err != nil {
				s.logger.Debug(slog.Int64("redis: error to delete user with id", dto.Id).String())
			} else {
				s.logger.Debug(slog.Int64("redis: success to delete user with id", dto.Id).String())
			}
		}()

		errCh <- s.storage.DeleteUser(ctx, dto)
	}()

	select {
	case <-ctx.Done():
		return errors.ErrorTimeOut
	case err := <-errCh:
		return err
	}
}

func (s userService) UpdateUser(ctx context.Context, dto *model.User) error {
	err := s.validor.ValidateUserToUpdate(dto)
	if err != nil {
		return err
	}
	errCh := make(chan error)

	go func() {
		defer close(errCh)
		go func() {
			err := s.hasher.Set(dto)
			if err == nil {
				s.logger.Debug(slog.Int64("redis: success to update user with id", dto.Id).String())
			}
		}()
		errCh <- s.storage.UpdateUser(ctx, dto)
	}()

	select {
	case <-ctx.Done():
		return errors.ErrorTimeOut
	case err := <-errCh:
		return err

	}
}
