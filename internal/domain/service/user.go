package service

import (
	"context"
	"fmt"
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

func (s userService) AddUser(ctx context.Context, user *model.RawUser) error {
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
				errCh <- err
				return
			}
			fmt.Println(result)
			age = result
		}()

		go func() {
			defer wg.Done()
			result, err := s.enricher.GetGender(ctx, user.Name)
			if err != nil {
				errCh <- err
				return
			}
			fmt.Println(result)
			gender = result
		}()

		go func() {
			defer wg.Done()
			result, err := s.enricher.GetNationalities(ctx, user.Name)
			if err != nil {
				errCh <- err
				return
			}
			fmt.Println(result)
			nationalities = result
		}()

		wg.Wait()
		fullUser := &model.User{
			RawUser:        *user,
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
