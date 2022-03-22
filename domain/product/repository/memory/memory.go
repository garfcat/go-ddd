// Package memory is a in memory implementation of the ProductRepository interface.
package memory

import (
	"sync"

	"github.com/garfcat/go-ddd/domain/product/entity"
	"github.com/garfcat/go-ddd/domain/product/repository"

	"github.com/google/uuid"
)

type ProductRepository struct {
	products map[uuid.UUID]entity.Product
	sync.Mutex
}

// New is a factory function to generate a new repository of customers
func New() *ProductRepository {
	return &ProductRepository{
		products: make(map[uuid.UUID]entity.Product),
	}
}

// GetAll returns all products as a slice
// Yes, it never returns an error, but
// A database implementation could return an error for instance
func (mpr *ProductRepository) GetAll() ([]entity.Product, error) {
	// Collect all Products from map
	var products []entity.Product
	for _, product := range mpr.products {
		products = append(products, product)
	}
	return products, nil
}

// GetByID searches for a product based on it's ID
func (mpr *ProductRepository) GetByID(id uuid.UUID) (entity.Product, error) {
	if product, ok := mpr.products[uuid.UUID(id)]; ok {
		return product, nil
	}
	return entity.Product{}, repository.ErrProductNotFound
}

// Add will add a new product to the repository
func (mpr *ProductRepository) Add(newprod entity.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[newprod.GetID()]; ok {
		return repository.ErrProductAlreadyExist
	}

	mpr.products[newprod.GetID()] = newprod

	return nil
}

// Update will change all values for a product based on it's ID
func (mpr *ProductRepository) Update(upprod entity.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[upprod.GetID()]; !ok {
		return repository.ErrProductNotFound
	}

	mpr.products[upprod.GetID()] = upprod
	return nil
}

// Delete remove an product from the repository
func (mpr *ProductRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return repository.ErrProductNotFound
	}
	delete(mpr.products, id)
	return nil
}
