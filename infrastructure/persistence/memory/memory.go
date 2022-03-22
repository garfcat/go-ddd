// Package memory is a in-memory implementation of the customer repository
package memory

import (
	"fmt"
	"sync"

	"github.com/garfcat/go-ddd/domain/customer/repository/po"

	"github.com/garfcat/go-ddd/domain/customer/repository"

	"github.com/google/uuid"
)

// Repository fulfills the Repository interface
type Repository struct {
	customers map[uuid.UUID]po.CustomerPo
	sync.Mutex
}

// New is a factory function to generate a new repository of customers
func New() *Repository {
	return &Repository{
		customers: make(map[uuid.UUID]po.CustomerPo),
	}
}

// Get finds a customer by ID
func (mr *Repository) Get(id uuid.UUID) (po.CustomerPo, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}

	return po.CustomerPo{}, repository.ErrCustomerNotFound
}

// Add will add a new customer to the repository
func (mr *Repository) Add(c po.CustomerPo) error {
	if mr.customers == nil {
		// Safety check if customers is not create, shouldn't happen if using the Factory, but you never know
		mr.Lock()
		mr.customers = make(map[uuid.UUID]po.CustomerPo)
		mr.Unlock()
	}
	// Make sure Customer isn't already in the repository
	if _, ok := mr.customers[c.ID]; ok {
		return fmt.Errorf("customer already exists: %w", repository.ErrFailedToAddCustomer)
	}
	mr.Lock()
	mr.customers[c.ID] = c
	mr.Unlock()
	return nil
}

// Update will replace an existing customer information with the new customer information
func (mr *Repository) Update(c po.CustomerPo) error {
	// Make sure Customer is in the repository
	if _, ok := mr.customers[c.ID]; !ok {
		return fmt.Errorf("customer does not exist: %w", repository.ErrUpdateCustomer)
	}
	mr.Lock()
	mr.customers[c.ID] = c
	mr.Unlock()
	return nil
}
