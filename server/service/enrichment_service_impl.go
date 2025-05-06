package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"person-enrichment-service/server/entity"
	"strings"
)

type EnrichmentServiceImpl struct {
	agifyURL       string
	genderizeURL   string
	nationalizeURL string
	logger         *slog.Logger
}

func NewEnrichmentService(agifyURL, genderizeURL, nationalizeURL string, logger *slog.Logger) EnrichmentService {
	return &EnrichmentServiceImpl{
		agifyURL:       agifyURL,
		genderizeURL:   genderizeURL,
		nationalizeURL: nationalizeURL,
		logger:         logger,
	}
}

func (s *EnrichmentServiceImpl) EnrichPersonData(ctx context.Context, name string) (*entity.Person, error) {
	person := &entity.Person{Name: name}

	age, err := s.getAge(ctx, name)
	if err != nil {
		s.logger.Error("Failed to get age", "error", err)
	} else {
		person.Age = age
	}

	gender, err := s.getGender(ctx, name)
	if err != nil {
		s.logger.Error("Failed to get gender", "error", err)
	} else {
		person.Gender = gender
	}

	nationality, err := s.getNationality(ctx, name)
	if err != nil {
		s.logger.Error("Failed to get nationality", "error", err)
	} else {
		person.Nationality = nationality
	}

	return person, nil
}

func (s *EnrichmentServiceImpl) getAge(ctx context.Context, name string) (int, error) {
	url := fmt.Sprintf("%s/?name=%s", s.agifyURL, name)
	resp, err := http.Get(url)
	if err != nil {
		s.logger.Error("Failed to get age", "error", err)
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Error("Failed to read response body", "error", err)
		return 0, err
	}

	var result struct {
		Age int `json:"age"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		s.logger.Error("Failed to unmarshal response", "error", err)
		return 0, err
	}

	return result.Age, nil
}

func (s *EnrichmentServiceImpl) getGender(ctx context.Context, name string) (string, error) {
	url := fmt.Sprintf("%s/?name=%s", s.genderizeURL, name)
	resp, err := http.Get(url)
	if err != nil {
		s.logger.Error("Failed to get gender", "error", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Error("Failed to read response body", "error", err)
		return "", err
	}

	var result struct {
		Gender string `json:"gender"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		s.logger.Error("Failed to unmarshal response", "error", err)
		return "", err
	}

	return result.Gender, nil
}

func (s *EnrichmentServiceImpl) getNationality(ctx context.Context, name string) (string, error) {
	url := fmt.Sprintf("%s/?name=%s", s.nationalizeURL, name)
	resp, err := http.Get(url)
	if err != nil {
		s.logger.Error("Failed to get naitionality", "error", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Error("Failed to read response body", "error", err)
		return "", err
	}
	// s.logger.Debug("Abdel Response body", "body", string(body))

	var result struct {
		Country []struct {
			CountryID   string  `json:"country_id"`
			Probability float64 `json:"probability"`
		} `json:"country"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		s.logger.Error("Failed to unmarshal response", "error", err)
		return "", err
	}

	if len(result.Country) == 0 {
		return "", fmt.Errorf("no nationality data found")
	}

	var topCountry string
	var topProb float64
	for _, country := range result.Country {
		if country.Probability > topProb {
			topProb = country.Probability
			topCountry = country.CountryID
		}
	}

	return strings.ToLower(topCountry), nil
}
