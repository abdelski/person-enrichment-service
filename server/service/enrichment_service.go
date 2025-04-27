package service

import (
	"context"
	"person-enrichment-service/server/entity"
)

type EnrichmentService interface {
	EnrichPersonData(ctx context.Context, name string) (*entity.Person, error)
}
