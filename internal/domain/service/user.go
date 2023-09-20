package service

import (
	"context"
	"go-kafka/internal/domain/errors"
	"go-kafka/internal/domain/model"
	"go-kafka/internal/domain/ports/enricher"
	"go-kafka/internal/domain/ports/storage"
	"sync"
)

type userService struct {
	storage  storage.UserStorage
	enricher enricher.Enricher
}

func NewUserStorage(s storage.UserStorage, e enricher.Enricher) *userService {
	return &userService{
		storage:  s,
		enricher: e,
	}
}

func (s userService) AddUser(ctx context.Context, user *model.User) error {
	errCh := make(chan error)

	age := &model.AgeDTO{}
	gender := &model.GenderDTO{}
	nationalities := []model.NationalityDTO{}

	go func() {
		defer close(errCh)
		wg := &sync.WaitGroup{}
		wg.Add(3)

		go func() {
			result, err := s.enricher.GetAge(ctx, user.Name)
			if err != nil {
				errCh <- err
				return
			}
			age = result
		}()

		go func() {
			result, err := s.enricher.GetGender(ctx, user.Name)
			if err != nil {
				errCh <- err
				return
			}
			gender = result
		}()

		go func() {
			result, err := s.enricher.GetNationalities(ctx, user.Name)
			if err != nil {
				errCh <- err
				return
			}
			nationalities = result
		}()

		wg.Wait()
		fullUser := &model.FullUser{
			User:        *user,
			Age:         age.Age,
			Gender:      gender.Gender,
			Nationality: extractNationality(nationalities),
		}

		err := s.storage.AddUser(ctx, fullUser)
		errCh <- err
	}()
	select {
	case <-ctx.Done():
		return errors.ErrorTimeOut
	case err := <-errCh:
		return err
	}

}

func extractNationality(nts []model.NationalityDTO) string {
	if len(nts) == 0 {
		return ""
	}
	maxProb := nts[0].Probability
	nat := nts[0].CountryId

	for i := 1; i < len(nts); i++ {
		if nts[i].Probability > maxProb {
			maxProb = nts[i].Probability
			nat = nts[i].CountryId
		}
	}

	return nat
}
