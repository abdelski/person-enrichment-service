package repository

import "person-enrichment-service/server/entity"

type PersonRepository interface {
	Create(person *entity.Person) error
	GetByID(id uint) (*entity.Person, error)
	GetAll(filterOptions entity.FilterOptions) ([]entity.Person, int64, error)
	Update(id uint, person *entity.UpdatePersonRequest) error
	Delete(id uint) error
}
