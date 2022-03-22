package customer_interface

import (
	"net/http"

	"github.com/garfcat/go-ddd/domain/customer/entity"
	"github.com/gin-gonic/gin"

	customerApplicationService "github.com/garfcat/go-ddd/application/service/customer"
)

type Customer interface {
	CreateCustomer(c *gin.Context)
}
type customer struct {
	customerApplicationService customerApplicationService.Service
}

// Configuration is an alias for a function that will take in a pointer to an CustomerService and modify it
type Configuration func(o *customer) error

func WithCustomerApplicationService(CustomerApplicationService customerApplicationService.Service) Configuration {
	return func(o *customer) error {
		o.customerApplicationService = CustomerApplicationService
		return nil
	}
}

// NewCustomerInterface create interface for customer
func NewCustomerInterface(cfgs ...Configuration) (Customer, error) {
	o := &customer{}
	for _, cfg := range cfgs {
		err := cfg(o)
		if err != nil {
			return nil, err
		}
	}
	return o, nil
}

func (o *customer) CreateCustomer(c *gin.Context) {
	var user entity.Customer
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}
	err := o.customerApplicationService.CreateCustomer(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "success")
}
