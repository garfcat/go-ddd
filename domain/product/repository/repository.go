// Package product holds the repository and the implementations for a ProductRepository
package repository

import (
	"errors"

	"github.com/garfcat/go-ddd/domain/product/entity"

	"github.com/google/uuid"
)

var (
	//ErrProductNotFound is returned when a product is not found
	ErrProductNotFound = errors.New("the product was not found")
	//ErrProductAlreadyExist is returned when trying to add a product that already exists
	ErrProductAlreadyExist = errors.New("the product already exists")
)

// ProductRepository is the repository interface to fulfill to use the product aggregate
type ProductRepository interface {
	GetAll() ([]entity.Product, error)
	GetByID(uuid.UUID) (entity.Product, error)
	Add(entity.Product) error
	Update(entity.Product) error
	Delete(uuid.UUID) error
}
