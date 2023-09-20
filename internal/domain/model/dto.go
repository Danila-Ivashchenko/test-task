package model

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

type AgeDTO struct {
	Age uint `json:"age"`
}

type GenderDTO struct {
	Gender      string `json:"gender"`
	Probability string `json:"probability"`
}

type NationalityDTO struct {
	CountryId   string  `json:"country_id"`
	Probability float32 `json:"probability"`
}
