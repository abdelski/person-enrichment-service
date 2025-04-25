package repository

import (
	"fmt"
	"person-enrichment-service/server/entity"

	"gorm.io/gorm"
)

type PersonRepositoryImpl struct {
	db *gorm.DB
}

func NewPersonRepository(db *gorm.DB) PersonRepository {
	return &PersonRepositoryImpl{db: db}
}

func (r *PersonRepositoryImpl) Create(person *entity.Person) error {
	return r.db.Create(person).Error
}

func (r *PersonRepositoryImpl) GetByID(id uint) (*entity.Person, error) {
	var person entity.Person
	if err := r.db.First(&person, id).Error; err != nil {
		return nil, err
	}
	return &person, nil
}

func (r *PersonRepositoryImpl) GetAll(filter entity.FilterOptions) ([]entity.Person, int64, error) {
	var people []entity.Person
	var total int64

	query := r.db.Model(&entity.Person{})

	if filter.Name != "" {
		query = query.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", filter.Name))
	}
	if filter.Surname != "" {
		query = query.Where("surname ILIKE ?", fmt.Sprintf("%%%s%%", filter.Surname))
	}
	if filter.Age > 0 {
		query = query.Where("age = ?", filter.Age)
	}
	if filter.Gender != "" {
		query = query.Where("gender = ?", filter.Gender)
	}
	if filter.Nationality != "" {
		query = query.Where("nationality = ?", filter.Nationality)
	}

	// Count total records
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	offset := (filter.Page - 1) * filter.PageSize
	if err := query.Offset(offset).Limit(filter.PageSize).Find(&people).Error; err != nil {
		return nil, 0, err
	}

	return people, total, nil
}

func (r *PersonRepositoryImpl) Update(id uint, person *entity.UpdatePersonRequest) error {
	return r.db.Model(&entity.Person{}).Where("id = ?", id).Updates(person).Error
}

func (r *PersonRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entity.Person{}, id).Error
}