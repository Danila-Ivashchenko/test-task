package model

import (
	"encoding/json"
	"fmt"
)

type RawUser struct {
	Name       string `json:"name" db:"name"`
	Surname    string `json:"surname" db:"surname"`
	Patronymic string `json:"patronymic" db:"patronymic"`
}

type UserInsertRequest struct {
	RawUser
	ToHash bool `json:"to_hash"`
}

type User struct {
	Id int64 `json:"id" db:"id"`
	RawUser
	Age         uint   `json:"age" db:"age"`
	Gender      string `json:"gender" db:"gender"`
	Nationality string `json:"nationality" db:"nationality"`
}

func (u User) MarshalBinary() ([]byte, error) {

	return json.Marshal(u)
}

func (u User) ExtractSQLToUpdate() string {
	result := ""
	if u.Name != "" {
		result += fmt.Sprintf(`name='%s'`, u.Name)
	}
	if u.Surname != "" {
		if result != "" {
			result += ", "
		}
		result += fmt.Sprintf(`sername='%s'`, u.Surname)
	}
	if u.Patronymic != "" {
		if result != "" {
			result += ", "
		}
		result += fmt.Sprintf(`patronymic='%s'`, u.Patronymic)
	}
	if u.Age != 0 {
		if result != "" {
			result += ", "
		}
		result += fmt.Sprintf(`age=%d`, u.Age)
	}
	if u.Gender != "" {
		if result != "" {
			result += ", "
		}
		result += fmt.Sprintf(`gender='%s'`, u.Gender)
	}
	if u.Nationality != "" {
		if result != "" {
			result += ", "
		}
		result += fmt.Sprintf(`nationality='%s'`, u.Nationality)
	}

	return result
}

type Users struct {
	Count    int64  `json:"count"`
	UsersArr []User `json:"users"`
}
