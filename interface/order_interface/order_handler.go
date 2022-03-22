package order_interface

import (
	"github.com/gin-gonic/gin"

	orderApplicationService "github.com/garfcat/go-ddd/application/service/order"
)

type Order interface {
	CreateOrder(c *gin.Context)
}
type order struct {
	orderApplicationService orderApplicationService.Service
}

func NewOrderInterface(cfgs ...Configuration) (Order, error) {
	o := &order{}
	for _, cfg := range cfgs {
		err := cfg(o)
		if err != nil {
			return nil, err
		}
	}
	return o, nil
}

func (o *order) CreateOrder(c *gin.Context) {
	//o.orderApplicationService.CreateOrder()
}

// Configuration OrderConfiguration is an alias for a function that will take in a pointer to an OrderService and modify it
type Configuration func(o *order) error

func WithOrderApplicationService(orderApplicationService orderApplicationService.Service) Configuration {
	return func(o *order) error {
		o.orderApplicationService = orderApplicationService
		return nil
	}
}
