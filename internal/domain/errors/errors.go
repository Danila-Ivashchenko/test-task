package errors

import "errors"

var (
	ErrorTimeOut = errors.New("time out error")

	ErrorNoName    = errors.New("no name in input")
	ErrorNoSurname = errors.New("no surname in input")

	ErrorInvalidName        = errors.New("invalid name")
	ErrorInvalidSurname     = errors.New("invalid surname")
	ErrorInvalidPatronymic  = errors.New("invalid patronymic")
	ErrorInvalidGender      = errors.New("invalid gender")
	ErrorInvalidAge         = errors.New("invalid age")
	ErrorInvalidNationality = errors.New("invalid nationality")

	ErrorNothingToChange = errors.New("nothing to change")
	ErrorNoId            = errors.New("no id")
	ErrorInvalidId       = errors.New("invalid id")

	ErrorNoSuchUser = errors.New("no such user")
	ErrorNoSuchUsers = errors.New("no such users")
	ErrorToAddUser  = errors.New("error to add user")
)
