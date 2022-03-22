// Package repository  holds all the domain logic for the customer domain.
package repository

import (
	"errors"

	"github.com/garfcat/go-ddd/domain/customer/repository/po"
	"github.com/google/uuid"
)

var (
	// ErrCustomerNotFound is returned when a customer is not found.
	ErrCustomerNotFound = errors.New("the customer was not found in the repository")
	// ErrFailedToAddCustomer is returned when the customer could not be added to the repository.
	ErrFailedToAddCustomer = errors.New("failed to add the customer to the repository")
	// ErrUpdateCustomer is returned when the customer could not be updated in the repository.
	ErrUpdateCustomer = errors.New("failed to update the customer in the repository")
)

// Repository is an interface that defines the rules around what a customer repository
// Has to be able to perform
type Repository interface {
	Get(uuid.UUID) (po.CustomerPo, error)
	Add(po.CustomerPo) error
	Update(po.CustomerPo) error
}
