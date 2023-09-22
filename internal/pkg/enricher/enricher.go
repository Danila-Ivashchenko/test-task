package enricher

import (
	"context"
	"encoding/json"
	"fmt"
	"go-kafka/internal/domain/model"
	"io"
	"net/http"
)

// type config interface {
// 	GetEnrichUrl() string
// }

type enricher struct {
	// url string
	client *http.Client
}

func New() *enricher {
	return &enricher{
		// url: cfg.GetEnrichUrl(),
		client: &http.Client{},
	}
}

func (e enricher) GetAge(ctx context.Context, name string) (*model.AgeDTO, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s?name=%s", "https://api.agify.io/", name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := e.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	dto := &model.AgeDTO{}
	err = json.Unmarshal(body, &dto)
	if err != nil {
		return nil, err
	}
	return dto, nil
}
func (e enricher) GetGender(ctx context.Context, name string) (*model.GenderDTO, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s?name=%s", "https://api.genderize.io/", name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := e.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	dto := &model.GenderDTO{}
	err = json.Unmarshal(body, &dto)
	if err != nil {
		return nil, err
	}
	return dto, nil
}
func (e enricher) GetNationalities(ctx context.Context, name string) (*model.NationalitiesDTO, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s?name=%s", "https://api.nationalize.io/", name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := e.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	dto := &model.NationalitiesDTO{}
	err = json.Unmarshal(body, &dto)
	if err != nil {
		return nil, err
	}
	return dto, nil
}
