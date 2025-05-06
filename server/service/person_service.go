package service

import (
	"context"
	"person-enrichment-service/server/entity"
)

type PersonService interface {
	CreatePerson(ctx context.Context, req *entity.CreatePersonRequest) (*entity.PersonResponse, error)
	GetPersonByID(ctx context.Context, id uint) (*entity.PersonResponse, error)
	GetAllPersons(ctx context.Context, filterOptions entity.FilterOptions) ([]entity.PersonResponse, int64, error)
	UpdatePerson(ctx context.Context, id uint, req *entity.UpdatePersonRequest) (*entity.PersonResponse, error)
	DeletePerson(ctx context.Context, id uint) error
}
