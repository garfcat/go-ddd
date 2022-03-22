// Package Customer holds all the services that connects repositories into a business flow related to Customering products
package customer

import (
	"github.com/garfcat/go-ddd/domain/customer/entity"
	customerService "github.com/garfcat/go-ddd/domain/customer/service"
)

type Service interface {
	CreateCustomer(customer entity.Customer) error
}

// Configuration CustomerConfiguration is an alias for a function that will take in a pointer to an Customerservice and modify it
type Configuration func(os *service) error

// service  is an implementation of the Customerservice
type service struct {
	customers customerService.CustomerDomainService
}

// NewCustomerService takes a variable amount of CustomerConfiguration functions and returns a new Customerservice
// Each CustomerConfiguration will be called in the Customer they are passed in
func NewCustomerService(cfgs ...Configuration) (Service, error) {
	// Create the Customer service
	os := &service{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// WithCustomerDomainService applies a given customer repository to the Customer service
func WithCustomerDomainService(domainService customerService.CustomerDomainService) Configuration {
	// return a function that matches the CustomerConfiguration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(cs *service) error {
		cs.customers = domainService
		return nil
	}
}

// CreateCustomer will chaintogether all repositories to create an Customer for a customer
// will return the collected price of all Products
func (s *service) CreateCustomer(customer entity.Customer) error {
	return s.customers.CreateCustomer(customer)
}
