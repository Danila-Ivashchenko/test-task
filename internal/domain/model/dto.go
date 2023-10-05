package model

import "fmt"

type GetUsersDTO struct {
	Name        string `json:"name" db:"name"`
	Surname     string `json:"surname" db:"surname"`
	Patronymic  string `json:"patronymic" db:"patronymic"`
	Age         uint   `json:"age" db:"age"`
	Gender      string `json:"gender" db:"gender"`
	Nationality string `json:"nationality" db:"nationality"`
	Limit       int    `json:"limit"`
	Offset      int    `json:"offset"`
}

func (dto GetUsersDTO) ExtractSQL() string {
	result := ""
	if dto.Name != "" {
		result += fmt.Sprintf(`name='%s'`, dto.Name)
	}
	if dto.Surname != "" {
		if result != "" {
			result += " AND "
		}
		result += fmt.Sprintf(`sername='%s'`, dto.Surname)
	}
	if dto.Patronymic != "" {
		if result != "" {
			result += " AND "
		}
		result += fmt.Sprintf(`patronymic='%s'`, dto.Patronymic)
	}
	if dto.Age != 0 {
		if result != "" {
			result += " AND "
		}
		result += fmt.Sprintf(`age=%d`, dto.Age)
	}
	if dto.Gender != "" {
		if result != "" {
			result += " AND "
		}
		result += fmt.Sprintf(`gender='%s'`, dto.Gender)
	}
	if dto.Nationality != "" {
		if result != "" {
			result += " AND "
		}
		result += fmt.Sprintf(`nationality='%s'`, dto.Nationality)
	}

	return result
}

type AgeDTO struct {
	Age uint `json:"age"`
}
type GenderDTO struct {
	Gender      string  `json:"gender"`
	Probability float32 `json:"probability"`
}

type Nationality struct {
	CountryId   string  `json:"country_id"`
	Probability float32 `json:"probability"`
}

type NationalitiesDTO struct {
	Country []Nationality `json:"country"`
}

type IdDTO struct {
	Id int64 `json:"id"`
}
