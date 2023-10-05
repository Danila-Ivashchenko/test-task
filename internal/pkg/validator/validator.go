package validator

import (
	"fmt"
	custom_errors "go-kafka/internal/domain/errors"
	"go-kafka/internal/domain/model"
	"regexp"

	"github.com/pkg/errors"
)

type validator struct{}

func New() *validator {
	return &validator{}
}

func (validator) ValidateRawUser(dto *model.RawUser) error {
	if dto.Name == "" {
		return custom_errors.ErrorNoName
	}

	if dto.Surname == "" {
		return custom_errors.ErrorNoSurname
	}

	if !isValidStr(dto.Name) {
		return errors.Wrap(custom_errors.ErrorInvalidName, dto.Name)
	}
	if !isValidStr(dto.Surname) {
		return errors.Wrap(custom_errors.ErrorInvalidSurname, dto.Surname)
	}
	if dto.Patronymic != "" && !isValidStr(dto.Patronymic) {
		return errors.Wrap(custom_errors.ErrorInvalidPatronymic, dto.Patronymic)
	}

	return nil
}

func isValidStr(str string) bool {
	pattern := "^[A-Z]+[a-z]+$"
	re := regexp.MustCompile(pattern)
	return re.MatchString(str)
}

func (validator) ValidateId(id int64) error {
	if id == 0 {
		return custom_errors.ErrorNoId
	}
	if id < 0 {
		return custom_errors.ErrorInvalidId
	}
	return nil
}

func (v validator) ValidateGetUsersDTO(dto *model.GetUsersDTO) error {
	toChange := 0

	if dto.Name != "" {
		toChange++
		if !isValidStr(dto.Name) {
			return errors.Wrap(custom_errors.ErrorInvalidName, dto.Name)
		}
	}
	if dto.Surname != "" {
		toChange++
		if !isValidStr(dto.Surname) {
			return errors.Wrap(custom_errors.ErrorInvalidSurname, dto.Surname)
		}
	}
	if dto.Patronymic != "" {
		toChange++
		if !isValidStr(dto.Patronymic) {
			return errors.Wrap(custom_errors.ErrorInvalidPatronymic, dto.Patronymic)
		}
	}
	if dto.Gender != "" {
		toChange++
		if dto.Gender != "female" && dto.Gender != "male" {
			return errors.Wrap(custom_errors.ErrorInvalidGender, dto.Gender)
		}
	}
	if dto.Age != 0 {
		toChange++
		if dto.Age < 1 && dto.Age > 150 {
			return errors.Wrap(custom_errors.ErrorInvalidAge, fmt.Sprintf("%d", dto.Age))
		}
	}
	if dto.Nationality != "" {
		toChange++
		if !isValidNationality(dto.Nationality) {
			return errors.Wrap(custom_errors.ErrorInvalidNationality, dto.Nationality)
		}
	}
	if toChange == 0 {
		return custom_errors.ErrorNothingToChange
	}
	return nil
}

func (v validator) ValidateUserToUpdate(dto *model.User) error {
	toChange := 0
	err := v.ValidateId(dto.Id)
	if err != nil {
		return err
	}
	if dto.Name != "" {
		toChange++
		if !isValidStr(dto.Name) {
			return errors.Wrap(custom_errors.ErrorInvalidName, dto.Name)
		}
	}
	if dto.Surname != "" {
		toChange++
		if !isValidStr(dto.Surname) {
			return errors.Wrap(custom_errors.ErrorInvalidSurname, dto.Surname)
		}
	}
	if dto.Patronymic != "" {
		toChange++
		if !isValidStr(dto.Patronymic) {
			return errors.Wrap(custom_errors.ErrorInvalidPatronymic, dto.Patronymic)
		}
	}
	if dto.Gender != "" {
		toChange++
		if dto.Gender != "female" && dto.Gender != "male" {
			return errors.Wrap(custom_errors.ErrorInvalidGender, dto.Gender)
		}
	}
	if dto.Age != 0 {
		toChange++
		if dto.Age < 1 && dto.Age > 150 {
			return errors.Wrap(custom_errors.ErrorInvalidAge, fmt.Sprintf("%d", dto.Age))
		}
	}
	if dto.Nationality != "" {
		toChange++
		if !isValidNationality(dto.Nationality) {
			return errors.Wrap(custom_errors.ErrorInvalidNationality, dto.Nationality)
		}
	}
	if toChange == 0 {
		return custom_errors.ErrorNothingToChange
	}
	return nil
}

func isValidNationality(nat string) bool {
	pattern := "^[A-Z]{2}$"
	re := regexp.MustCompile(pattern)
	return re.MatchString(nat)
}
