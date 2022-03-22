// Package customer holds aggregates that combines many entities into a full object
package entity

import (
	"errors"

	"github.com/garfcat/go-ddd/domain/customer/entity/valueObject"
	"github.com/google/uuid"
)

var (
	// ErrInvalidPerson is returned when the person is not valid in the NewCustom factory
	ErrInvalidPerson = errors.New("a customer has to have an valid person")
)

// Customer is an aggregate that combines all entities needed to represent a customer
type Customer struct {
	// ID is the identifier of the Entity, the ID is shared for all sub domains
	ID uuid.UUID
	// Name is the name of the person
	Name string
	// Age is the age of the person
	Age int
	// a customer can perform many transactions
	transactions []valueObject.Transaction
}

// NewCustomer is a factory to create a new Customer aggregate
// It will validate that the name is not empty
func NewCustomer(name string) (Customer, error) {
	// Validate that the Name is not empty
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	// Create a customer object and initialize all the values to avoid nil pointer exceptions
	return Customer{
		Name:         name,
		ID:           uuid.New(),
		transactions: make([]valueObject.Transaction, 0),
	}, nil
}

// GetID returns the customers root entity ID
func (c *Customer) GetID() uuid.UUID {
	return c.ID
}

// SetID sets the root ID
func (c *Customer) SetID(id uuid.UUID) {
	c.ID = id
}

// SetName changes the name of the Customer
func (c *Customer) SetName(name string) {
	c.Name = name
}

// GetName get the name of the Customer
func (c *Customer) GetName() string {
	return c.Name
}
