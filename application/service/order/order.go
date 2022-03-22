// Package order holds all the services that connects repositories into a business flow related to ordering products
package order

import (
	"log"

	"github.com/google/uuid"

	"github.com/garfcat/go-ddd/domain/customer/entity"
	customerDomain "github.com/garfcat/go-ddd/domain/customer/service"
	orderDomain "github.com/garfcat/go-ddd/domain/order/service"
	orderService "github.com/garfcat/go-ddd/domain/order/service"
	entity2 "github.com/garfcat/go-ddd/domain/product/entity"
	prodmemory "github.com/garfcat/go-ddd/domain/product/repository/memory"
	productService "github.com/garfcat/go-ddd/domain/product/service"
)

// Configuration OrderConfiguration is an alias for a function that will take in a pointer to an OrderService and modify it
type Configuration func(os *service) error
type Service interface {
}

// Service  is an implementation of the OrderService
type service struct {
	customers          customerDomain.CustomerDomainService
	products           productService.ProductDomainService
	orderDomainService orderService.OrderDomainService
}

// NewOrderService takes a variable amount of OrderConfiguration functions and returns a new OrderService
// Each OrderConfiguration will be called in the order they are passed in
func NewOrderService(cfgs ...Configuration) (Service, error) {
	// Create the order service
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

//// WithCustomerRepository applies a given customer repository to the OrderService
//func WithCustomerRepository(cr repository.Repository) Configuration {
//	// return a function that matches the OrderConfiguration alias,
//	// You need to return this so that the parent function can take in all the needed parameters
//	return func(os *Service) error {
//		os.customers = cr
//		return nil
//	}
//}

// WithCustomerDomainService applies a given customer repository to the OrderService
func WithCustomerDomainService(domainService customerDomain.CustomerDomainService) Configuration {
	// return a function that matches the OrderConfiguration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(os *service) error {
		os.customers = domainService
		return nil
	}
}

// WithOrderDomainService  applies a given customer repository to the OrderService
func WithOrderDomainService(domainService orderDomain.OrderDomainService) Configuration {
	// return a function that matches the OrderConfiguration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(os *service) error {
		os.orderDomainService = domainService
		return nil
	}
}

// WithProductDomainService applies a given customer repository to the OrderService
func WithProductDomainService(domainService productService.ProductDomainService) Configuration {
	// return a function that matches the OrderConfiguration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(os *service) error {
		os.products = domainService
		return nil
	}
}

// WithMemoryProductRepository adds an in memory product repo and adds all input products
func WithMemoryProductRepository(products []entity2.Product) Configuration {
	return func(os *service) error {
		// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
		pr := prodmemory.New()

		// Add Items to repo
		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

// CreateOrder will chaintogether all repositories to create an order for a customer
// will return the collected price of all Products
func (o *service) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	// Get the customer
	c, err := o.customers.GetCustomer(customerID)
	if err != nil {
		return 0, err
	}

	// Get each Product, Ouchie, We need a ProductRepository
	var products []entity2.Product
	var price float64
	for _, id := range productIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		price += p.GetPrice()
	}

	// All Products exist in store, now we can create the order
	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))
	// Add Products and Update Customer

	return price, nil
}

// AddCustomer will add a new customer and return the customerID
func (o *service) AddCustomer(name string) (uuid.UUID, error) {
	c, err := entity.NewCustomer(name)
	if err != nil {
		return uuid.Nil, err
	}
	// Add to Repo
	err = o.customers.CreateCustomer(c)
	if err != nil {
		return uuid.Nil, err
	}

	return c.GetID(), nil
}
