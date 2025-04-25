package service

import (
	"context"
	"person-enrichment-service/server/entity"
	"person-enrichment-service/server/repository"
)


type PersonServiceImpl struct {
	repo           repository.PersonRepository
	enrichmentSvc  EnrichmentService
}

func NewPersonService(repo repository.PersonRepository, enrichmentSvc EnrichmentService) PersonService {
	return &PersonServiceImpl{
		repo:          repo,
		enrichmentSvc: enrichmentSvc,
	}
}

func (s *PersonServiceImpl) CreatePerson(ctx context.Context, req *entity.CreatePersonRequest) (*entity.PersonResponse, error) {
	person := &entity.Person{
		Name:       req.Name,
		Surname:    req.Surname,
	}

	enrichedPerson, err := s.enrichmentSvc.EnrichPersonData(ctx, req.Name)
	if err != nil {
		return nil, err
	} else {
		person.Age = enrichedPerson.Age
		person.Gender = enrichedPerson.Gender
		person.Nationality = enrichedPerson.Nationality
	}

	if err := s.repo.Create(person); err != nil {
		return nil, err
	}

	return personToResponse(person), nil
}

func (s *PersonServiceImpl) GetPersonByID(ctx context.Context, id uint) (*entity.PersonResponse, error) {
	person, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return personToResponse(person), nil
}

func (s *PersonServiceImpl) GetAllPersons(ctx context.Context, filter entity.FilterOptions) ([]entity.PersonResponse, int64, error) {
	people, total, err := s.repo.GetAll(filter)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]entity.PersonResponse, len(people))
	for i, p := range people {
		responses[i] = *personToResponse(&p)
	}

	return responses, total, nil
}

func (s *PersonServiceImpl) UpdatePerson(ctx context.Context, id uint, req *entity.UpdatePersonRequest) (*entity.PersonResponse, error) {
	person := &entity.UpdatePersonRequest{
		Name:        req.Name,
		Surname:     req.Surname,
		Age:         req.Age,
		Gender:      req.Gender,
		Nationality: req.Nationality,
	}

	if err := s.repo.Update(id, person); err != nil {
		return nil, err
	}

	updatedPerson, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return personToResponse(updatedPerson), nil
}

func (s *PersonServiceImpl) DeletePerson(ctx context.Context, id uint) error {

	if err := s.repo.Delete(id); err != nil {
		return err
	}

	return nil
}

func personToResponse(person *entity.Person) *entity.PersonResponse {
	return &entity.PersonResponse{
		ID:          person.ID,
		Name:        person.Name,
		Surname:     person.Surname,
		Age:         person.Age,
		Gender:      person.Gender,
		Nationality: person.Nationality,
		CreatedAt:   person.CreatedAt,
		UpdatedAt:   person.UpdatedAt,
	}
}