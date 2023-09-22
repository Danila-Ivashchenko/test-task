package enricher

import (
	"context"
	"go-kafka/internal/domain/model"
)

type Enricher interface {
	GetAge(context.Context, string) (*model.AgeDTO, error)
	GetGender(context.Context, string) (*model.GenderDTO, error)
	GetNationalities(context.Context, string) (*model.NationalitiesDTO, error)
}
