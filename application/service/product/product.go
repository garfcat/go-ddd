package product

import (
	productService "github.com/garfcat/go-ddd/domain/product/service"
)

// Configuration is an alias for a function that will take in a pointer to a Service and modify it
type Configuration func(os *service) error
type Service interface {
}

// Service  is an implementation of the ProductService
type service struct {
	products productService.ProductDomainService
}

// NewProductService takes a variable amount of Configuration functions and returns a new ProductService
// Each Configuration will be called in the Product they are passed in
func NewProductService(cfgs ...Configuration) (Service, error) {
	// Create the product service
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

// WithProductDomainService applies a given customer repository to the ProductService
func WithProductDomainService(domainService productService.ProductDomainService) Configuration {
	// return a function that matches the ProductConfiguration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(os *service) error {
		os.products = domainService
		return nil
	}
}
