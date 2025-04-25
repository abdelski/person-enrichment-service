package entity

import (
	"time"

	"gorm.io/gorm"
)


type Person struct {
	gorm.Model
	Name 		string `gorm:"not null"`
	Surname 	string `gorm:"not null"`
	Age 		int
	Gender 		string
	Nationality string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

type CreatePersonRequest struct {
	Name 	string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
}

type UpdatePersonRequest struct {
	Name       	string `json:"name,omitempty"`
	Surname    	string `json:"surname,omitempty"`
	Age        	int    `json:"age,omitempty"`
	Gender     	string `json:"gender,omitempty"`
	Nationality string `json:"nationality,omitempty"`
}

type PersonResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	Age         int       `json:"age,omitempty"`
	Gender      string    `json:"gender,omitempty"`
	Nationality string    `json:"nationality,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type FilterOptions struct {
	Name        string
	Surname     string
	Age         int
	Gender      string
	Nationality string
	Page        int
	PageSize    int
}