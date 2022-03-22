package service

import (
	"github.com/garfcat/go-ddd/domain/customer/entity"
	"github.com/garfcat/go-ddd/domain/customer/repository/po"
)

func CreateCustomer(po po.CustomerPo) entity.Customer {
	return entity.Customer{}
}

func CreateCustomerPo(customer entity.Customer) po.CustomerPo {
	return po.CustomerPo{}
}
