package product_interface

import (
	productApplicationService "github.com/garfcat/go-ddd/application/service/product"
	"github.com/gin-gonic/gin"
)

type Product interface {
	CreateProduct(c *gin.Context)
}
type product struct {
	ProductApplicationService productApplicationService.Service
}

func NewProductInterface(cfgs ...Configuration) (Product, error) {
	o := &product{}
	for _, cfg := range cfgs {
		err := cfg(o)
		if err != nil {
			return nil, err
		}
	}
	return o, nil
}

func (o *product) CreateProduct(c *gin.Context) {
}

// Configuration ProductConfiguration is an alias for a function that will take in a pointer to an ProductService and modify it
type Configuration func(o *product) error

func WithProductApplicationService(ProductApplicationService productApplicationService.Service) Configuration {
	return func(o *product) error {
		o.ProductApplicationService = ProductApplicationService
		return nil
	}
}
